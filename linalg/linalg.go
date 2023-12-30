package linalg

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strconv"
	"text/tabwriter"
)

// row represents a row in an augmented matrix. The last (constant)
// column is pulled out into its own field, k.
type row struct {
	cs []*big.Rat // coefficients
	k  *big.Rat   // constant
}

var (
	ratZero *big.Rat           // 0
	ratOne  = big.NewRat(1, 1) // 1
)

// isRatZero checks if the given rational is zero.
func isRatZero(x *big.Rat) bool {
	return x.Num().BitLen() == 0
}

// zero returns true if all coefficients and the constant are 0.
func (r *row) zero() bool {
	if !isRatZero(r.k) {
		return false
	}
	return r.empty()
}

// empty returns true if all coefficients are 0.
func (r *row) empty() bool {
	for _, c := range r.cs {
		if !isRatZero(c) {
			return false
		}
	}
	return true
}

// impossible returns true if all coefficients are 0, but the constant is not.
func (r *row) impossible() bool {
	if isRatZero(r.k) {
		return false
	}
	return r.empty()
}

// coefficients returns the number of coefficients in the row.
func (r *row) coefficients() int {
	return len(r.cs)
}

// normalize multiplies r by whatever factor will make the first non-zero coefficient 1.
// If all coefficients are zero, and the constant is not, it'll set the constant to 1.
func (r *row) normalize() {
	factor := &big.Rat{}
	found := false

	for _, c := range r.cs {
		if isRatZero(c) {
			continue
		}
		if !found {
			factor.Inv(c)
			found = true
			c.Set(ratOne)
		} else {
			c.Mul(c, factor)
		}
	}

	if found {
		r.k.Mul(r.k, factor)
	} else {
		if !isRatZero(r.k) {
			r.k.Set(ratOne)
		}
	}
}

// firstNonZero returns the index of the first nonZero coefficient.
func (r *row) firstNonZero() int {
	for i, c := range r.cs {
		if !isRatZero(c) {
			return i
		}
	}
	return -1
}

// cmp compares two rows. It searches corresponding coefficients until
// they compare different, and returns that compare result. If all
// coefficients are equal, it returns the compare result of the
// constants.
func (r *row) cmp(other *row) int {
	for i, rc := range r.cs {
		if verdict := rc.Cmp(other.cs[i]); verdict != 0 {
			return verdict
		}
	}

	return r.k.Cmp(other.k)
}

// sub subtracts the other row from this one, leaving this row
// containing the result.
func (r *row) sub(other *row) {
	for i, rc := range r.cs {
		rc.Sub(rc, other.cs[i])
	}
	r.k.Sub(r.k, other.k)
}

// mul multiplies a row by a rational, leaving the original unchanged
// and returning the product.
func (r *row) mul(factor *big.Rat) *row {
	res := &row{}

	for _, c := range r.cs {
		newC := &big.Rat{}
		newC.Set(c).Mul(newC, factor)
		res.cs = append(res.cs, newC)
	}
	newK := &big.Rat{}
	res.k = newK.Set(r.k).Mul(newK, factor)
	return res
}

// orderBefore compares two rows for sorting during matrix reduction.
// coefficient. Row a should be ordered before row b if row a has a
// non-zero coefficient first. If all coefficients are zero, the row
// with the smaller constant should come first.
func (r *row) orderBefore(other *row) bool {
	rIndex := r.firstNonZero()
	otherIndex := other.firstNonZero()
	if rIndex == -1 {
		if otherIndex == -1 {
			return r.k.Cmp(other.k) < 0
		}
		return false
	}
	if otherIndex == -1 {
		return true
	}
	return rIndex < otherIndex
}

// matrix represents an augmented matrix, for reduction.
type Matrix struct {
	rows       []*row
	reduced    bool
	impossible bool
	rank       int
}

// NewMatrix creates a new matrix and returns it.
func NewMatrix(rows [][]int) *Matrix {
	res := &Matrix{}
	if len(rows) == 0 {
		return res
	}

	ll := len(rows[0])
	if ll < 2 {
		panic("Cannot create an augmented matrix with row length < 2; got " + strconv.Itoa(ll))
	}

	for i, row := range rows {
		if len(row) != ll {
			panic(fmt.Sprintf("Uneven row lengths: row 0 had %d ints, row %d has %d", ll, i, len(row)))
		}

		res.rows = append(res.rows, newRow(row[:ll-1], row[ll-1]))
	}

	return res
}

// newRow creates a new row object from integer coefficients and
// constant.
func newRow(coefficients []int, constant int) *row {
	res := &row{}
	for _, c := range coefficients {
		res.cs = append(res.cs, big.NewRat(int64(c), 1))
	}
	res.k = big.NewRat(int64(constant), 1)

	return res
}

// Impossible returns true if there are rows with impossible values:
// all zero coefficients adding to a non-zero constant. It performs
// reduction first, if it hasn't been done already.
func (m *Matrix) Impossible() bool {
	if !m.reduced {
		m.reduce()
	}

	return m.impossible
}

// Rank returns the rank of the matrix. It performs reduction first,
// if it hasn't been done already.
func (m *Matrix) Rank() int {
	if !m.reduced {
		m.reduce()
	}

	return m.rank
}

// Coefficients returns the number of coefficients in the matrix.
func (m *Matrix) Coefficients() int {
	return m.rows[0].coefficients()
}

// downwardPass does the downward subtracting pass of reduction,
// resulting in an upper-triangular matrix (or something somehow
// degenerate).
func (m *Matrix) downwardPass() {
	needSort := true

NEXT_ROW:
	for current := 0; current < len(m.rows)-1; current++ {
		// Largest first.
		if needSort {
			sort.Slice(m.rows[current:], func(i, j int) bool {
				return m.rows[i+current].orderBefore(m.rows[j+current])
			})
		}

		firstRow := m.rows[current]
		nz0 := firstRow.firstNonZero()
		if nz0 == -1 {
			// All the remaining rows have all-zero coefficients.
			return
		}

		for _, r := range m.rows[current+1:] {
			nzi := r.firstNonZero()
			if nzi == -1 || nzi > nz0 {
				continue NEXT_ROW
			}

			r.sub(firstRow)
			r.normalize()
			needSort = true
		}
	}
}

// upwardPass does the upward subtracting pass of reduction, trying to
// clear out all coefficients except the leading ones. If it
func (m *Matrix) upwardPass() {
	for current := 1; current < len(m.rows); current++ {
		clearer := m.rows[current]
		clnz := clearer.firstNonZero()
		if clnz == -1 {
			// All zero coefficients from here on; nothing more to do.
			return
		}

		for _, r := range m.rows[:current] {
			coeff := r.cs[clnz]
			if isRatZero(coeff) {
				continue
			}
			toSub := clearer.mul(coeff)
			r.sub(toSub)
		}
	}
}

// reduce tries to reduce the matrix by calling downwardPass and then
// upwardPass.
func (m *Matrix) reduce() {
	m.reduced = true
	for _, r := range m.rows {
		r.normalize()
		if r.impossible() {
			m.rank = -1
			m.impossible = true
			return
		}
	}

	m.downwardPass()
	m.upwardPass()

	// Calculate rank.
	rank := 0
	for _, r := range m.rows {
		if r.impossible() {
			m.rank = -1
			m.impossible = true
			return
		}
		if r.firstNonZero() != -1 {
			rank++
		}
	}
	m.rank = rank
}

// Equal returns true if the matrices have the same rows.
func (m *Matrix) Equal(other *Matrix) bool {
	if m.Coefficients() != other.Coefficients() {
		return false
	}

	if len(m.rows) != len(other.rows) {
		return false
	}

	for i, r := range m.rows {
		if r.cmp(other.rows[i]) != 0 {
			return false
		}
	}

	return true
}

// Printable returns a grid representation of the matrix.
func (m *Matrix) Printable(prefix string, brackets bool) string {
	prettyRat := func(r *big.Rat) string {
		if r.IsInt() {
			return r.Num().String()
		}
		return r.String()
	}

	if len(m.rows) == 0 {
		return prefix + "[no rows]"
	}
	b := &bytes.Buffer{}

	w := tabwriter.NewWriter(b, 0, 0, 1, ' ', tabwriter.AlignRight)

	tab := []byte("\t")

	for i, r := range m.rows {
		brack, ets := "", ""
		switch {
		case brackets && i == 0:
			brack, ets = "⎡ ", " ⎤"
		case brackets && i == len(m.rows)-1:
			brack, ets = "⎣ ", " ⎦"
		case brackets:
			brack, ets = "⎢ ", " ⎥"
		}

		io.WriteString(w, prefix+brack)
		// io.WriteString(w, prefix)
		for _, c := range r.cs {
			io.WriteString(w, prettyRat(c))
			w.Write(tab)
		}

		fmt.Fprintf(w, "⎪\t%s\t%s\n", prettyRat(r.k), ets)
	}

	w.Flush()
	// Remove trailing newline.
	b.Truncate(b.Len() - 1)
	return b.String()
}

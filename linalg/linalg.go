package linalg

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/big"
	"slices"
	"sort"
	"strconv"
	"text/tabwriter"
)

var (
	ratOne = big.NewRat(1, 1) // 1
)

// isRatZero checks if the given rational is zero.
func isRatZero(x *big.Rat) bool {
	return x.Num().BitLen() == 0
}

// Row represents a row in an augmented matrix. The last (constant)
// column is pulled out into its own field, k.
type Row struct {
	cs []*big.Rat // coefficients
	k  *big.Rat   // constant
}

// cloneRats clones a slice of *big.Rat
func cloneRats(rats []*big.Rat) []*big.Rat {
	if rats == nil {
		return nil
	}
	res := make([]*big.Rat, len(rats))
	for i, rat := range rats {
		res[i] = new(big.Rat).Set(rat)
	}
	return res
}

// Clone creates a clone of a Row.
func (r *Row) Clone() *Row {
	return &Row{
		cs: cloneRats(r.cs),
		k:  new(big.Rat).Set(r.k),
	}
}

// Coefficients returns the coefficients of the row.
func (r *Row) Coefficients() []*big.Rat {
	return r.cs
}

// Constant returns the constant of the row.
func (r *Row) Constant() *big.Rat {
	return r.k
}

// Zero returns true if all coefficients and the constant are 0.
func (r *Row) Zero() bool {
	if !isRatZero(r.k) {
		return false
	}
	return r.Empty()
}

// Empty returns true if all coefficients are 0.
func (r *Row) Empty() bool {
	for _, c := range r.cs {
		if !isRatZero(c) {
			return false
		}
	}
	return true
}

// NonZeroCoefficientCount returns the number of non-zero coefficients in the row.
func (r *Row) NonZeroCoefficientCount() int {
	count := 0
	for _, c := range r.cs {
		if !isRatZero(c) {
			count++
		}
	}

	return count
}

// impossible returns true if all coefficients are 0, but the constant is not.
func (r *Row) impossible() bool {
	if isRatZero(r.k) {
		return false
	}
	return r.Empty()
}

// knownCoefficient returns the coefficient this row identifies, if it
// does. For a row to identify a coefficient, it needs exactly one
// non-zero coefficient.
//
// The return value is intended to be slightly useful, even if there
// is no known coefficient: all zero coefficients will return -1, the
// row's constant, and false, while a non-zero coefficient followed by
// a later non-zero coefficient will return the index of that
// coeffiecint, the value it and the constant imply, and false.
func (r *Row) knownCoefficient() (int, *big.Rat, bool) {
	rat := &big.Rat{}
	rat.Set(r.k)

	index := r.firstNonZero()
	if index == -1 {
		return -1, rat, false
	}

	if r.cs[index].Cmp(ratOne) != 0 {
		rat.Inv(r.cs[index]).Mul(rat, r.k)
	}

	for _, c := range r.cs[index+1:] {
		if !isRatZero(c) {
			return index, rat, false
		}
	}

	return index, rat, true
}

// NumCoefficients returns the number of coefficients in the row.
func (r *Row) NumCoefficients() int {
	return len(r.cs)
}

// normalize multiplies r by whatever factor will make the first non-zero coefficient 1.
// If all coefficients are zero, and the constant is not, it'll set the constant to 1.
func (r *Row) normalize() {
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
func (r *Row) firstNonZero() int {
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
func (r *Row) cmp(other *Row) int {
	for i, rc := range r.cs {
		if verdict := rc.Cmp(other.cs[i]); verdict != 0 {
			return verdict
		}
	}

	return r.k.Cmp(other.k)
}

// sub subtracts the other row from this one, leaving this row
// containing the result.
func (r *Row) sub(other *Row) {
	for i, rc := range r.cs {
		rc.Sub(rc, other.cs[i])
	}
	r.k.Sub(r.k, other.k)
}

// mul multiplies a row by a rational, leaving the original unchanged
// and returning the product.
func (r *Row) mul(factor *big.Rat) *Row {
	res := &Row{}

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
func (r *Row) orderBefore(other *Row) bool {
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
	rows []*Row

	// cached values
	// filled in by reduce()
	reduced    bool
	impossible bool
	rank       int

	// filled in by Coefficients()
	coefficients      []*big.Rat
	coefficientsKnown []bool
}

// Clone creates a clone of a matrix as-is.
func (m *Matrix) Clone() *Matrix {
	var rows []*Row
	for _, row := range m.rows {
		rows = append(rows, row.Clone())
	}

	return &Matrix{
		rows:              rows,
		reduced:           m.reduced,
		impossible:        m.impossible,
		rank:              m.rank,
		coefficients:      cloneRats(m.coefficients),
		coefficientsKnown: slices.Clone(m.coefficientsKnown),
	}
}

// AddRow adds a new row to the matrix. If the matrix is already
// reduced, it reduces it again. If the row contains a different
// number of coefficients from rows already in the matrix, AddRow
// panics.
func (m *Matrix) AddRow(r *Row) {

	if len(m.rows) == 0 || m.NumCoefficients() == r.NumCoefficients() {
		m.rows = append(m.rows, r)
	} else {
		panic(fmt.Sprintf("Trying to add a row with %d coefficients to a matrix containing rows with %d coefficients",
			r.NumCoefficients(), m.NumCoefficients()))
	}

	if m.reduced {
		m.reduced, m.impossible, m.rank = false, false, 0
		m.coefficients, m.coefficientsKnown = nil, nil
		m.reduce()
	}
}

// AddRows adds multiple new rows to the matrix. If the matrix is
// already reduced, it reduces it again. If the row contains a
// different number of coefficients from rows already in the matrix,
// AddRow panics.
func (m *Matrix) AddRows(rows []*Row) {

	for _, r := range rows {
		if len(m.rows) == 0 || m.NumCoefficients() == r.NumCoefficients() {
			m.rows = append(m.rows, r)
		} else {
			panic(fmt.Sprintf("Trying to add a row with %d coefficients to a matrix containing rows with %d coefficients",
				r.NumCoefficients(), m.NumCoefficients()))
		}
	}

	if m.reduced {
		m.reduced, m.impossible, m.rank = false, false, 0
		m.coefficients, m.coefficientsKnown = nil, nil
		m.reduce()
	}
}

// AddRowInts adds a new row to the matrix, given integer coefficients
// and constant. If the matrix is already reduced, it reduces it
// again. If the row contains a different number of coefficients from
// rows already in the matrix, AddRow panics.
func (m *Matrix) AddRowInts(coefficients []int, constant int) {
	m.AddRow(NewRow(coefficients, constant))
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

		res.rows = append(res.rows, NewRow(row[:ll-1], row[ll-1]))
	}

	return res
}

// Rows returns the matrix's rows.
func (m *Matrix) Rows() []*Row {
	return m.rows
}

// NewRow creates a new row object from integer coefficients and
// constant.
func NewRow(coefficients []int, constant int) *Row {
	res := &Row{}
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
	m.reduceOnce()
	return m.impossible
}

// Rank returns the rank of the matrix. It performs reduction first,
// if it hasn't been done already.
func (m *Matrix) Rank() int {
	m.reduceOnce()
	return m.rank
}

// NumCoefficients returns the number of coefficients in the matrix.
func (m *Matrix) NumCoefficients() int {
	return m.rows[0].NumCoefficients()
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

func (m *Matrix) reduceOnce() {
	if m.reduced {
		return
	}
	m.reduce()
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

// KnownCoefficients reduces the matrix, and returns a slice of
// coefficient values and a slice of booleans denoting whether the
// corresponding coefficient is actually known.
func (m *Matrix) KnownCoefficients() ([]*big.Rat, []bool) {
	if m.coefficients != nil {
		return m.coefficients, m.coefficientsKnown
	}

	m.reduceOnce()
	count := m.NumCoefficients()

	m.coefficients = make([]*big.Rat, count)
	for i := range m.coefficients {
		m.coefficients[i] = &big.Rat{}
	}
	m.coefficientsKnown = make([]bool, count)

	if m.impossible {
		return m.coefficients, m.coefficientsKnown
	}

	for _, r := range m.rows {
		if index, val, known := r.knownCoefficient(); known {
			m.coefficients[index].Set(val)
			m.coefficientsKnown[index] = true
		}
	}

	return m.coefficients, m.coefficientsKnown
}

// KnownCoefficientCount calls KnownCoefficients, but just returns the
// number of known coefficients.
func (m *Matrix) KnownCoefficientCount() int {
	_, known := m.KnownCoefficients()
	count := 0
	for _, k := range known {
		if k {
			count++
		}
	}
	return count
}

// KnownCoefficientFloats returns almost the same result as
// KnownCoefficients, with Rationals converted to float64s. The
// difference is that if the magnitude of a rational is too large to
// be represented as a float64, the corresponding float will be an
// infinity, and the corresponding boolean will be set to false if it
// wasn't already false.
func (m *Matrix) KnownCoefficientFloats() ([]float64, []bool) {
	ratCoefficients, known := m.KnownCoefficients()
	known = slices.Clone(known)
	coefficients := make([]float64, len(ratCoefficients))
	for i, c := range ratCoefficients {
		floatVal, exact := c.Float64()
		coefficients[i] = floatVal
		if math.IsInf(floatVal, 0) && !exact {
			known[i] = false
		}
	}
	return coefficients, known
}

// Equal returns true if the matrices have the same rows.
func (m *Matrix) Equal(other *Matrix) bool {
	if m.NumCoefficients() != other.NumCoefficients() {
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
			io.WriteString(w, c.RatString())
			w.Write(tab)
		}

		fmt.Fprintf(w, "⎪\t%s\t%s\n", r.k.RatString(), ets)
	}

	w.Flush()
	// Remove trailing newline.
	b.Truncate(b.Len() - 1)
	return b.String()
}

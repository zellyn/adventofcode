package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type board struct {
	numbers [5][5]int
	seen    [5][5]bool
}

type state struct {
	draws  []int
	boards []*board
}

func (b *board) print() {
	printf("Board:\n")
	for _, row := range b.numbers {
		for _, i := range row {
			printf(" %2d", i)
		}
		printf("\n")
	}
}

func (b *board) see(n int) (win bool) {
	row, col := 0, 0

	found := false
OUTER:
	for ; row < 5; row++ {
		for col = 0; col < 5; col++ {
			if b.numbers[row][col] == n {
				found = true
				break OUTER
			}
		}
	}
	if !found {
		return false
	}

	b.seen[row][col] = true

	return b.rowSeen(row) || b.colSeen(col)
}

func (b *board) rowSeen(row int) bool {
	for col := 0; col < 5; col++ {
		if !b.seen[row][col] {
			return false
		}
	}
	return true
}

func (b *board) colSeen(col int) bool {
	for row := 0; row < 5; row++ {
		if !b.seen[row][col] {
			return false
		}
	}
	return true
}

func (b *board) unseen() []int {
	var res []int

	for row, seens := range b.seen {
		for col, seen := range seens {
			if !seen {
				res = append(res, b.numbers[row][col])
			}
		}
	}
	return res
}

func (s state) print() {
	printf("Draws: %v\n", s.draws)
	for _, b := range s.boards {
		b.print()
	}
}

func parseBoard(lines []string) (*board, error) {
	grid, err := util.ParseGrid(lines)
	if err != nil {
		return nil, err
	}

	if len(grid) != 5 {
		return nil, fmt.Errorf("want 5 lines in grid; got %d", len(grid))
	}

	res := board{}
	for i, row := range grid {
		if len(row) != 5 {
			return nil, fmt.Errorf("want 5 numbers in row; got %d: %v", len(row), row)
		}

		for j, num := range row {
			res.numbers[i][j] = num
		}
	}

	return &res, nil
}

func parse(inputs []string) (state, error) {
	paras := util.LinesByParagraph(inputs)
	draws, err := util.ParseInts(paras[0][0], ",")
	if err != nil {
		return state{}, err
	}

	boards, err := util.MapE(paras[1:], parseBoard)
	if err != nil {
		return state{}, err
	}

	return state{
		draws:  draws,
		boards: boards,
	}, nil
}

func part1(inputs []string) (int, error) {
	s, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	for _, num := range s.draws {
		for _, board := range s.boards {
			win := board.see(num)
			if !win {
				continue
			}
			return num * util.Sum(board.unseen()), nil
		}
	}

	return 0, errors.New("no wins seen")
}

func part2(inputs []string) (int, error) {
	s, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	winsLeft := len(s.boards)
	wins := make([]bool, len(s.boards))

	for _, num := range s.draws {
		for i, board := range s.boards {
			win := board.see(num)
			if win {
				if wins[i] {
					continue
				}
				wins[i] = true
				winsLeft--
				if winsLeft == 0 {
					return num * util.Sum(board.unseen()), nil
				}
			}
		}
	}

	return 0, errors.New("no wins seen")
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

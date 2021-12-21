package board

import (
	"fmt"
	"strconv"
)

type Position struct {
	Row int // rank
	Col int // file
}

func (p Position) String() string {
	// Given a position in row, column form we return a readable
	// string that represents a rank and file such as E4 or G7
	return fmt.Sprintf("%c%d", p.Col+'A', p.Row+1)
}

func ToPosition(s string) Position {
	// Given a string such as E4 or G7, we return a position in
	// row, column form to be used as a position in the board's
	parsed, _ := strconv.Atoi(string(s[1]))
	row := parsed - 1
	return Position{row, int(s[0] - 'A')}
}

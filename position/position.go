package position

import (
	"fmt"
	"strconv"

	"github.com/maaslalani/gambit/squares"
)

// Position represents a position on the board
type Position struct {
	// Row represents the row number of the cell in the board,
	// this can easily be converted to a human readable rank
	Row int // rank
	// Col represents the column number of the cell in the board,
	// this can easily be converted to a human readable file
	Col int // file
}

// String takes the current position and returns a human readable file and
// rank in chess notation
func (p Position) String() string {
	return ColumnToFile(p.Col) + RowToRank(p.Row)
}

// ToPosition reads a rank and file number and returns the corresponding
// position on the board's grid
func ToPosition(s squares.Square) Position {
	return Position{RankToRow(s[1]), FileToColumn(s[0])}
}

// RowToRank converts a row number to a human readable rank
func RowToRank(row int) string {
	return fmt.Sprintf("%d", row+1)
}

// RankToRow converts a human readable rank number to a board row number
func RankToRow(rank byte) int {
	parsed, _ := strconv.Atoi(string(rank))
	return parsed - 1
}

// ColumnToFile converts a column number to a human readable file
func ColumnToFile(column int) string {
	return fmt.Sprintf("%c", column+'A')
}

// FileToColumn converts a human readable file to a board column number
func FileToColumn(file byte) int {
	return int(file - 'A')
}

var NoPosition = Position{-1, -1}

package position

import (
	"fmt"
	"strconv"

	"github.com/maaslalani/gambit/board"
)

// colToFile returns the File given a column
func colToFile(col int) string {
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return fmt.Sprintf("%c", col+'a')
}

// rowToRank returns a Rank given a row
func rowToRank(row int) int {
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row + 1
}

// ToSquare returns the square position of a given row and
// column for display or checking legal moves.
//
// For example, ToSquare(0, 0) returns a1.
func ToSquare(row, col int) string {
	return colToFile(col) + strconv.Itoa(rowToRank(row))
}

package position

import (
	"fmt"
	"strconv"

	"github.com/maaslalani/gambit/board"
)

// colToFile returns the file given a column
func colToFile(col int) string {
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return fmt.Sprintf("%c", col+'a')
}

// rowToRank returns a rank given a row
func rowToRank(row int) int {
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row + 1
}

// ToSquare returns the square position (e.g. a1) of a given row and column
// (e.g. 0,0) for display or checking legal moves.
func ToSquare(row, col int, flipped bool) string {
	// If the board is flipped, the row and column are reversed
	// i.e. h becomes a and 8 becomes 1
	if flipped {
		col = board.LastCol - col
	} else {
		row = board.LastRow - row
	}
	return colToFile(col) + strconv.Itoa(rowToRank(row))
}

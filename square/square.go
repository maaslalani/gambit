package square

import (
	"strconv"

	"github.com/maaslalani/gambit/board"
)

// fileToCol returns column number for a given file rune
func fileToCol(file rune) int {
	col := int(file - 'a')
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return col
}

// rankToRow returns a row number for a given rank
func rankToRow(rank int) int {
	row := rank - 1
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row
}

// ToPosition takes a square string and returns the
// corresponding row and column for compatibility with the grid
// (8x8 matrix).
//
// For example: "a1" returns (0, 0) and "h8" returns (7, 7).
func ToPosition(square string) (int, int) {
	col := fileToCol(rune(square[0]))
	row, _ := strconv.Atoi(string(square[1]))
	row = rankToRow(row)
	return col, row
}

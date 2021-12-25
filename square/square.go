package square

import (
	"strconv"

	"github.com/maaslalani/gambit/board"
)

// fileToCol returns column number (e.g. 0) for a given file (e.g. 'a').
func fileToCol(file rune) int {
	col := int(file - 'a')
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return col
}

// rankToRow returns a row number (e.g. 0) for a given rank (e.g. 1).
func rankToRow(rank int) int {
	row := rank - 1
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row
}

// ToPosition takes a square (e.g. a1) and returns the corresponding row and
// column (e.g. 0,0) for compatibility with the grid (8x8 matrix).
func ToPosition(square string) (int, int) {
	col := fileToCol(rune(square[0]))
	row, _ := strconv.Atoi(string(square[1]))
	row = rankToRow(row)
	return col, row
}

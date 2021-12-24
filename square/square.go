package square

import (
	"strconv"

	"github.com/maaslalani/gambit/board"
)

func fileToCol(file rune) int {
	col := int(file - 'a')
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return col
}

func rankToRow(rank int) int {
	row := rank - 1
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row
}

func ToPosition(square string) (int, int) {
	col := fileToCol(rune(square[0]))
	row, _ := strconv.Atoi(string(square[1]))
	row = rankToRow(row)
	return col, row
}

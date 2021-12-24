package position

import (
	"fmt"
	"strconv"

	"github.com/maaslalani/gambit/board"
)

func colToFile(col int) string {
	if col < board.FirstCol {
		col = board.FirstCol
	} else if col > board.LastCol {
		col = board.LastCol
	}
	return fmt.Sprintf("%c", col+'a')
}

func rowToRank(row int) int {
	if row < board.FirstRow {
		row = board.FirstRow
	} else if row > board.LastRow {
		row = board.LastRow
	}
	return row + 1
}

func ToSquare(row, col int) string {
	return colToFile(col) + strconv.Itoa(rowToRank(row))
}

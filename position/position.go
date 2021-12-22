package position

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
	return fmt.Sprintf("%c%d", ColumnToFile(p.Col), RowToRank(p.Row))
}

func ToPosition(s string) Position {
	// Given a string such as E4 or G7, we return a position in
	// row, column form to be used as a position in the board's
	return Position{RankToRow(s[1]), FileToColumn(s[0])}
}

func RowToRank(row int) int {
	return row + 1
}

func RankToRow(rank byte) int {
	parsed, _ := strconv.Atoi(string(rank))
	return parsed - 1
}

func ColumnToFile(column int) string {
	return fmt.Sprintf("%c", column+'A')
}

func FileToColumn(file byte) int {
	return int(file - 'A')
}

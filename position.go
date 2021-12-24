package main

import (
	"fmt"
	"strconv"
)

func ColToFile(col int) string {
	if col < firstCol {
		col = firstCol
	} else if col > lastCol {
		col = lastCol
	}
	return fmt.Sprintf("%c", col+'a')
}

func FileToCol(file rune) int {
	col := int(file - 'a')
	if col < firstCol {
		col = firstCol
	} else if col > lastCol {
		col = lastCol
	}
	return col
}

func RowToRank(row int) int {
	if row < firstRow {
		row = firstRow
	} else if row > lastRow {
		row = lastRow
	}
	return row + 1
}

func RankToRow(rank int) int {
	row := rank - 1
	if row < firstRow {
		row = firstRow
	} else if row > lastRow {
		row = lastRow
	}
	return row
}

func SquareToPosition(square string) (int, int) {
	col := FileToCol(rune(square[0]))
	row, _ := strconv.Atoi(string(square[1]))
	row = RankToRow(row)
	return col, row
}

func PositionToSquare(row, col int) string {
	return ColToFile(col) + strconv.Itoa(RowToRank(row))
}

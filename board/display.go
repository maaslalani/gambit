package board

import (
	"fmt"
	"strings"

	"github.com/maaslalani/gambit/position"
)

var (
	border = []string{"┌", "┬", "┐", "├", "┼", "┤", "└", "┴", "┘"}

	borderBottomOffset = 6
	borderMiddleOffset = 3
	borderTopOffset    = 0

	files = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	ranks = []int{7, 6, 5, 4, 3, 2, 1, 0}
)

const (
	vertical   = "│"
	horizontal = "─"
	marginLeft = "    "
)

// String prints the board in a human readable format.
// The left and bottom sides have labels for ranks and files respectively
// All pieces are surrounded with borders.
func (b Board) String() string {
	var s string

	if b.reversed {
		ranks = reverse(ranks)
	}

	for r, row := range ranks {

		if isFirstRow(r) {
			s += buildRow(borderTopOffset) + "\n"
		}

		for c, cell := range b.grid[row] {
			if isFirstColumn(c) {
				s += fmt.Sprintf(" %d  ", position.RowToRank(row))
			}

			s += fmt.Sprintf("%s %s ", vertical, cell)

			if isLastColumn(c) {
				s += vertical
			}
		}

		if !isLastRow(r) {
			s += buildRow(borderMiddleOffset)
		} else {
			s += buildRow(borderBottomOffset)
			s += "\n      "
			s += strings.Join(files, "   ")
		}
		s += "\n"
	}

	return s
}

// isLastRow returns whether or not the given row is the last row of the
// board based on the number of ranks on the board.
func isLastRow(i int) bool {
	return i == len(ranks)-1
}

// isLastColumn returns whether or not the given column is the last column of
// the board based on the number of files on the board.
func isLastColumn(i int) bool {
	return i == len(files)-1
}

// isFirstRow returns whether or not the given row is the first row
func isFirstRow(i int) bool {
	return i == 0
}

// isFirstColumn returns whether or not the given column is the first column
func isFirstColumn(i int) bool {
	return i == 0
}

// buildRow builds a row string based on the given borders for the left and
// right side and correctly pads the middle with the given character adjusted
// to the number of rows on the board.
func buildRow(borderOffset int) string {
	var left, middle, right = border[borderOffset], border[borderOffset+1], border[borderOffset+2]
	var row []string
	row = append(row, left)
	for i := 0; i < len(ranks)-1; i++ {
		row = append(row, middle)
	}
	row = append(row, right)
	return fmt.Sprintf("\n    %s", strings.Join(row, horizontal+horizontal+horizontal))
}

// reverse reverses the given slice of ints.
func reverse(ranks []int) []int {
	var reversed []int
	for i := len(ranks) - 1; i >= 0; i-- {
		reversed = append(reversed, ranks[i])
	}
	return reversed
}

package board

import (
	"strings"

	"github.com/maaslalani/gambit/position"
	"github.com/maaslalani/gambit/style"
)

const (
	top int = iota
	middle
	bottom
)

var (
	border = map[int][]rune{
		top:    []rune("┌┬┐"),
		middle: []rune("├┼┤"),
		bottom: []rune("└┴┘"),
	}

	files = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	ranks = []int{7, 6, 5, 4, 3, 2, 1, 0}
)

const (
	vertical   = "│"
	horizontal = "─"
)

// String prints the board in a human readable format.
// The left and bottom sides have labels for ranks and files respectively
// All pieces are surrounded with borders.
func (b Board) String() string {
	var s string

	if b.Reversed {
		ranks = reverse(ranks)
	}

	for r, row := range ranks {

		if isFirstRow(r) {
			s += buildRow(border[top]) + "\n"
		}

		for c, cell := range b.Grid[row] {
			// Rank labels
			if isFirstColumn(c) {
				s += " " + style.Faint.Render(position.RowToRank(row)) + "  "
			}

			s += vertical + " "
			cellStyle := cell.Style()
			if b.Selected != position.NoPosition && row == b.Selected.Row && c == b.Selected.Col {
				cellStyle = style.Selected
			}
			s += cellStyle.Render(cell.String()) + " "

			if isLastColumn(c) {
				s += vertical
			}
		}

		if !isLastRow(r) {
			s += buildRow(border[middle]) + "\n"
		} else {
			s += buildRow(border[bottom]) + "\n"
			// File labels
			s += "      " + style.Faint.Render(strings.Join(files, "   ")) + "\n"
		}
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
func buildRow(border []rune) string {
	var row [9]string
	row[0] = string(border[0])
	row[8] = string(border[2])
	for i := 1; i < len(ranks); i++ {
		row[i] = string(border[1])
	}
	return "\n    " + strings.Join(row[:], strings.Repeat(horizontal, 3))
}

// reverse reverses the given slice of ints.
func reverse(ranks []int) []int {
	var reversed []int
	for i := len(ranks) - 1; i >= 0; i-- {
		reversed = append(reversed, ranks[i])
	}
	return reversed
}

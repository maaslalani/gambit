package border

import (
	"strings"

	"github.com/maaslalani/gambit/board"
)

const (
	cellHeight = 2
	cellWidth  = 4
	marginLeft = 3
	marginTop  = 1

	Vertical   = "│"
	Horizontal = "─"
)

func Cell(x, y int) (int, int) {
	col := (x - marginLeft) / cellWidth
	row := board.LastRow - (y-marginTop)/cellHeight
	return col, row
}

// withMarginLeft returns a string with a prepended left margin
func withMarginLeft(s string) string {
	return strings.Repeat(" ", marginLeft) + s
}

// Build returns a string with a border for a given row (top, middle, bottom)
func Build(left, middle, right string) string {
	border := left + Horizontal + strings.Repeat(Horizontal+Horizontal+middle+Horizontal, board.LastRow)
	border += Horizontal + Horizontal + right + "\n"
	return withMarginLeft(border)
}

func Top() string {
	return Build("┌", "┬", "┐")
}

func Middle() string {
	return Build("├", "┼", "┤")
}

func Bottom() string {
	return Build("└", "┴", "┘")
}

func BottomLabels() string {
	return withMarginLeft("  A   B   C   D   E   F   G   H\n")
}

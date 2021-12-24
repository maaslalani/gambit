package border

import (
	"strings"

	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/position"
)

const (
	// cellHeight represents how many rows are in a cell
	cellHeight = 2
	// cellWidth represents how many columns are in a cell
	cellWidth = 4

	// marginLeft and marginTop represent how much to offset the
	// chess board from the top left of the terminal to account for
	// padding and rank labels
	marginLeft = 3
	marginTop  = 1

	Vertical   = "│"
	Horizontal = "─"
)

// Cell returns the square that was clicked based on mouse
// coordinates adjusting for margins and cell dimensions.
func Cell(x, y int) string {
	col := (x - marginLeft) / cellWidth
	row := board.LastRow - (y-marginTop)/cellHeight
	return position.ToSquare(row, col)

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

// Top returns a built border with the top row
func Top() string {
	return Build("┌", "┬", "┐")
}

// Middle returns a built border with the middle row
func Middle() string {
	return Build("├", "┼", "┤")
}

// Bottom returns a built border with the bottom row
func Bottom() string {
	return Build("└", "┴", "┘")
}

// BottomLabels returns the labels for the files
func BottomLabels() string {
	return withMarginLeft("  A   B   C   D   E   F   G   H\n")
}

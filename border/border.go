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

	// marginLeft and marginTop represent the offset of the chess
	// board from the top left of the terminal window. This is to
	// account for padding and rank labels
	marginLeft = 3
	marginTop  = 1

	Vertical   = "│"
	Horizontal = "─"
)

// Cell returns the square that was clicked based on mouse coordinates adjusted
// for margins and cell dimensions.
func Cell(x, y int, flipped bool) string {
	col := (x - marginLeft) / cellWidth
	row := (y - marginTop) / cellHeight
	if !flipped {
		// Careful: `flipped` is a bit strange here.
		//
		// `flipped` represents whether the _chess board_ is flipped.
		// Normally, the order of the rows will be from 8 (top) to 1 (bottom).
		// If the board is flipped (i.e. black's turn) then the order of the rows
		// would be how one would normally print of the rows: 1 (top) to 8 (bottom).
		//
		// In other words, flipped is flipped. This trade-off is for consistency
		// with the other parts of the code base, namely the model.flipped property
		row = board.LastRow - row
	}
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

package border

import (
	"fmt"
	"strings"

	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/position"
	"github.com/maaslalani/gambit/style"
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
	return position.ToSquare(row, col, flipped)
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

func topWithTheme(theme style.Theme) string {
	var s strings.Builder
	// top-left square is light
	isLightSquare := true
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < cellWidth; j++ {
			s.WriteString(theme.Fg("\u2583", isLightSquare))
		}
		isLightSquare = !isLightSquare
	}
	s.WriteRune('\n')
	return withMarginLeft(s.String())
}

// Top returns a built border with the top row
func Top(theme style.Theme) string {
	if !theme.IsValid() {
		return Build("┌", "┬", "┐")
	}
	return topWithTheme(theme)
}

func middleWithTheme(theme style.Theme, firstSquareIsLight bool) string {
	var s strings.Builder
	isLightSquare := firstSquareIsLight
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < cellWidth; j++ {
			s.WriteString(theme.Border("\u2580", isLightSquare))
		}
		isLightSquare = !isLightSquare
	}
	s.WriteRune('\n')
	return withMarginLeft(s.String())
}

// Middle returns a built border with the middle row
func Middle(theme style.Theme, firstSquareIsLight bool) string {
	if !theme.IsValid() {
		return Build("├", "┼", "┤")
	}
	return middleWithTheme(theme, firstSquareIsLight)
}

// bottomWithTheme returns built border with the bottom row
func bottomWithTheme(theme style.Theme) string {
	var s strings.Builder
	// bottom-left square is dark
	isLightSquare := false
	for i := 0; i < board.Cols; i++ {
		for j := 0; j < cellWidth; j++ {
			s.WriteString(theme.Fg("\u2580", isLightSquare))
		}
		isLightSquare = !isLightSquare
	}
	s.WriteRune('\n')
	return withMarginLeft(s.String())
}

// Bottom returns a built border with the bottom row
func Bottom(theme style.Theme) string {
	if !theme.IsValid() {
		return Build("└", "┴", "┘")
	}
	return bottomWithTheme(theme)
}

// BottomLabels returns the labels for the files
func BottomLabels(flipped bool) string {
	labels := ""
	for i := 0; i < board.Cols; i++ {
		c := i
		if flipped {
			c = board.LastCol - i
		}
		labels += fmt.Sprintf("%c", c+'A')
		if i != board.LastCol {
			labels += withMarginLeft("")
		}
	}
	return withMarginLeft(fmt.Sprintf("  %s\n", labels))
}

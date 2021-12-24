package main

import (
	"strings"

	"github.com/maaslalani/gambit/board"
)

const (
	cellHeight = 2
	cellWidth  = 4
	marginLeft = 3
	marginTop  = 1

	vertical   = "│"
	horizontal = "─"
)

// withMarginLeft returns a string with a prepended left margin
func withMarginLeft(s string) string {
	return strings.Repeat(" ", marginLeft) + s
}

// buildBorder returns a string with a border for a given row (top, middle, bottom)
func buildBorder(left, middle, right string) string {
	border := left + horizontal + strings.Repeat(horizontal+horizontal+middle+horizontal, board.LastRow)
	border += horizontal + horizontal + right + "\n"
	return withMarginLeft(border)
}

func topBorder() string {
	return buildBorder("┌", "┬", "┐")
}

func middleBorder() string {
	return buildBorder("├", "┼", "┤")
}

func bottomBorder() string {
	return buildBorder("└", "┴", "┘")
}

func bottomLabels() string {
	return withMarginLeft("  A   B   C   D   E   F   G   H\n")
}

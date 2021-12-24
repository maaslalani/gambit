package fen

import (
	"fmt"
	"strings"
)

// Tokens returns the tokens of a FEN string
func Tokens(fen string) []string {
	return strings.Split(fen, " ")
}

// Ranks returns a slice of ranks from the FEN string
func Ranks(fen string) []string {
	return strings.Split(Tokens(fen)[0], "/")
}

// Grid returns a 8x8 grid of the board represented by the FEN string
func Grid(fen string) [8][8]string {
	var grid [8][8]string
	for r, rank := range Ranks(fen) {
		var row [8]string
		c := 0
		for _, col := range rank {
			skip := 1
			if isNumeric(col) {
				skip = runeToInt(col)
			} else {
				row[c] = fmt.Sprintf("%c", col)
			}
			c += skip
		}
		grid[r] = row
	}
	return grid
}

// isNumeric returns true if the current rune is a number
func isNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

// runeToInt converts a rune to an int
func runeToInt(r rune) int {
	return int(r - '0')
}

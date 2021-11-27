package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	dimensions = 8
	firstRow   = 0
	firstCol   = 0
	lastRow    = dimensions - 1
	lastCol    = dimensions - 1
)

const (
	header       = "\n    ┌───┬───┬───┬───┬───┬───┬───┬───┐ \n"
	middle       = "\n    ├───┼───┼───┼───┼───┼───┼───┼───┤ \n"
	footer       = "\n    └───┴───┴───┴───┴───┴───┴───┴───┘ \n"
	footerLabels = "\n      A   B   C   D   E   F   G   H   \n"
	divider      = " │ "
)

var faintStyle = lipgloss.NewStyle().Faint(true)

type Board struct {
	Players []Player
	Grid    [dimensions][dimensions]string
}

func (b *Board) Draw() {
	for _, player := range b.Players {
		for _, piece := range player.Pieces {
			b.Grid[piece.Position[0]-1][piece.Position[1]-1] = piece.String()
		}
	}
}

func (b *Board) String() string {
	var s = header
	for row := firstRow; row < dimensions; row++ {
		for col := firstCol; col < dimensions; col++ {
			if col == firstCol {
				s += faintStyle.Render(fmt.Sprintf(" %d ", dimensions-row))
			}
			s += divider
			if b.Grid[row][col] == "" {
				s += " "
			} else {
				s += b.Grid[row][col]
			}
			if col == lastCol {
				s += divider
			}
		}
		if row != lastRow {
			s += middle
		}
	}
	return s + footer + faintStyle.Render(footerLabels)
}

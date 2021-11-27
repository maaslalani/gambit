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
	flipped bool
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

	var ranks []int

	if b.flipped {
		ranks = []int{0, 1, 2, 3, 4, 5, 6, 7}
	} else {
		ranks = []int{7, 6, 5, 4, 3, 2, 1, 0}
	}

	for row, rank := range ranks {
		for col := firstCol; col < dimensions; col++ {
			if col == firstCol {
				s += faintStyle.Render(fmt.Sprintf(" %d ", rank+1))
			}
			s += divider
			if b.Grid[rank][col] == "" {
				s += " "
			} else {
				s += b.Grid[rank][col]
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

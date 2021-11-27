package main

import (
	"fmt"
	"strings"
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

type Board struct {
	Players []Player
	Grid    [dimensions][dimensions]*Piece
	flipped bool
}

func (b *Board) Move(from, to position) {
	p := b.Grid[from[0]-1][from[1]-1]
	if p == nil {
		return
	}
	b.Grid[from[0]-1][from[1]-1] = nil
	b.Grid[to[0]-1][to[1]-1] = p
	p.Position = to
}

func (b *Board) Draw() {
	for _, player := range b.Players {
		for _, piece := range player.Pieces {
			p := piece
			b.Grid[piece.Position[0]-1][piece.Position[1]-1] = p
		}
	}
}

func FileToColumn(file string) int {
	if len(file) < 1 {
		return 0
	}
	return int([]rune(strings.ToUpper(file))[0] - 64)
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
		for file := firstCol; file < dimensions; file++ {
			if file == firstCol {
				s += faintStyle.Render(fmt.Sprintf(" %d ", rank+1))
			}
			s += divider
			if b.Grid[rank][file] == nil {
				s += " "
			} else {
				s += b.Grid[rank][file].String()
			}
			if file == lastCol {
				s += divider
			}
		}
		if row != lastRow {
			s += middle
		}
	}
	return s + footer + faintStyle.Render(footerLabels)
}

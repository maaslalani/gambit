package main

const dimensions = 8
const firstRow = 0
const firstCol = 0
const lastRow = dimensions - 1
const lastCol = dimensions - 1

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
	var s string
	for row := 0; row < dimensions; row++ {
		if row == firstRow {
			s += " ┌───┬───┬───┬───┬───┬───┬───┬───┐ \n"
		}
		for col := 0; col < dimensions; col++ {
			s += " │ "
			if b.Grid[row][col] == "" {
				s += " "
			} else {
				s += b.Grid[row][col]
			}
			if col == lastCol {
				s += " │ "
			}
		}
		if row == lastRow {
			s += "\n └───┴───┴───┴───┴───┴───┴───┴───┘ \n"
		} else {
			s += "\n ├───┼───┼───┼───┼───┼───┼───┼───┤ \n"
		}
	}
	return s
}

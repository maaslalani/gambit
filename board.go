package main

type Board struct {
	Players []Player
	Grid    [8][8]string
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
	for row := 0; row < 8; row++ {
		if row == 0 {
			s += " ┌───┬───┬───┬───┬───┬───┬───┬───┐ \n"
		}
		for col := 0; col < 8; col++ {
			s += " │ "
			if b.Grid[row][col] == "" {
				s += " "
			} else {
				s += b.Grid[row][col]
			}
			if col == 7 {
				s += " │ "
			}
		}
		if row == 7 {
			s += "\n └───┴───┴───┴───┴───┴───┴───┴───┘ \n"
		} else {
			s += "\n ├───┼───┼───┼───┼───┼───┼───┼───┤ \n"
		}
	}
	return s
}

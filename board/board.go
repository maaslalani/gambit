package board

import (
	"github.com/maaslalani/gambit/piece"
	"github.com/maaslalani/gambit/position"
)

type Board struct {
	// The board is represented as a 2D array of cells.
	// The first index is the row, the second is the column.
	Grid     [8][8]piece.Piece
	Reversed bool
	Selected position.Position
	Turn     piece.Color
}

func New() Board {
	ep := piece.Empty()
	er := [8]piece.Piece{ep, ep, ep, ep, ep, ep, ep, ep}
	return Board{
		Grid:     [8][8]piece.Piece{er, er, er, er, er, er, er, er},
		Turn:     piece.White,
		Selected: position.NoPosition,
	}
}

func (b *Board) At(p position.Position) piece.Piece {
	return b.Grid[p.Row][p.Col]
}

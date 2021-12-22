package board

import (
	"github.com/maaslalani/gambit/piece"
)

type Board struct {
	// The board is represented as a 2D array of cells.
	// The first index is the row, the second is the column.
	Grid     [8][8]piece.Piece
	Reversed bool
	Turn     piece.Color
}

func New() Board {
	ep := piece.Empty()
	er := [8]piece.Piece{ep, ep, ep, ep, ep, ep, ep, ep}
	return Board{
		Grid: [8][8]piece.Piece{er, er, er, er, er, er, er, er},
		Turn: piece.White,
	}
}

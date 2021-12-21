package board

import "github.com/maaslalani/gambit/piece"

type Board struct {
	// The board is represented as a 2D array of cells.
	// The first index is the row, the second is the column.
	grid [8][8]*piece.Piece
}
package board

import (
	"github.com/maaslalani/gambit/piece"
	"github.com/maaslalani/gambit/position"
	"github.com/maaslalani/gambit/squares"
)

type Move struct {
	From squares.Square
	To   squares.Square
}

func (b *Board) Move(m Move) {
	f := position.ToPosition(m.From)
	t := position.ToPosition(m.To)
	b.Grid[t.Row][t.Col] = b.Grid[f.Row][f.Col]
	b.Grid[f.Row][f.Col] = piece.Empty()
}

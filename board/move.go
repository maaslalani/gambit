package board

import (
	"github.com/maaslalani/gambit/piece"
	"github.com/maaslalani/gambit/position"
)

func (b *Board) Move(from, to string) {
	f := position.ToPosition(from)
	t := position.ToPosition(to)
	b.Grid[t.Row][t.Col] = b.Grid[f.Row][f.Col]
	b.Grid[f.Row][f.Col] = piece.Empty()
}

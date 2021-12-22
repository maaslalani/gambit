package piece

import (
	"github.com/maaslalani/gambit/color"
)

type Piece struct {
	Type  Type
	Color color.Color
}

func Empty() Piece {
	return Piece{None, color.None}
}

func (p Piece) String() string {
	return Display[p.Type]
}

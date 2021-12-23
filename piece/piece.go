package piece

import "github.com/maaslalani/gambit/style"

type Piece struct {
	Type  Type
	Color Color
}

func Empty() Piece {
	return Piece{NoType, NoColor}
}

func (p Piece) String() string {
	if p.Color == White {
		return style.White.Render(Display[p.Type])
	} else {
		return style.Black.Render(Display[p.Type])
	}
}

var (
	BW = Piece{Bishop, White}
	KW = Piece{King, White}
	NW = Piece{Knight, White}
	PW = Piece{Pawn, White}
	QW = Piece{Queen, White}
	RW = Piece{Rook, White}
)

var (
	BB = Piece{Bishop, Black}
	KB = Piece{King, Black}
	NB = Piece{Knight, Black}
	PB = Piece{Pawn, Black}
	QB = Piece{Queen, Black}
	RB = Piece{Rook, Black}
)

var OO = Piece{NoType, NoColor}
var EmptyPiece = Piece{NoType, NoColor}

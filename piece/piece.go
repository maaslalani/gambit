package piece

type Piece struct {
	Type  Type
	Color Color
}

func Empty() Piece {
	return Piece{NoType, NoColor}
}

func (p Piece) String() string {
	return Display[p.Type]
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

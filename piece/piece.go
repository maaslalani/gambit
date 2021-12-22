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

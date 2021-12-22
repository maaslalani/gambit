package piece

type Type string

func (t Type) String() string {
	return string(t)
}

const (
	Pawn   Type = "P"
	Knight Type = "N"
	Bishop Type = "B"
	Rook   Type = "R"
	Queen  Type = "Q"
	King   Type = "K"
	NoType Type = " "
)

var Display = map[Type]string{
	Bishop: "♝",
	King:   "♚",
	Knight: "♞",
	Pawn:   "♟",
	Queen:  "♛",
	Rook:   "♜",
	NoType: " ",
}

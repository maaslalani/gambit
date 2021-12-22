package piece

type Type string

const (
	Pawn   Type = "P"
	Knight Type = "N"
	Bishop Type = "B"
	Rook   Type = "R"
	Queen  Type = "Q"
	King   Type = "K"
	None   Type = " "
)

var Display = map[Type]string{
	Bishop: "♝",
	King:   "♚",
	Knight: "♞",
	Pawn:   "♟",
	Queen:  "♛",
	Rook:   "♜",
	None:   " ",
}

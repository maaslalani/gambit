package piece

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/gambit/style"
)

// Color of pieces
type Color string

const (
	White   Color = "w"
	Black   Color = "b"
	NoColor Color = ""
)

// Types of pieces
type Type string

const (
	Pawn   Type = "P"
	Knight Type = "N"
	Bishop Type = "B"
	Rook   Type = "R"
	Queen  Type = "Q"
	King   Type = "K"
	NoType Type = ""
)

func (t Type) String() string {
	return string(t)
}

// Pieces
type Piece struct {
	Type  Type
	Color Color
}

func (p Piece) String() string {
	return Display[p.Type]
}

func Empty() Piece {
	return Piece{NoType, NoColor}
}

var Display = map[Type]string{
	Bishop: "♝",
	King:   "♚",
	Knight: "♞",
	Pawn:   "♟",
	Queen:  "♛",
	Rook:   "♜",
	NoType: " ",
}

// ToFen converts a piece into its FEN representation
// i.e. White Knight -> "N"
//      Black Bishop -> "b"
// if a piece is empty, it returns an empty string
func (p Piece) ToFen() string {
	t := string(p.Type)
	if p.Color == Black {
		return strings.ToLower(t)
	}
	return t
}

// FromFen converts a FEN representation of a piece into a
// piece. Reverses ToFen.
func FromFen(fen string) Piece {
	u := strings.ToUpper(fen)
	t := Type(u)
	var c Color
	if u == fen {
		c = White
	} else {
		c = Black
	}
	return Piece{Type: t, Color: c}
}

func (p Piece) Style() lipgloss.Style {
	if p.Color == White {
		return style.White
	} else {
		return style.Black
	}
}

var EmptyPiece = Piece{NoType, NoColor}

var (
	BB = Piece{Bishop, Black}
	BW = Piece{Bishop, White}
	KB = Piece{King, Black}
	KW = Piece{King, White}
	NB = Piece{Knight, Black}
	NW = Piece{Knight, White}
	PB = Piece{Pawn, Black}
	PW = Piece{Pawn, White}
	QB = Piece{Queen, Black}
	QW = Piece{Queen, White}
	RB = Piece{Rook, Black}
	RW = Piece{Rook, White}
	OO = Piece{NoType, NoColor}
)

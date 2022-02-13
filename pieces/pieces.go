package pieces

import "strings"

// Piece represents a chess piece.
type Piece string

// display maps pieces from their FEN representations to their ASCII
// representations for a more human readable experience.
var display = map[Piece]string{
	"":  " ",
	"B": "♝",
	"K": "♚",
	"N": "♞",
	"P": "♟",
	"Q": "♛",
	"R": "♜",
	"b": "♗",
	"k": "♔",
	"n": "♘",
	"p": "♙",
	"q": "♕",
	"r": "♖",
}

// IsWhite returns true if the piece is white.
func (p Piece) IsWhite() bool {
	s := p.String()
	return strings.ToUpper(s) == s
}

// IsBlack returns true if the piece is black.
func (p Piece) IsBlack() bool {
	s := p.String()
	return strings.ToLower(s) == s
}

// Display returns the ASCII representation of the piece.
func (p Piece) Display() string {
	return display[p]
}

// String implements the stringer interface.
func (p Piece) String() string {
	return string(p)
}

// IsKing returns true if the piece is a king.
func (p Piece) IsKing() bool {
	return p == "K" || p == "k"
}

// IsPawn returns true if the piece is a pawn.
func (p Piece) IsPawn() bool {
	return p == "P" || p == "p"
}

// IsRook returns true if the piece is a rook.
func (p Piece) IsRook() bool {
	return p == "R" || p == "r"
}

// IsBishop returns true if the piece is a bishop.
func (p Piece) IsBishop() bool {
	return p == "B" || p == "b"
}

// IsKnight returns true if the piece is a knight.
func (p Piece) IsKnight() bool {
	return p == "N" || p == "n"
}

// IsQueen returns true if the piece is a queen.
func (p Piece) IsQueen() bool {
	return p == "Q" || p == "q"
}

// IsEmpty returns true if the piece is empty.
func (p Piece) IsEmpty() bool {
	return p == ""
}

// ToPieces converts a slice of FEN string to a slice of pieces.
func ToPieces(r [8]string) [8]Piece {
	var p [8]Piece
	for i, s := range r {
		p[i] = Piece(s)
	}
	return p
}

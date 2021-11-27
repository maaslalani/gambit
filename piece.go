package main

import "fmt"

type piece string

const (
	Pawn   piece = "pawn"
	Bishop piece = "bishop"
	Knight piece = "knight"
	Rook   piece = "rook"
	Queen  piece = "queen"
	King   piece = "king"
)

type color string

const (
	White color = "white"
	Black color = "black"
)

var pieces = map[color]map[piece]string{
	Black: {
		Pawn:   "♟",
		Bishop: "♝",
		Rook:   "♜",
		Knight: "♞",
		Queen:  "♛",
		King:   "♚",
	},
	White: {
		Pawn:   "♙",
		Bishop: "♗",
		Rook:   "♖",
		Knight: "♘",
		Queen:  "♕",
		King:   "♔",
	},
}

// array of [row, column]
// e.g. [1, 4]
type position [2]int

func (p position) String() string {
	return p.Row() + p.Col()
}

func (p position) Row() string {
	return string(rune('A' - 1 + p[0]))
}

func (p position) Col() string {
	return fmt.Sprint(p[1])
}

type Piece struct {
	Type     piece
	Color    color
	Position position
	Active   bool
}

func (p Piece) String() string {
	return pieces[p.Color][p.Type]
}

func NewPiece(piece piece, position position, color color) Piece {
	return Piece{
		Type:     piece,
		Position: position,
		Color:    color,
	}
}

func InitialPieces(color color) []Piece {
	var pieces []Piece
	var (
		pawnRank int
		backRank int
	)

	if color == White {
		backRank, pawnRank = 1, 2
	} else {
		backRank, pawnRank = 8, 7
	}

	for i := 1; i <= 8; i++ {
		pieces = append(pieces, NewPiece(Pawn, position{pawnRank, i}, color))
	}

	for i, p := range []piece{Rook, Knight, Bishop, Queen, King, Knight, Bishop, Rook} {
		pieces = append(pieces, NewPiece(p, position{backRank, i}, color))
	}

	return pieces
}

func BlackPieces() []Piece {
	return InitialPieces(Black)
}

func WhitePieces() []Piece {
	return InitialPieces(White)
}

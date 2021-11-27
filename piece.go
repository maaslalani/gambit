package main

import (
	"fmt"
)

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

var pieces = map[piece]string{
	Pawn:   "♟",
	Bishop: "♝",
	Rook:   "♜",
	Knight: "♞",
	Queen:  "♛",
	King:   "♚",
}

// array of [row,  column]
// array of [rank, file]
// e.g. [1, 4] (D1)
type position [2]int

func (p position) String() string {
	return p.File() + p.Rank()
}

func (p position) File() string {
	return string(rune('A' - 1 + p[1]))
}

func (p position) Rank() string {
	return fmt.Sprint(p[0])
}

type Piece struct {
	Type     piece
	Color    color
	Position position
	Selected bool
}

func (p Piece) String() string {
	if p.Selected {
		return activeStyle.Render(pieces[p.Type])
	}
	if p.Color == Black {
		return faintStyle.Render(pieces[p.Type])
	}
	return pieces[p.Type]
}

func NewPiece(piece piece, position position, color color) *Piece {
	return &Piece{
		Type:     piece,
		Position: position,
		Color:    color,
	}
}

func InitialPieces(color color) []*Piece {
	var pieces []*Piece
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
		pieces = append(pieces, NewPiece(p, position{backRank, i + 1}, color))
	}

	return pieces
}

func BlackPieces() []*Piece {
	return InitialPieces(Black)
}

func WhitePieces() []*Piece {
	return InitialPieces(White)
}

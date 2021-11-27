package main

import (
	"fmt"
	"strconv"
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

func ToPosition(square string) position {
	if len(square) != 2 {
		return position{0, 0}
	}

	rank, _ := strconv.Atoi(string(square[1]))
	file := FileToColumn(string(square[0]))
	return position{rank, file}
}

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

package main

type Player struct {
	Points   int
	Pieces   []*Piece
	Captured []*Piece
}

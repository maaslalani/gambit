package main

import "fmt"

type Board struct {
	Players []Player
}

func (b Board) String() string {
	var s string
	for _, player := range b.Players {
		for _, piece := range player.Pieces {
			s += fmt.Sprint(piece)
		}
	}
	return s
}

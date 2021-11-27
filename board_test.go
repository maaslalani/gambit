package main

import "testing"

func TestSerialize(t *testing.T) {
	positions := `RNBQKB.R
PPPP.PPP
.....N..
....P...
....p...
..n.....
pppp.ppp
r.bqkbnr`

	board := Deserialize(positions)
	serialized := Serialize(board)

	if positions != serialized {
		t.Log("want: " + positions)
		t.Log("got: " + serialized)
		t.Fatal("Board was not equivalent to deserialization and serialization")
	}
}

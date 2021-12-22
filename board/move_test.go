package board

import (
	"testing"

	. "github.com/maaslalani/gambit/squares"
)

func TestMove(t *testing.T) {
	initial := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	tests := []struct {
		board    string
		moves    []Move
		expected string
	}{
		{initial, []Move{{E2, E4}}, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"},
		{initial, []Move{{A7, A6}}, "rnbqkbnr/1ppppppp/p7/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"},
		{initial, []Move{{G1, F3}}, "rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R w KQkq - 0 1"},
		{initial, []Move{{E2, E4}, {C7, C5}, {G1, F3}}, "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 0 1"},
	}

	for _, tc := range tests {
		b, _ := FromFen(tc.board)
		for _, move := range tc.moves {
			b.Move(move)
		}
		if b.ToFen() != tc.expected {
			t.Errorf("\nwant %s\ngot  %s", tc.expected, b.ToFen())
			// t.Log(b)
		}
	}
}

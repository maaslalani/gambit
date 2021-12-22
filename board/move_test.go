package board

import (
	"testing"

	. "github.com/maaslalani/gambit/squares"
)

func TestMove(t *testing.T) {
	initial := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	expected := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
	b, _ := FromFen(initial)
	b.reversed = true
	b.Move(E2, E4)

	if b.ToFen() != expected {
		t.Errorf("Expected %s, got %s", expected, b.ToFen())
	}
}

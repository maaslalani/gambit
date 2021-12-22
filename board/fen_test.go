package board_test

import (
	"testing"

	"github.com/maaslalani/gambit/board"
	. "github.com/maaslalani/gambit/piece"
)

func TestFen(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	grid := [8][8]Piece{
		{RW, NW, BW, QW, KW, BW, NW, RW},
		{PW, PW, PW, PW, PW, PW, PW, PW},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{PB, PB, PB, PB, PB, PB, PB, PB},
		{RB, NB, BB, QB, KB, BB, NB, RB},
	}
	b := board.Board{
		Grid: grid,
		Turn: White,
	}

	if b.ToFen() != fen {
		t.Errorf("ToFen()\nActual = %s\nTarget = %s", b.ToFen(), fen)
	}

	fromFen, _ := board.FromFen(fen)
	if fromFen.ToFen() != fen {
		t.Errorf("FromFen()\nActual = %s\nTarget = %s", fromFen.ToFen(), fen)
	}
}

func TestFenMoved(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
	grid := [8][8]Piece{
		{RW, NW, BW, QW, KW, BW, NW, RW},
		{PW, PW, PW, PW, OO, PW, PW, PW},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{OO, OO, OO, OO, PW, OO, OO, OO},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{OO, OO, OO, OO, OO, OO, OO, OO},
		{PB, PB, PB, PB, PB, PB, PB, PB},
		{RB, NB, BB, QB, KB, BB, NB, RB},
	}
	b := board.Board{
		Grid: grid,
		Turn: White,
	}

	if b.ToFen() != fen {
		t.Errorf("ToFen()\nActual = %s\nTarget = %s", b.ToFen(), fen)
	}

	fromFen, _ := board.FromFen(fen)
	if fromFen.ToFen() != fen {
		t.Errorf("FromFen()\nActual = %s\nTarget = %s", fromFen, fen)
	}
}

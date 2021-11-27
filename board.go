package main

import (
	"fmt"
	"strings"
)

const (
	dimensions = 8
	firstRow   = 0
	firstCol   = 0
	lastRow    = dimensions - 1
	lastCol    = dimensions - 1
)

const (
	header       = "\n    ┌───┬───┬───┬───┬───┬───┬───┬───┐ \n"
	middle       = "\n    ├───┼───┼───┼───┼───┼───┼───┼───┤ \n"
	footer       = "\n    └───┴───┴───┴───┴───┴───┴───┴───┘ \n"
	footerLabels = "\n      A   B   C   D   E   F   G   H   \n"
	divider      = " │ "
)

type Board struct {
	Grid    [dimensions][dimensions]*Piece
	flipped bool
}

func (b *Board) Move(from, to string) {
	fromPos, toPos := ToPosition(from), ToPosition(to)
	p := b.Grid[fromPos[0]-1][fromPos[1]-1]
	if p == nil {
		return
	}
	cp := b.Grid[toPos[0]-1][toPos[1]-1]
	if cp != nil {
		// TODO: Capture
	}
	b.Grid[fromPos[0]-1][fromPos[1]-1] = nil
	b.Grid[toPos[0]-1][toPos[1]-1] = p
	p.Position = toPos
}

// Deserialize sets the board with pieces from a serialized format
// usually read from a file. Example: board.txt
func Deserialize(board string) *Board {
	var b Board
	lines := strings.Split(board, "\n")[:8]
	for row, line := range lines {
		for col, square := range line {
			if square == '.' {
				b.Grid[row][col] = nil
				continue
			}

			var piece = &Piece{}
			if square < 97 {
				piece.Color = White
			} else {
				piece.Color = Black
			}
			switch strings.ToUpper(string(square)) {
			case "B":
				piece.Type = Bishop
			case "K":
				piece.Type = King
			case "N":
				piece.Type = Knight
			case "P":
				piece.Type = Pawn
			case "Q":
				piece.Type = Queen
			case "R":
				piece.Type = Rook
			}
			piece.Position = position{row, col}
			b.Grid[row][col] = piece
		}
	}
	return &b
}

// Serialize returns a string of the current board
// usually read from a file. Example: boards/board.1
func Serialize(b *Board) string {
	var sb strings.Builder
	for _, rank := range b.Grid {
		for _, piece := range rank {
			if piece == nil {
				sb.WriteString(".")
				continue
			}
			var sp string
			switch piece.Type {
			case Bishop:
				sp = "B"
			case King:
				sp = "K"
			case Knight:
				sp = "N"
			case Pawn:
				sp = "P"
			case Queen:
				sp = "Q"
			case Rook:
				sp = "R"
			}
			if piece.Color == Black {
				sp = strings.ToLower(sp)
			}
			sb.WriteString(sp)
		}
		sb.WriteString("\n")
	}
	return strings.TrimSuffix(sb.String(), "\n")
}

func FileToColumn(file string) int {
	if len(file) < 1 {
		return 0
	}
	return int([]rune(strings.ToUpper(file))[0] - 64)
}

func (b *Board) String() string {
	var s = header

	var ranks []int

	if b.flipped {
		ranks = []int{0, 1, 2, 3, 4, 5, 6, 7}
	} else {
		ranks = []int{7, 6, 5, 4, 3, 2, 1, 0}
	}

	for row, rank := range ranks {
		for file := firstCol; file < dimensions; file++ {
			if file == firstCol {
				s += faintStyle.Render(fmt.Sprintf(" %d ", rank+1))
			}
			s += divider
			if b.Grid[rank][file] == nil {
				s += " "
			} else {
				s += b.Grid[rank][file].String()
			}
			if file == lastCol {
				s += divider
			}
		}
		if row != lastRow {
			s += middle
		}
	}
	return s + footer + faintStyle.Render(footerLabels)
}

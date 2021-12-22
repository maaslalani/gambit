package board

import (
	"errors"
	"strconv"
	"strings"

	"github.com/maaslalani/gambit/piece"
)

var ErrInvalidFEN = errors.New("Invalid FEN")

func (b Board) ToFen() string {
	var sb strings.Builder
	// Loop through the entire board and build the FEN string for each rank
	// The FEN string is built from the bottom up, so we need to reverse the grid
	for r := len(b.Grid) - 1; r >= 0; r-- {
		// Track the number of empty squares we have encountered so far before reaching a
		// non-empty square, dump this count as a replacement for pieces
		emptyCounter := 0
		// Loop through each column in the rank and convert it to its FEN equivalent
		for c := 0; c < len(b.Grid[r]); c++ {
			p := b.Grid[r][c]
			if p.Color == piece.NoColor {
				// Empty square
				emptyCounter += 1
			} else {
				// If we have encountered an empty square, dump the number of
				// empty squares we have encountered so far
				if emptyCounter > 0 {
					sb.WriteString(strconv.Itoa(emptyCounter))
				}
			}

			// Display the piece's Fen representation
			sb.WriteString(p.ToFen())
		}

		// If we have reached the end of the rank and we have encountered
		// empty squares dump the number of empty squares
		if emptyCounter > 0 {
			sb.WriteString(strconv.Itoa(emptyCounter))
		}

		if r > 0 {
			sb.WriteRune('/')
		}
	}

	sb.WriteString(" " + string(b.Turn) + " ")
	// TODO: Add castling
	sb.WriteString("KQkq")
	sb.WriteString(" - ")
	sb.WriteString("0 1")
	return sb.String()
}

func FromFen(fen string) (Board, error) {
	var b Board

	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return b, ErrInvalidFEN
	}

	b.Turn = piece.Color(parts[1])

	ranks := strings.Split(parts[0], "/")

	for r, rank := range ranks {
		for c, char := range rank {
			// Empty squares, append n empty pieces into the board
			if char >= '1' && char <= '8' {
				n := int(char - '0')
				for i := 0; i < n; i++ {
					b.Grid[7-r][c+i] = piece.Empty()
				}
				continue
			}

			p := piece.FromFen(string(char))
			b.Grid[7-r][c] = p
		}
	}

	return b, nil
}

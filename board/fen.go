package board

import (
	"strconv"
	"strings"

	"github.com/maaslalani/gambit/piece"
)

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
			sb.WriteString(p.Fen())
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

	sb.WriteString(" w ")
	// TODO: Add castling
	sb.WriteString("KQkq")
	sb.WriteString(" - ")
	sb.WriteString("0 1")
	return sb.String()
}

func FromFen(fen string) (Board, error) {
	return Board{}, nil
}

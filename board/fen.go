package board

import (
	"strconv"
	"strings"

	"github.com/maaslalani/gambit/piece"
)

// FromFen parses a FEN string and returns a Board with the corresponding
// attributes and pieces
func FromFen(fen string) (Board, error) {
	var b Board = New()

	// Split the FEN string into its component parts
	parts := strings.Split(fen, " ")

	// 1. Piece placement (from White's perspective). Each rank
	// is described, starting with rank 8 and ending with rank 1;
	// within each rank, the contents of each square are described
	// from file "a" through file "h". Following the Standard
	// Algebraic Notation (SAN), each piece is identified by a
	// single letter taken from the standard English names (pawn =
	// "P", knight = "N", bishop = "B", rook = "R", queen = "Q" and
	// king = "K"). White pieces are designated using upper-case
	// letters ("PNBRQK") while black pieces use lowercase
	// ("pnbrqk"). Empty squares are noted using digits 1 through 8
	// (the number of empty squares), and "/" separates ranks.
	ranks := strings.Split(parts[0], "/")
	for r, rank := range ranks {
		col := 0
		for _, char := range rank {
			if char >= '1' && char <= '8' {
				col += int(char - '0')
				continue
			}
			p := piece.FromFen(string(char))
			b.Grid[7-r][col] = p
			col += 1
		}
	}

	// 2. Active color.
	// "w" means White moves next, "b" means Black moves next.
	b.Turn = piece.Color(parts[1])

	// 3. Castling availability.
	// If neither side can castle, this is "-". Otherwise, this has one or more
	// letters: "K" (White can castle kingside), "Q" (White can castle queenside),
	// "k" (Black can castle kingside), and/or "q" (Black can castle queenside). A
	// move that temporarily prevents castling does not negate this notation.

	// 4. En passant target square in algebraic notation.
	// If there's no en passant target square, this is "-". If a pawn has just
	// made a two-square move, this is the position "behind" the pawn. This is
	// recorded regardless of whether there is a pawn in position to make an en
	// passant capture.

	// 5. Halfmove clock.
	// The number of halfmoves since the last capture or pawn advance, used for
	// the fifty-move rule.

	// 6. Fullmove number.
	// The number of the full move. It starts at 1, and is incremented after
	// Black's move.

	return b, nil
}

// ToFen converts a board into a FEN string
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
					emptyCounter = 0
				}
			}

			// Display the piece's Fen representation
			sb.WriteString(p.ToFen())
		}

		// If we have reached the end of the rank and we have encountered
		// empty squares dump the number of empty squares
		if emptyCounter > 0 {
			sb.WriteString(strconv.Itoa(emptyCounter))
			emptyCounter = 0
		}

		if r > 0 {
			sb.WriteRune('/')
		}
	}

	sb.WriteString(" " + string(b.Turn) + " ")

	// TODO: Add castling
	sb.WriteString("KQkq")

	// TODO: Add En passant target squares
	sb.WriteString(" - ")

	// TODO: Add halfmove + fullmove clock
	sb.WriteString("0 1")
	return sb.String()
}

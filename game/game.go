package game

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"

	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/border"
	"github.com/maaslalani/gambit/fen"
	"github.com/maaslalani/gambit/pieces"
	"github.com/maaslalani/gambit/position"
	. "github.com/maaslalani/gambit/style"
)

// model stores the state of the chess game. It tracks the board, legal moves,
// and the selected piece as well as the subset of legal moves for the selected
// piece.
type model struct {
	board      *dt.Board
	moves      []dt.Move
	pieceMoves []dt.Move
	selected   string
}

// InitialModel returns an initial model of the game board using the starting
// position of a normal chess game and generating the legal moves from the
// starting position.
func InitialModel() tea.Model {
	board := dt.ParseFen(dt.Startpos)
	return model{
		board: &board,
		moves: board.GenerateLegalMoves(),
	}
}

// Init Initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// View converts a FEN string into a chess board with all pieces and empty
// squares in a grid like pattern. We highlight the selected piece and the
// legal moves for that piece so the user know where they can move.
//
// For example, if the user selects the white pawn on E2 we indicate that they
// can move to E3 and E4 legally.
//
//    ┌───┬───┬───┬───┬───┬───┬───┬───┐
//  8 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  7 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  6 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  5 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  4 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  3 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  2 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  1 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │
//    └───┴───┴───┴───┴───┴───┴───┴───┘
//      A   B   C   D   E   F   G   H
//
func (m model) View() string {
	var s strings.Builder
	s.WriteString(border.Top())

	for r, row := range fen.Grid(m.board.ToFen()) {
		rr := board.LastRow - r
		s.WriteString(Faint(fmt.Sprintf(" %d ", rr+1)) + border.Vertical)

		for c, cell := range row {
			display := pieces.Display[cell]

			if m.selected == position.ToSquare(rr, c) {
				display = Cyan(display)
			}

			if isLegalMove(m.pieceMoves, position.ToSquare(rr, c)) {
				if cell == "" {
					display = "."
				}
				display = Magenta(display)
			}

			s.WriteString(fmt.Sprintf(" %s %s", display, border.Vertical))
		}
		s.WriteRune('\n')

		if r != board.LastRow {
			s.WriteString(border.Middle())
		}
	}

	s.WriteString(border.Bottom() + Faint(border.BottomLabels()))
	return s.String()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}

		// Find the square the user clicked on, this will either be our square
		// square for our piece or the destination square for a move if a piece is
		// already square and that destination square completes a legal move
		square := border.Cell(msg.X, msg.Y)

		// If the user has already selected a piece, check see if the square that
		// the user clicked on is a legal move for that piece. If so, make the move.
		if m.selected != "" {
			from := m.selected
			to := square

			for _, move := range m.pieceMoves {
				if move.String() == from+to {
					m.board.Apply(move)

					// We have applied a new move and the chess board is in a new state.
					// We must generate the new legal moves for the new state.
					m.moves = m.board.GenerateLegalMoves()

					// We have made a move, so we no longer have a selected piece or
					// legal moves for any selected pieces.
					m.selected = ""
					m.pieceMoves = []dt.Move{}

					return m, nil
				}
			}

			// The user clicked on a square that wasn't a legal move for the selected
			// piece, so we select the piece that was clicked on instead
			m.selected = to
		} else {
			m.selected = square
		}

		// After a mouse click, we must generate the legal moves for the selected
		// piece, if there is a newly selected piece
		m.pieceMoves = legalPieceMoves(m.moves, m.selected)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// isLegalMove determines whether it is legal to move the the destination
// square given a piece's legal moves
func isLegalMove(legalMoves []dt.Move, destination string) bool {
	for _, move := range legalMoves {
		if strings.HasSuffix(move.String(), destination) {
			return true
		}
	}
	return false
}

// legalPieceMoves returns the legal moves for a given piece this is usually
// for the selected piece so that we know to which we can move. If there is no
// selected piece we return an empty array of moves.
func legalPieceMoves(moves []dt.Move, selected string) []dt.Move {
	var legalMoves []dt.Move

	if selected == "" {
		return legalMoves
	}

	for _, move := range moves {
		if strings.HasPrefix(move.String(), selected) {
			legalMoves = append(legalMoves, move)
		}
	}

	return legalMoves
}

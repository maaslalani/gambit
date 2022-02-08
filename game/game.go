package game

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"

	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/border"
	"github.com/maaslalani/gambit/fen"
	"github.com/maaslalani/gambit/moves"
	"github.com/maaslalani/gambit/pieces"
	"github.com/maaslalani/gambit/position"
	. "github.com/maaslalani/gambit/style"
)

// model stores the state of the chess game.
//
// It tracks the board, legal moves, and the selected piece. It also keeps
// track of the subset of legal moves for the currently selected piece
type model struct {
	board      *dt.Board
	moves      []dt.Move
	pieceMoves []dt.Move
	selected   string
	buffer     string
	flipped    bool
}

// InitialModel returns an initial model of the game board.
func InitialModel(position string) tea.Model {
	if !fen.IsValid(position) {
		position = dt.Startpos
	}
	board := dt.ParseFen(position)

	return model{
		board: &board,
		moves: board.GenerateLegalMoves(),
	}
}

// Init Initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// View converts a FEN string into a human readable chess board. All pieces and
// empty squares are arranged in a grid-like pattern. The selected piece is
// highlighted and the legal moves for the selected piece are indicated by a
// dot (.) for empty squares. Pieces that may be captured by the selected piece
// are highlighted.
//
// For example, if the user selects the white pawn on E2 we indicate that they
// can move to E3 and E4 legally.
//
//    ┌───┬───┬───┬───┬───┬───┬───┬───┐
//  8 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  7 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  6 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  5 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  4 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  3 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  2 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  1 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │
//    └───┴───┴───┴───┴───┴───┴───┴───┘
//      A   B   C   D   E   F   G   H
//
func (m model) View() string {
	var s strings.Builder
	s.WriteString(border.Top())

	// Traverse through the rows and columns of the board and print out the
	// pieces and empty squares. Once a piece is selected, highlight the legal
	// moves and pieces that may be captured by the selected piece.
	var rows = fen.Grid(m.board.ToFen())

	for r := board.FirstRow; r < board.Rows; r++ {
		row := rows[r]
		rr := board.LastRow - r

		// reverse the row if the board is flipped
		if m.flipped {
			row = rows[board.LastRow-r]
			for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
				row[i], row[j] = row[j], row[i]
			}
			rr = r
		}

		s.WriteString(Faint(fmt.Sprintf(" %d ", rr+1)) + border.Vertical)

		for c, cell := range row {
			display := pieces.Display[cell]
			selected := position.ToSquare(r, c, m.flipped)

			// The user selected the current cell, highlight it so they know it is
			// selected.
			if m.selected == selected {
				display = Cyan(display)
			}

			// Show all the cells to which the piece may move. If it is an empty cell
			// we present a coloured dot, otherwise color the capturable piece.
			if moves.IsLegal(m.pieceMoves, selected) {
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

	s.WriteString(border.Bottom() + Faint(border.BottomLabels(m.flipped)))
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
		square := border.Cell(msg.X, msg.Y, m.flipped)
		return m.Select(square)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+f":
			m.flipped = !m.flipped
		case "a", "b", "c", "d", "e", "f", "g", "h":
			m.buffer = msg.String()
		case "1", "2", "3", "4", "5", "6", "7", "8":
			var move string
			if m.buffer != "" {
				move = m.buffer + msg.String()
				m.buffer = ""
			}
			return m.Select(move)
		case "esc":
			return m.Deselect()
		}
	}

	return m, nil
}

func (m model) Deselect() (tea.Model, tea.Cmd) {
	m.selected = ""
	m.pieceMoves = []dt.Move{}
	return m, nil
}

func (m model) Select(square string) (tea.Model, tea.Cmd) {
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
				return m.Deselect()
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
	m.pieceMoves = moves.LegalSelected(m.moves, m.selected)

	return m, nil
}

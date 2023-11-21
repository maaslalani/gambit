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

// MoveMsg is a message that controls the board from outside the model.
type MoveMsg struct {
	From string
	To   string
}

// NotifyMsg is a message that gets emitted when the user makes a move.
type NotifyMsg struct {
	From      string
	To        string
	Turn      bool
	Check     bool
	Checkmate bool
}

// Game stores the state of the chess game.
//
// It tracks the board, legal moves, and the selected piece. It also keeps
// track of the subset of legal moves for the currently selected piece
type Game struct {
	board      *dt.Board
	moves      []dt.Move
	pieceMoves []dt.Move
	selected   string
	buffer     string
	flipped    bool
	theme      Theme
}

// NewGame returns an initial model of the game board.
func NewGame() *Game {
	return NewGameWithPosition(dt.Startpos)
}

// NewGameWithPosition returns an initial model of the game board with the
// specified FEN position.
func NewGameWithPosition(position string) *Game {
	m := &Game{}

	if !fen.IsValid(position) {
		position = dt.Startpos
	}
	board := dt.ParseFen(position)
	m.board = &board
	m.moves = m.board.GenerateLegalMoves()

	return m
}

// NewGameWithPosition returns an initial model of the game board with the
// specified FEN position.
func NewGameWithPositionAndTheme(position string, theme Theme) *Game {
	m := NewGameWithPosition(position)
	m.theme = theme

	return m
}

// Init Initializes the model
func (m *Game) Init() tea.Cmd {
	return nil
}

// drawRankNumber returns the string with the rank number and the eventual
// border
func (m *Game) drawRankNumber(rank int) string {
	rankString := Faint(fmt.Sprintf(" %d ", rank+1))
	if !m.theme.IsValid() {
		rankString = rankString + border.Vertical
	}
	return rankString
}

// renderSelectedPiece returns the string of the piece with the selected
// color based on the theme
func (m *Game) renderSelectedPiece(display string, isLightSquare bool) string {
	if m.theme.IsValid() {
		return m.theme.SelectedSquare(display, isLightSquare)
	}
	return Cyan(display)
}

// renderPieceInCheck returns the string of the piece with the check
// color based on the theme
func (m *Game) renderPieceInCheck(display string, isLightSquare bool) string {
	if m.theme.IsValid() {
		return m.theme.CheckSquare(display, isLightSquare)
	}
	return Red(display)
}

// renderAvailableMove returns the string of the piece with the availableMove
// color based on the theme
func (m *Game) renderAvailableMove(display string, isLightSquare bool) string {
	if m.theme.IsValid() {
		return m.theme.AvailableMove(display, isLightSquare)
	}
	return Magenta(display)
}

// renderNormalPiece returns the string of the piece based on the theme
func (m *Game) renderNormalPiece(display string, isWhite bool, isLightSquare bool) string {
	if m.theme.IsValid() {
		display = m.theme.Piece(display, isWhite, isLightSquare)
	}
	return display
}

// drawSquare returns the string of the piece and square to show
func (m *Game) drawSquare(piece pieces.Piece, square string) string {
	var s strings.Builder
	whiteTurn := m.board.Wtomove
	display := piece.Display()
	check := m.board.OurKingInCheck()
	selected := square
	isLightSquare := position.IsLightSquare(square)

	if m.theme.IsValid() {
		// display first space
		s.WriteString(m.theme.Bg(" ", isLightSquare))
	}

	// The user selected the current cell, highlight it so they know it is
	// selected. If it is a check, highlight the king in red.
	if m.selected == selected {
		display = m.renderSelectedPiece(display, isLightSquare)
	} else if check && piece.IsKing() &&
		((whiteTurn && piece.IsWhite()) || (!whiteTurn && piece.IsBlack())) {
		display = m.renderPieceInCheck(display, isLightSquare)
	} else {
		display = m.renderNormalPiece(display, piece.IsWhite(), isLightSquare)
	}

	// Show all the cells to which the piece may move. If it is an empty cell
	// we present a coloured dot, otherwise color the capturable piece.
	if moves.IsLegal(m.pieceMoves, selected) && piece.IsEmpty() {
		display = m.renderAvailableMove(".", isLightSquare)
	}

	if m.theme.IsValid() {
		s.WriteString(display)
		s.WriteString(m.theme.Bg("  ", isLightSquare))
	} else {
		s.WriteString(fmt.Sprintf(" %s %s", display, border.Vertical))
	}

	return s.String()
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
//  8 │ ♖ │ ♘ │ ♗ │ ♕ │ ♔ │ ♗ │ ♘ │ ♖ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  7 │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │ ♙ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  6 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  5 │   │   │   │   │   │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  4 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  3 │   │   │   │   │ . │   │   │   │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  2 │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │ ♟ │
//    ├───┼───┼───┼───┼───┼───┼───┼───┤
//  1 │ ♜ │ ♞ │ ♝ │ ♛ │ ♚ │ ♝ │ ♞ │ ♜ │
//    └───┴───┴───┴───┴───┴───┴───┴───┘
//      A   B   C   D   E   F   G   H
//
func (m *Game) View() string {
	var s strings.Builder
	s.WriteString(border.Top(m.theme))

	// Traverse through the rows and columns of the board and print out the
	// pieces and empty squares. Once a piece is selected, highlight the legal
	// moves and pieces that may be captured by the selected piece.
	var rows = fen.Grid(m.board.ToFen())

	for r := board.FirstRow; r < board.Rows; r++ {
		row := pieces.ToPieces(rows[r])
		rr := board.LastRow - r

		// reverse the row if the board is flipped
		if m.flipped {
			row = pieces.ToPieces(rows[board.LastRow-r])
			for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
				row[i], row[j] = row[j], row[i]
			}
			rr = r
		}

		s.WriteString(m.drawRankNumber(rr))

		for c, piece := range row {
			square := position.ToSquare(r, c, m.flipped)
			s.WriteString(m.drawSquare(piece, square))
		}
		s.WriteRune('\n')

		if r != board.LastRow {
			firstSquareIsLight := position.IsLightSquare(position.ToSquare(r, 0, m.flipped))
			s.WriteString(border.Middle(m.theme, firstSquareIsLight))
		}
	}

	s.WriteString(border.Bottom(m.theme) + Faint(border.BottomLabels(m.flipped)))
	return s.String()
}

func (m *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case MoveMsg:
		m.selected = msg.From
		m.pieceMoves = moves.LegalSelected(m.moves, m.selected)
		return m.Select(msg.To)
	}

	return m, nil
}

func (m *Game) Notify(from, to string, turn, check, checkmate bool) tea.Cmd {
	return func() tea.Msg {
		return NotifyMsg{
			From: from, To: to, Turn: turn,
			Check: check, Checkmate: checkmate,
		}
	}
}

func (m *Game) Deselect() (tea.Model, tea.Cmd) {
	m.selected = ""
	m.pieceMoves = []dt.Move{}
	return m, nil
}

func (m *Game) Select(square string) (tea.Model, tea.Cmd) {
	// If the user has already selected a piece, check see if the square that
	// the user clicked on is a legal move for that piece. If so, make the move.
	if m.selected != "" {
		from := m.selected
		to := square

		for _, move := range m.pieceMoves {
			if move.String() == from+to || (move.Promote() > 1 && move.String() == from+to+"q") {
				var cmds []tea.Cmd
				m.board.Apply(move)

				// We have applied a new move and the chess board is in a new state.
				// We must generate the new legal moves for the new state.
				m.moves = m.board.GenerateLegalMoves()
				check := m.board.OurKingInCheck()
				checkmate := check && len(m.moves) == 0

				// We have made a move, so we no longer have a selected piece or
				// legal moves for any selected pieces.
				g, cmd := m.Deselect()
				cmds = append(cmds, cmd, m.Notify(from, to, m.board.Wtomove, check, checkmate))
				return g, tea.Batch(cmds...)
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

// SetFlipped sets the board to be flipped or not.
func (g *Game) SetFlipped(flip bool) {
	g.flipped = flip
}

// Position returns the current FEN position of the board.
func (g *Game) Position() string {
	return g.board.ToFen()
}

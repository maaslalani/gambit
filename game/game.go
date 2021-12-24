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
	"github.com/maaslalani/gambit/style"
)

type model struct {
	board      *dt.Board
	moves      []dt.Move
	pieceMoves []dt.Move
	selected   string
}

func Model() tea.Model {
	board := dt.ParseFen(dt.Startpos)
	moves := board.GenerateLegalMoves()
	return model{board: &board, moves: moves}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	var s strings.Builder

	grid := fen.Grid(m.board.ToFen())

	for r, row := range grid {
		if r == board.FirstRow {
			s.WriteString(border.Top())
		}

		for c, cell := range row {
			if c == board.FirstCol {
				label := fmt.Sprintf(" %d ", board.LastRow-r+1)
				s.WriteString(style.Faint.Render(label) + border.Vertical)
			}

			display := " " + pieces.Display[cell] + " "
			if m.selected == position.ToSquare(board.LastRow-r, c) {
				display = " " + style.Selected.Render(pieces.Display[cell]) + " "
			}

			// Loop through all piece legal moves and see if this square matches any
			for _, move := range m.pieceMoves {
				if strings.HasSuffix(move.String(), position.ToSquare(board.LastRow-r, c)) {
					if cell == "" {
						display = style.Cyan.Render(" . ")
					} else {
						display = style.Red.Render(display)
					}
					break
				}
			}
			s.WriteString(display + border.Vertical)
		}
		s.WriteString("\n")

		if r != board.LastRow {
			s.WriteString(border.Middle())
		} else {
			s.WriteString(border.Bottom() + style.Faint.Render(border.BottomLabels()))
		}
	}
	return s.String()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Type != tea.MouseLeft {
			return m, nil
		}

		col, row := border.Cell(msg.X, msg.Y)

		if m.selected != "" {
			from := m.selected
			to := fmt.Sprintf("%s", position.ToSquare(row, col))

			for _, move := range m.moves {
				if move.String() == from+to {
					// Perform move
					m.board.Apply(move)
					m.moves = m.board.GenerateLegalMoves()

					// Unselect piece
					m.selected = ""
					m.pieceMoves = []dt.Move{}

					return m, nil
				}
			}

			// We didn't encounter a valid move
			m.selected = to
			m.pieceMoves = []dt.Move{}
			for _, move := range m.moves {
				if strings.HasPrefix(move.String(), m.selected) {
					m.pieceMoves = append(m.pieceMoves, move)
				}
			}
		} else {
			m.selected = fmt.Sprintf("%s", position.ToSquare(row, col))
			m.pieceMoves = []dt.Move{}
			for _, move := range m.moves {
				if strings.HasPrefix(move.String(), m.selected) {
					m.pieceMoves = append(m.pieceMoves, move)
				}
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

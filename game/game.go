package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/piece"
	"github.com/maaslalani/gambit/position"
	. "github.com/maaslalani/gambit/squares"
)

type model struct {
	board board.Board
	moves []board.Move
}

func Model() tea.Model {
	b, _ := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return model{
		board: b,
	}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() string  { return m.board.String() }

const (
	cellHeight = 2
	cellWidth  = 4
	marginLeft = 4
	marginTop  = 2

	maxCol = 7
	maxRow = 7
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		col := (msg.X - marginLeft) / cellWidth
		row := (msg.Y - marginTop) / cellHeight

		if col < 0 || col > maxCol || row < 0 || row > maxRow {
			m.board.Selected = position.NoPosition
			return m, nil
		}

		if !m.board.Reversed {
			row = maxRow - row
		}

		if msg.Type != tea.MouseRelease {
			return m, nil
		}

		if m.board.Selected == position.NoPosition {
			pos := position.Position{Row: row, Col: col}
			if m.board.At(pos) == piece.EmptyPiece {
				return m, nil
			}
			m.board.Selected = pos
		} else {
			from := Square(m.board.Selected.String())
			toPos := position.Position{Row: row, Col: col}
			to := Square(toPos.String())

			// Don't allow moving to the same square
			if from == to {
				return m, nil
			}

			// Don't allow moving to a square with a piece of the same
			// color as the selected piece
			if m.board.At(m.board.Selected).Color == m.board.At(toPos).Color {
				m.board.Selected = toPos
				return m, nil
			}

			// Valid move
			move := board.Move{From: from, To: to}
			m.moves = append(m.moves, move)
			m.board.Move(move)
			m.board.Selected = position.NoPosition
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

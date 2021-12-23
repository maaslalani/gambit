package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/piece"
	"github.com/maaslalani/gambit/position"
	. "github.com/maaslalani/gambit/squares"
)

type model struct {
	board    board.Board
	selected position.Position
	moves    []board.Move
}

func Model() tea.Model {
	b, _ := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return model{
		board:    b,
		selected: position.NoPosition,
	}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() string  { return m.board.String() }

const (
	cellHeight = 2
	cellWidth  = 4
	marginLeft = 4
	marginTop  = 2
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		col := (msg.X - marginLeft) / cellWidth
		row := (msg.Y - marginTop) / cellHeight

		if col < 0 || col > 7 || row < 0 || row > 7 {
			m.selected = position.NoPosition
			return m, nil
		}

		if !m.board.Reversed {
			row = 7 - row
		}

		if msg.Type != tea.MouseRelease {
			return m, nil
		}

		if m.selected == position.NoPosition {
			pos := position.Position{Row: row, Col: col}
			if m.board.At(pos) == piece.EmptyPiece {
				return m, nil
			}
			m.selected = pos
		} else {
			from := Square(m.selected.String())
			toPos := position.Position{Row: row, Col: col}
			to := Square(toPos.String())

			// Don't allow moving to the same square
			if from == to {
				return m, nil
			}

			// Don't allow moving to a square with a piece of the same
			// color as the selected piece
			if m.board.At(m.selected).Color == m.board.At(toPos).Color {
				m.selected = toPos
				return m, nil
			}

			// Valid move
			move := board.Move{From: from, To: to}
			m.moves = append(m.moves, move)
			m.board.Move(move)
			m.selected = position.NoPosition
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

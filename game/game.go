package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/board"
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

		if !m.board.Reversed {
			row = 7 - row
		}

		if msg.Type == tea.MouseRelease {
			return m, nil
		}

		if m.selected == position.NoPosition {
			m.selected = position.Position{Row: row, Col: col}
		} else {
			from := Square(m.selected.String())
			to := Square(position.Position{Row: row, Col: col}.String())
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

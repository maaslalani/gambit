package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/board"
	"github.com/maaslalani/gambit/position"
	. "github.com/maaslalani/gambit/squares"
)

type model struct {
	board    board.Board
	selected *position.Position
	moves    []board.Move
}

func Model() tea.Model {
	b, _ := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return model{
		board: b,
	}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() string  { return m.board.String() + m.selected.String() }

const (
	marginLeft = 4
	marginTop  = 2
	cellWidth  = 4
	cellHeight = 2
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		col := (msg.X - marginLeft) / cellWidth
		row := (msg.Y - marginTop) / cellHeight

		if !m.board.Reversed {
			row = 7 - row
		}

		if m.selected != nil {
			m.selected = &position.Position{Row: row, Col: col}
		} else {
			from := m.selected
			to := &position.Position{Row: row, Col: col}
			m.moves = append(m.moves, board.Move{
				From: Square((*from).String()),
				To:   Square((*to).String()),
			})
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

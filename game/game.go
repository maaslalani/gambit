package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/board"
)

type model struct {
	board board.Board
}

func Model() tea.Model {
	b, _ := board.FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return model{
		board: b,
	}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() string  { return m.board.String() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"
)

type model struct {
	board dt.Board
}

func Model() tea.Model {
	board := dt.ParseFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	return model{
		board: board,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) View() string {
	var s strings.Builder
	ranks := strings.Split(strings.Split(m.board.ToFen(), " ")[0], "/")
	for _, r := range ranks {
		for _, c := range r {
			if c >= '1' && c <= '8' {
				for i := 0; i < int(c-'0'); i++ {
					s.WriteString(Display[""])
				}
			} else {
				s.WriteString(Display[string(c)])
			}
		}
		s.WriteRune('\n')
	}
	return s.String()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			move := m.board.GenerateLegalMoves()[0]
			m.board.Apply(move)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

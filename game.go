package main

import (
	"fmt"
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
	for r, rank := range ranks {
		if r == 0 {
			s.WriteString("   ┌───┬───┬───┬───┬───┬───┬───┬───┐\n")
		}
		for c, cell := range rank {
			if c == 0 {
				s.WriteString(fmt.Sprintf(" %d │", r))
			}
			if cell >= '1' && cell <= '8' {
				s.WriteString(strings.Repeat("   │", int(cell-'0')))
			} else {
				s.WriteString(" " + Display[string(cell)] + " │")
			}
		}
		s.WriteRune('\n')
		if r == 7 {
			s.WriteString("   └───┴───┴───┴───┴───┴───┴───┴───┘\n")
			s.WriteString("     A   B   C   D   E   F   G   H\n")
		} else {
			s.WriteString("   ├───┼───┼───┼───┼───┼───┼───┼───┤\n")
		}
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

package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model{}
	m.Board.Players = []Player{
		{Pieces: BlackPieces()},
		{Pieces: WhitePieces()},
	}
	p := tea.NewProgram(m, tea.WithMouseAllMotion())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	Board Board
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.Board.String()
}

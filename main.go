package main

import (
	_ "embed"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//go:embed board.txt
var initialBoard string

func main() {
	m := model{}
	m.Board = *Deserialize(string(initialBoard))
	p := tea.NewProgram(m, tea.WithMouseAllMotion())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	Board Board
	From  string
	To    string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			m.From, m.To = "", ""
		case "1", "2", "3", "4", "5", "6", "7", "8":
			rank := msg.String()
			// `From` and `To` moves will be complete, perform move
			if len(m.To) >= 1 {
				m.To += rank
				m.Board.Move(m.From, m.To)
				m.From, m.To = "", ""
			} else {
				m.From += rank
			}
			return m, nil
		case "a", "b", "c", "d", "e", "f", "g", "h",
			"A", "B", "C", "D", "E", "F", "G", "H":
			file := strings.ToUpper(msg.String())
			// If we already have the `From` rank set, set the `To` column
			if len(m.From) == 2 {
				m.To = file
			} else {
				m.From = file
			}

			return m, nil
		case "ctrl+f":
			m.Board.flipped = !m.Board.flipped
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.Board.String() + "\n"
}

package main

import (
	"log"
	"strconv"

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
	Move  struct {
		From position
		To   position
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Board.Draw()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			m.Move.From, m.Move.To = position{}, position{}
		case "1", "2", "3", "4", "5", "6", "7", "8":
			if m.Move.To[1] > 0 {
				i, _ := strconv.Atoi(msg.String())
				m.Move.To[0] = i
				m.Board.Move(m.Move.From, m.Move.To)
				m.Move.From, m.Move.To = position{}, position{}
				return m, nil
			}
			i, _ := strconv.Atoi(msg.String())
			m.Move.From[0] = i
		case "a", "b", "c", "d", "e", "f", "g", "h",
			"A", "B", "C", "D", "E", "F", "G", "H":
			if m.Move.From[0] > 0 {
				m.Move.To[1] = FileToColumn(msg.String())
				return m, nil
			}

			m.Move.From[1] = FileToColumn(msg.String())
			return m, nil
		case "ctrl+f":
			m.Board.flipped = !m.Board.flipped
			return m, nil
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.Board.String() + "\n\n" + m.Move.From.String() + " " + m.Move.To.String()
}

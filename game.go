package main

import (
	"fmt"
	"math/rand"
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

	fen := m.board.ToFen()
	board := strings.Split(fen, " ")[0]
	ranks := strings.Split(board, "/")

	for r, rank := range ranks {
		if r == firstRow {
			s.WriteString(topBorder())
		}

		for c, cell := range rank {
			if c == firstCol {
				label := fmt.Sprintf(" %d ", lastRow-r+1)
				s.WriteString(Faint.Render(label) + vertical)
			}

			if isNumeric(cell) {
				s.WriteString(strings.Repeat("   "+vertical, runeToInt(cell)))
			} else {
				s.WriteString(" " + Display[string(cell)] + " " + vertical)
			}
		}
		s.WriteRune('\n')

		if r != lastRow {
			s.WriteString(middleBorder())
		} else {
			s.WriteString(bottomBorder() + Faint.Render(bottomLabels()))
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
			moves := m.board.GenerateLegalMoves()
			move := moves[rand.Intn(len(moves))]
			m.board.Apply(move)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// isNumeric returns true if the current rune is a number
func isNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

// runeToInt converts a rune to an int
func runeToInt(r rune) int {
	return int(r - '0')
}

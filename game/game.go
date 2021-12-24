package game

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	dt "github.com/dylhunn/dragontoothmg"
	"github.com/maaslalani/gambit/piece"
)

type model struct {
	board dt.Board
}

func Model() tea.Model {
	board := dt.ParseFen("rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1")
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
					s.WriteString(piece.Display[""])
				}
			} else {
				s.WriteString(piece.Display[string(c)])
			}
		}
		s.WriteRune('\n')
	}
	return s.String()
}

const (
	cellHeight = 2
	cellWidth  = 4
	marginLeft = 4
	marginTop  = 2

	maxCol = 7
	maxRow = 7
)

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

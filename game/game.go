package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func Model() tea.Model {
	return model{}
}

func (m model) Init() tea.Cmd { return nil }
func (m model) View() string  { return "game" }

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

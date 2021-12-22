package style

import "github.com/charmbracelet/lipgloss"

var (
	Faint = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	White = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	Black = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
)

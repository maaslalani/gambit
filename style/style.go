package style

import (
	. "github.com/charmbracelet/lipgloss"
)

var (
	Faint    = NewStyle().Foreground(Color("8"))
	White    = NewStyle().Foreground(Color("7"))
	Black    = NewStyle().Foreground(Color("4"))
	Selected = NewStyle().Foreground(Color("2"))
)

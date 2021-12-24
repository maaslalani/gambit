package style

import . "github.com/charmbracelet/lipgloss"

func foreground(color string) Style {
	return NewStyle().Foreground(Color(color))
}

var Cyan = foreground("5")
var Faint = foreground("8")
var Red = foreground("1")
var Selected = foreground("6")

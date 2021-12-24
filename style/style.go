package style

import . "github.com/charmbracelet/lipgloss"

func fg(color string) Style {
	return NewStyle().Foreground(Color(color))
}

var Cyan = fg("5")
var Faint = fg("8")
var Red = fg("1")
var Selected = fg("6")

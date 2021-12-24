package style

import . "github.com/charmbracelet/lipgloss"

type colorFunc func(s string) string

func fg(color string) colorFunc {
	return NewStyle().Foreground(Color(color)).Render
}

var Magenta = fg("5")
var Faint = fg("8")
var Red = fg("1")
var Cyan = fg("6")

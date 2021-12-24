package main

import (
	lg "github.com/charmbracelet/lipgloss"
)

var (
	Faint    = lg.NewStyle().Foreground(lg.Color("8"))
	White    = lg.NewStyle().Foreground(lg.Color("7"))
	Black    = lg.NewStyle().Foreground(lg.Color("4"))
	Selected = lg.NewStyle().Foreground(lg.Color("2"))
)

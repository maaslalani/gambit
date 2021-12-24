package main

import (
	lg "github.com/charmbracelet/lipgloss"
)

var (
	Black    = lg.NewStyle().Foreground(lg.Color("4"))
	Faint    = lg.NewStyle().Foreground(lg.Color("8"))
	Cyan     = lg.NewStyle().Foreground(lg.Color("5"))
	Red      = lg.NewStyle().Foreground(lg.Color("1"))
	Selected = lg.NewStyle().Foreground(lg.Color("6"))
	White    = lg.NewStyle().Foreground(lg.Color("7"))
)

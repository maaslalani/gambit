package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		Model(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	err := p.Start()
	if err != nil {
		log.Fatal(err)
	}
}

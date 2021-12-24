package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/game"
)

func main() {
	p := tea.NewProgram(
		game.InitialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	err := p.Start()
	if err != nil {
		log.Fatal(err)
	}
}

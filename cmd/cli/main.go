package main

import (
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/eduardpeters/go-game/internal/ui"
)

func main() {
	m := ui.NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

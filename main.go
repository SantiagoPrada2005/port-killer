package main

import (
	"fmt"
	"os"

	"port-killer/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Initialize the TUI program
	p := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/SnehGhetiya/kea/internal"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	agentModel := internal.New()

	p := tea.NewProgram(agentModel)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

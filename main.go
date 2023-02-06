package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Main function that runs the program
func main() {
	// Uses the Kanban list function
	models = []tea.Model{New(), NewForm(todo)}
	m := models[model]
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

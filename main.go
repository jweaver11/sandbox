package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Main function that runs the program
func main() {
	// Sets p as a new tea program using out startUpModel
	p := tea.NewProgram(startUpModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running the program: %v", err)
		os.Exit(1)
	}

}

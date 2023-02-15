package main

import (
	"fmt"
	"os"

	//"sandbox/styling"		// How to import packages

	tea "github.com/charmbracelet/bubbletea"
)

// Main function that runs the program
func main() {
	// Checks to see if the server is up first
	// For now it just uses the charm server
	//
	// CHECK SERVER BEFORE BULDING
	//if err := tea.NewProgram(CheckServerModel{}).Start(); err != nil {
	//fmt.Printf("There was an error fool: %v\n", err)
	//os.Exit(1)
	//} else {
	// Runs the program if there is not an error with the server

	// Sets p as a new tea program using out startUpModel
	p := tea.NewProgram(CreateProjectViewModel(), tea.WithAltScreen()) //tea.WithAltScreen starts in fullscreen mode
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running the program: %v", err)
		os.Exit(1)
	}
	//

}

package main

import (
	"fmt"
	"os"

	// How to import packages
	//"sandbox/helpers"
	//"sandbox/styling"

	tea "github.com/charmbracelet/bubbletea"
)

// Main function that runs the program
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

var ActiveModelInt int

func run() (err error) {
	// Checks to see if the server is up first
	// For now it just uses the charm server
	//
	// CHECK SERVER BEFORE BULDING
	//if err := tea.NewProgram(CheckServerModel{}).Start(); err != nil {
	//fmt.Printf("There was an error connecting to the server: %v\n", err)
	//os.Exit(1)
	//} else {

	// Runs the program if there is not an error with the server
	// Sets p as a new tea program using out startUpModel

	ActiveModelInt = 0

	p := tea.NewProgram(CreateDescriptionModel(), tea.WithAltScreen()) //tea.WithAltScreen starts in fullscreen mode
	//n := tea.NewProgram(CreateNewModel(), tea.WithAltScreen())

	if ActiveModelInt == 0 {
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running the program: %v", err)
			os.Exit(1)
			return err
		}
	} //else {
	//if _, err := n.Run(); err != nil {
	//fmt.Printf("Error running the program: %v", err)
	//os.Exit(1)
	//return err
	//}
	//}

	return err
}

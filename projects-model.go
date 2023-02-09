// OPEN SANDBOX PROGRAM WITH THE GOALS OF ADDING ITEMS TO A LIST, AND BEING ABLE TO SELECT THEM TO PRESENT A DIFFERENT
// MODEL DESCRIBING THE ITEM

// map example --> moons := make(map[string]string)
// moons["Jupiter"] = "Europa"
// HELP
// https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Sets the tasks as a struct
type TasksModel struct {
	items        []string
	descriptions []string
	cursor       int
}

// Sets the items and descriptions of our start up model
func startUpModel() TasksModel {
	return TasksModel{
		items:        []string{"Pirates of the Cryptobbean", "Haramgay"},
		descriptions: []string{"Dank Pirates", "Gay Harambe NFT's"},
	}
}

/************* MAIN FUNCTIONS ********************/
// Initializes the model at start of program.
// Returns a command if there is one
func (t TasksModel) Init() tea.Cmd {
	return nil
}

// Runs whenever there is an update or event
func (t TasksModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Sets the msg to types
	switch msg := msg.(type) {

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("\n\n\n\nbye bye bozo")
			return t, tea.Quit

		// Moves the cursor up
		case "up":
			if t.cursor > 0 {
				t.cursor--
			}

		// Moves the cursor up
		case "down":
			if t.cursor < len(t.items) {
				t.cursor++
			}

			// Toggle selected view to return a new model of item cursor is hovering
		case " ", "enter":
			// Placeholder
			return t, nil

		}

	}

	return t, nil
}

// Renders the view so the user can see the updated model
func (t TasksModel) View() string {
	// Sets s as a string to return out entire model
	// This Sets the header before s returns the model
	s := "What project would you like to know more about?\n\n"

	// Iterate over the individual projects in items
	for i, item := range t.items {

		// Is the cursor pointing at this choice
		cursor := " " // No cursor

		if t.cursor == i {
			cursor = ">" //Sets cursor as >
		}

		//
		s += fmt.Sprintf("%s  %s %s\n", cursor, item, t.descriptions)
	}

	// Sets a footer at end of s
	s += "\n\n\nPress q to quit"
	return s
}

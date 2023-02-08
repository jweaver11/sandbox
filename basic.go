// OPEN SANDBOX PROGRAM WITH THE GOALS OF ADDING ITEMS TO A LIST, AND BEING ABLE TO SELECT THEM TO PRESENT A DIFFERENT
// MODEL DESCRIBING THE ITEM
package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Sets an array of tea models
var models []tea.Model

// Sets the tasks as a struct
type Task struct {
	title       string
	description string
}

// Creates a new model
func NewModel() *Model {
	return &Model{}
}

/************* MAIN FUNCTIONS ********************/
// Initializes the model at start of program.
// Returns a command if there is one
func (m Model) Init() tea.Cmd {
	return nil
}

// Runs whenever there is an update or event
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

// Renders the view so the user can see the updated model
func (m model) View() string {
	return nil
}

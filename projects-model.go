// OPEN SANDBOX PROGRAM WITH THE GOALS OF ADDING ITEMS TO A LIST, AND BEING ABLE TO SELECT THEM TO PRESENT A DIFFERENT
// MODEL DESCRIBING THE ITEM
// https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics

// TASKS:
// Setup correct help options at bottom
// Begin work on more detailed description models
// Better formatting and design of main project view model

package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Sets the Projects as a struct
// This is our main model for the projects page
type ProjectViewModel struct {
	items, descriptions []string        // Each project with a short description
	cursor              int             // Used to track the cursor's location
	keys                keyMap          // Sets a keymap needed to use the help view
	help                help.Model      // Sets help as a help.Model so we can add it automatically to the bottom of our model
	inputStyle          lipgloss.Style  // Styling
	paginator           paginator.Model // Adds page scrolling to bottom of page
}

// This function is run in main to start a new program
// It sets our previously defined model with values
func createProjectViewModel() ProjectViewModel {
	// Sets items and descriptions new so we can change them easier here, and return them later
	var items, descriptions []string

	items = []string{"Pirates of the Cryptobbean", "Haramgay", "Another Dank Project here", "Midget Wrestling"}
	descriptions = []string{"Dank Pirates", "Gay Harambe NFT's", "Description of Dank Project", "Is Badass"}

	// Initializes the page scrolling for our list of items
	p := paginator.New()    // Sets p as a new paginator we can return later
	p.Type = paginator.Dots // Renders dots for our itmes
	p.PerPage = 1           // Items per page
	p.ActiveDot = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•") // Selected page formatting
	p.InactiveDot = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•") // Non-selected pages formatting
	p.SetTotalPages(len(items)) // Total number of pages

	return ProjectViewModel{
		items:        items,
		descriptions: descriptions,
		keys:         keys,
		help:         help.New(),
		inputStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#FF75B7")),
		paginator:    p,
	}
}

/************* MAIN FUNCTIONS TO SETUP PROJECT VIEW MODEL ********************/
// Initializes the model at start of program.
// Returns a command if there is one
func (p ProjectViewModel) Init() tea.Cmd {
	return nil
}

// Runs whenever there is an update or event
func (p ProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Sets the msg to types
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		p.help.Width = msg.Width

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("\n\n\n\nbye bye bozo")
			return p, tea.Quit

		// Moves the cursor up
		case "up":
			if p.cursor > 0 {
				p.cursor--
			}

		// Moves the cursor up
		case "down":
			if p.cursor < len(p.items) {
				p.cursor++
			}

		// Toggle selected view to return a new model of item cursor is hovering
		case " ", "enter":
			// Placeholder
			return p, nil

		// Toggles the help view between mini and full view
		case "h":
			p.help.ShowAll = !p.help.ShowAll

		}
		//var cmd tea.Cmd
		p.paginator, cmd = p.paginator.Update(msg)
	}

	// Returns our updated model with no command
	return p, cmd
}

// Renders the view so the user can see the updated model
func (p ProjectViewModel) View() string {
	// Sets s as a string to return out entire model
	// This Sets the header before s returns the model
	s := "What project would you like to know more about?\n\n"

	// Iterate over the individual projects in items
	for i, item := range p.items {

		// Is the cursor pointing at this choice
		cursor := " " // No cursor

		if p.cursor == i {
			cursor = ">" //Sets cursor as >
		}

		// Sets the individual descriptions as one variable to be returned
		description := p.descriptions[i]

		// Returns the model as a string, starting with the cursor, the item, then description
		s += fmt.Sprintf("%s  %s    [%s]\n", cursor, item, description)

	}

	// Sets a variable fullHelpView as a string to return our pages menu help view,
	// Which is managed automatically by the help package
	fullHelpView := ("   " + p.paginator.View() + "\n" + p.help.View(p.keys))
	height := 8 - strings.Count("0", "\n") - strings.Count(fullHelpView, "\n")
	s += "\n" + strings.Repeat("\n", height) + fullHelpView

	return s
}

package main

import (
	"sandbox/styling"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type DescriptionModel struct {
	descriptions []string
	width        int
	height       int
}

func CreateDescriptionModel() DescriptionModel {
	var descriptions []string

	descriptions = []string{
		"Dank Pirates",
		"Gay Harambe NFT's",
		"Description of Dank Project",
		"Is Badass",
	}

	return DescriptionModel{
		descriptions: descriptions,
	}
}

// ********************** BUBBLE TEA BUILT IN FUNCTIONS ***********************
// Initializes the model at start of program.
// Returns a command if there is one
func (d DescriptionModel) Init() tea.Cmd {
	return nil
}

// Runs whenever there is an update or event
func (d DescriptionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Sets the msg to types
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		// Sets the height and width of terminal so the border shows correctly
		d.width = msg.Width - 6
		d.height = msg.Height - 6

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return d, tea.Quit

		case "p":
			return CreateProjectViewModel(), nil
		}
	}
	// Returns our updated model with no command
	return d, cmd
}

// Renders the view so the user can see the updated model
func (d DescriptionModel) View() string {
	// Sets s as a string builder to return out entire model
	// Will return as a string later
	var s strings.Builder

	// A final string that is used to format all the styles
	// And can add one background/border too in the end
	var finalStr string

	// Renders the header
	finalStr += styling.HeaderStyle.Render("This is a different Models Header")

	finalStr += "\n\n"

	for _, item := range d.descriptions {
		finalStr += styling.ItemStyle.Render(item) + "\n"
	}

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(d.width).Height(d.height).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

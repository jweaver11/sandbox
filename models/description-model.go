package main

import (
	"sandbox/styling"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DescriptionModel struct {
	project     string
	description string
	width       int
	height      int
	selected    int
}

func CreateDescriptionModel(projectName string, cursor int) DescriptionModel {
	var description string

	project := projectName

	switch cursor {
	case 0:
		description = "Project 1 Description"
	case 1:
		description = "Project 2 Description"
	case 2:
		description = "Project 3 Description"
	case 3:
		description = "Project 4 Description"
	case 4:
		description = "Project 5 Description"
	case 5:
		description = "Project 6 Description"
	case 6:
		description = "Project 7 Description"

	default:
		description = "No Project selected"
	}

	return DescriptionModel{
		project:     project,
		description: description,
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
		case "ctrl+c", "q":
			return d, tea.Quit

		case "esc", " ":
			return CreateProjectViewModel(), nil

		default:
			return d, cmd
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

	finalStr += "\n"

	// Renders the header
	finalStr += styling.HeaderStyle.UnsetForeground().Render("Projects")
	finalStr += styling.HeaderStyle.Foreground(lipgloss.Color("#7D56F4")).Render("Descriptions")

	finalStr += "\n\n"

	finalStr += styling.ItemStyle.Render(d.description) + "\n\n"

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(d.width).Height(d.height).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

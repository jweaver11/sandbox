package main

import (
	"sandbox/styling"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2  // Universal padding
	maxWidth = 80 // Universal width
)

type DescriptionModel struct {
	project     string
	description string
	width       int
	height      int
	selected    int
	progressBar progress.Model // Progress bar
}

// Time variable that returns a message every tick
// We set every tick to one second
type tickMsg time.Time

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
	case 7:
		description = "Project 8 description"

	default:
		description = "No Project selected"
	}

	return DescriptionModel{
		project:     project,
		description: description,
		progressBar: progress.New(progress.WithDefaultGradient()),
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// ********************** BUBBLE TEA BUILT IN FUNCTIONS ***********************
// Initializes the model at start of program.
// Returns a command if there is one
func (d DescriptionModel) Init() tea.Cmd {
	return tickCmd()
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

		// Sets the progress bar width
		// If its bigger than the window, then it sets it to the size of the window
		d.progressBar.Width = msg.Width - padding*2 - 4
		if d.progressBar.Width > maxWidth {
			d.progressBar.Width = maxWidth
		}

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return d, tea.Quit

		case "esc", " ":
			return CreateProjectViewModel(), nil

		case "up", "right":
			barUp := d.progressBar.IncrPercent(0.2)
			return d, tea.Batch(tickCmd(), barUp)

		case "down", "left":
			barDown := d.progressBar.DecrPercent(0.2)
			return d, tea.Batch(tickCmd(), barDown)

		default:
			return d, cmd
		}

		// Returns every tick
	case tickMsg:

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := d.progressBar.Update(msg)
		d.progressBar = progressModel.(progress.Model)
		return d, cmd

	default:
		return d, nil
	}

	return d, nil
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

	finalStr += styling.ItemStyle.Render(d.description) + "\n\n\n"

	finalStr += d.progressBar.View()

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(d.width).Height(d.height).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

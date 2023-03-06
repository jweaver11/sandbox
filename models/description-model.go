package main

import (
	"fmt"
	"sandbox/styling"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding   = 2   // Universal padding
	maxWidth  = 133 // Width of terminal at full screen -- just on my laptop??
	maxHeight = 33  // Height of terminal at full screen
)

type DescriptionModel struct {
	project     string         // Projects name
	description string         // Projects description
	width       int            // Width of model
	height      int            // Height of Model
	progressBar progress.Model // Progress bar
}

// Creates our defined model with actual values and then returns its
func CreateDescriptionModel(projectName string, cursor int) DescriptionModel {

	// Sets our project name and description that are passed when model is built
	project := projectName
	description := projectName + " Description biaaaatch"

	tickCmd()

	// Returns our model
	return DescriptionModel{
		project:     project,
		description: description,
		progressBar: progress.New(progress.WithDefaultGradient()),
	}
}

// ********************** BUBBLE TEA BUILT IN FUNCTIONS ***********************
// Initializes the model at start of program.
func (d DescriptionModel) Init() tea.Cmd {
	return tickCmd() // Returns our tick command used for the progress bar
}

// Time variable that returns a message every tick
type tickMsg time.Time

// tickCmd returns a tea command every second
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Runs whenever there is an update or event
func (d DescriptionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Sets the cmd for easy return later
	var cmd tea.Cmd

	// Sets switch cases for the msg, which is the key press
	switch msg := msg.(type) {

	// Doesnt run on first start up since its switched to this model
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		d.width = msg.Width
		d.height = msg.Height

		// Sets the progress bar width
		// If its bigger than the window, then it sets it to the size of the window
		d.progressBar.Width = d.width - padding

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return d, tea.Quit

		// When esc or space pressed, return ProjectViewModel
		case "esc", " ":
			return CreateProjectViewModel(), cmd

		// When up or right arrow pressed, move progress bar up
		case "up", "right":
			barUp := d.progressBar.IncrPercent(0.2)
			return d, tea.Batch(tickCmd(), barUp)

		// When down or left arrow presseed, move progress bar down
		case "down", "left":
			barDown := d.progressBar.DecrPercent(0.2)
			return d, tea.Batch(tickCmd(), barDown)

		case "p":
			fmt.Println(d.width)

		default:
			return d, cmd
		}

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
	// Sets s as a string builder needed for paginator
	var s strings.Builder

	// A final string that is used to pass styles onto it
	var finalStr string

	// Renders the header
	finalStr += styling.HeaderStyle.UnsetForeground().Render("Projects")
	finalStr += styling.HeaderStyle.Foreground(lipgloss.Color("#7D56F4")).Render("Descriptions") + "\n\n"

	// Adds the project name
	finalStr += styling.ItemStyle.Foreground(lipgloss.Color("12")).Render(d.project)

	// Adds the project description
	finalStr += styling.FullDescStyle.Render(d.description) + "\n\n\n"

	// Adds the progress bar
	finalStr += d.progressBar.View()

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(d.width).Height(d.height).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

package main

import (
	"sandbox/styling"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type DescriptionModel struct {
	project     string         // Projects name
	description string         // Projects description
	progressBar progress.Model // Progress bar
	termWidth   int            // Sets terminal width
	termHeight  int            // Sets terminal height
}

// Creates our defined model with actual values and then returns its
func CreateDescriptionModel(projectName string) DescriptionModel {

	// Sets our project name and description that are passed when model is built
	project := projectName
	description := projectName + " Description biaaaatch"

	PB := progress.New(progress.WithDefaultGradient()) // scaleRamp = false

	// Gets Height and Width of Terminal Size
	TW, TH, _ := term.GetSize(0)

	// Sets progress bar width small enough that it wont take up two lines
	PB.Width = TW - 20

	// Returns our model
	return DescriptionModel{
		project:     project,
		description: description,
		progressBar: progress.New(progress.WithDefaultGradient()),
		termWidth:   TW,
		termHeight:  TH,
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

		// Gets terminal size
		d.termWidth, d.termHeight, _ = term.GetSize(0)

		// Sets progress bar width to account for padding
		d.progressBar.Width = d.termWidth - styling.ProgressBarStyle.GetPaddingLeft() - 6

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return d, tea.Quit

		// When esc or space pressed, return ProjectViewModel
		case "esc", " ":
			cmd = tea.Batch(tea.ClearScreen, tea.EnterAltScreen)
			return CreateProjectViewModel(), cmd

		// When up or right arrow pressed, move progress bar up
		case "up", "right":
			barUp := d.progressBar.IncrPercent(0.2)
			return d, tea.Batch(tickCmd(), barUp)

		// When down or left arrow presseed, move progress bar down
		case "down", "left":
			barDown := d.progressBar.DecrPercent(0.2)
			return d, tea.Batch(tickCmd(), barDown)

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
	finalStr += "\n" + styling.HeaderStyle.UnsetForeground().Render("Projects")
	finalStr += styling.HeaderStyle.Foreground(lipgloss.Color("#7D56F4")).Render("Descriptions") + "\n\n"

	// Adds the project name
	finalStr += styling.ItemStyle.Foreground(lipgloss.Color("12")).Render(d.project)

	// Adds the project description
	finalStr += styling.FullDescStyle.Render(d.description)

	// Sets the height == to max height - each new line in the finalStr
	height := d.termHeight - strings.Count(finalStr, "\n") // Counts the number of lines the string takes up
	height -= strings.Count(d.progressBar.View(), "\n") + 2
	height -= strings.Count("0", "\n") // Counts all remaining lines left before bottom of terminal

	// If progress bar width bigger than terminal width set them equal

	finalStr += strings.Repeat("\n", height) + styling.ProgressBarStyle.Render(d.progressBar.View())

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(d.termWidth).Height(d.termHeight).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

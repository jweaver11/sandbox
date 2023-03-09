// OPEN SANDBOX PROGRAM WITH THE GOALS OF ADDING ITEMS TO A LIST, AND BEING ABLE TO SELECT THEM TO PRESENT A DIFFERENT

// TASKS:
// Model breaks if terminal gets to 1 item per page or full help view called below six
// Add copy to clipboard command for future github open source stuff
// Try other new stuff
// Spinner freezes after model switch

package main

import (
	"fmt"
	"strings"

	"sandbox/styling"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Our project page model
type ProjectViewModel struct {
	items, descriptions []string        // Each project item name and a short description
	cursor              int             // Used to track the cursor's location
	keys                keyMap          // Sets a keymap needed to use the help view
	help                help.Model      // Sets help as a help.Model so we can add it automatically to the bottom of our model
	paginator           paginator.Model // Adds page scrolling to bottom of page
	spinner             spinner.Model   // Adds the spinner to be used as a cursor
	err                 error           // Error that can be returned
	termWidth           int             // Sets terminal width
	termHeight          int             // Sets terminal height
}

// Creates our defined model with actual values and then returns it
func CreateProjectViewModel() ProjectViewModel {
	// Sets items and descriptions new so we can change them easier here, and return them later
	var items, descriptions []string

	// Temporary just numbers a bunch of project titles and descriptions
	for i := 1; i < 36; i++ {
		text := fmt.Sprintf("Project: %d", i)
		desc := fmt.Sprintf("Short Description: %d", i)
		items = append(items, text)
		descriptions = append(descriptions, desc)
	}
	TW, TH, _ := term.GetSize(0)

	// SPINNER -- Sets up our spinner
	s := spinner.New()             // Sets s as a new spinner
	s.Spinner = spinner.Line       // Sets the dots style spinner
	s.Style = styling.SpinnerStyle // Uses the spinner styling

	// PAGINATOR -- Initializes the page scrolling for our list of items
	p := paginator.New()                                 // Sets p as a new paginator
	p.Type = paginator.Dots                              // Using dots as the pages
	p.PerPage = 8                                        // Number of items per page
	p.ActiveDot = styling.ActiveDotStyle.Render("•")     // Selected page styling
	p.InactiveDot = styling.InactiveDotStyle.Render("•") // Non-selected pages formatting
	p.SetTotalPages(len(items))                          // Sets the total number of pages
	p.KeyMap.NextPage.Unbind()                           // Unbinds the Next page command so we can customize it ourselves
	p.KeyMap.PrevPage.Unbind()                           // Unbinds the Prev page command so we can customize it

	// Returns our model
	return ProjectViewModel{
		items:        items,
		descriptions: descriptions,
		keys:         keys,
		help:         help.New(),
		paginator:    p,
		spinner:      s,
		termWidth:    TW,
		termHeight:   TH,
	}
}

// ********************** BUBBLE TEA BUILT IN FUNCTIONS ***********************
// Initializes the model at start of program and returns a command if there is one
func (p ProjectViewModel) Init() tea.Cmd {
	// return p.spinner.Tick // Starts the spinner when program starts
	return nil
}

// Runs whenever there is an update or event
func (p ProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Sets the cmd for easy return later
	var cmd tea.Cmd

	// Sets switch cases for the msg, which is the key press
	switch msg := msg.(type) {

	// Runs whenever the window is resized or first loaded
	case tea.WindowSizeMsg:

		// Sets the help model and main model width for sizing later
		p.help.Width = msg.Width - styling.HelpBarStyle.GetPaddingLeft()

		// Sets terminal width and height
		p.termWidth, p.termHeight, _ = term.GetSize(0)

		// Set Per page accordance to window size

		// Sets the minimum height so the model will resize to fit smaller terminal instead of breaking
		//........... Items per page + new lines per item....... paginator lines.....................Help view lines.........Additional lines added
		minHeight := p.paginator.PerPage*3 + strings.Count(p.paginator.View(), "\n") + strings.Count(p.help.View(p.keys), "\n") + 4
		minWidth := 10

		// Check if the window size is less than the minimum width and height
		if p.termWidth < minWidth || p.termHeight < minHeight {

			// If there is more than one item per page, subtract one item per page
			if p.paginator.PerPage > 1 {
				p.paginator.PerPage -= 1
			}
			if p.paginator.PerPage*p.paginator.TotalPages < len(p.items) {
				p.paginator.TotalPages += 1
			}
		}

	// Evertime the spinner ticks, so every second
	case spinner.TickMsg:
		s, cmd := p.spinner.Update(msg) // Update the spinner
		p.spinner = s
		return p, cmd

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {

		// Three keys that quit the program
		case "ctrl+c", "q", "esc":
			return p, tea.Quit

		// When up is pressed, move the cursor up
		case "up":

			// If on first page and cursor on first item, don't go up
			if p.paginator.Page == 0 && p.cursor == 0 {

				// Otherwise move cursor up
			} else {
				p.cursor-- // Move cursor up
			}

			// Moving cursor up when on first item of page that is not the first, change page to previous one
			if p.cursor == -1 {
				p.paginator.PrevPage()             // Go to previous page
				p.cursor = p.paginator.PerPage - 1 // Go to last item on previous page
			}

		// When down pressed, move cursor down
		case "down":

			// Won't move cursor down if on last item of last page
			if p.paginator.OnLastPage() == true { // Check if on last page

				// Move cursor correctly if all items fit evenly into num of pages
				if len(p.items)%p.paginator.PerPage == 0 {

					if p.cursor == p.paginator.PerPage-1 { // Check if on last item
						// Do Nothing
					} else {
						p.cursor++
					}

					// Move cursor correctly if items dont fit evenly in num of pages
				} else if p.cursor < (len(p.items)%p.paginator.PerPage - 1) { // Check if on last item
					p.cursor++
				}

			} else { // If not on last page move the cursor down
				p.cursor++
			}

			// Move cursor to start of next page if on last item
			if p.cursor == p.paginator.PerPage { // Move cursor to next page if on last item of page
				p.paginator.NextPage()
				p.cursor = 0
			}

		// When left arrow pressed, move to previous page
		case "left":
			p.paginator.PrevPage()

		// When right arrow pressed, move to next page
		case "right":
			p.paginator.NextPage()

			// If cursor below last item on last page, put it on last item
			if p.paginator.OnLastPage() && p.cursor >= len(p.items)%p.paginator.PerPage {
				if len(p.items)%p.paginator.PerPage == 0 {
					// Do Nothing
				} else {
					p.cursor = len(p.items)%p.paginator.PerPage - 1
				}
			}

		// When ? pressed, channge between short help view and full help view
		case "?":
			p.help.ShowAll = !p.help.ShowAll

		// If space pressed, switches to description model of selected project
		case " ":
			if p.cursor > len(p.items)%p.paginator.PerPage && p.paginator.OnLastPage() { // Prevent error if no project selected
				return p, nil
			} else {
				cmd = tea.Batch(tea.ClearScreen, tea.EnterAltScreen)
				return CreateDescriptionModel(p.items[p.cursor+(p.paginator.Page*p.paginator.PerPage)]), cmd
			}

		}
	}

	// Updates spinner and paginator
	p.spinner, cmd = p.spinner.Update(msg)
	p.paginator, cmd = p.paginator.Update(msg)

	// Returns our updated model with any commands
	return p, cmd
}

// Renders the view so the user can see the updated model
func (p ProjectViewModel) View() string {
	// Sets s as a string builder needed for paginator
	var s strings.Builder

	// A final string that is used to pass styles onto it
	var finalStr string

	// Renders the header
	finalStr += "\n" + styling.HeaderStyle.Foreground(lipgloss.Color("#7D56F4")).Render("Projects")
	finalStr += styling.HeaderStyle.UnsetForeground().Render("Descriptions") + "\n\n"

	// Iterate over the individual projects in items
	// Using the paginator function GetSliceBounds
	start, end := p.paginator.GetSliceBounds(len(p.items))

	for i, item := range p.items[start:end] {

		styling.ItemStyle.UnsetForeground() // Unset the font color by default

		cursor := " " // Sets cursor as blank by default

		if p.cursor == i {
			cursor = p.spinner.View() // Sets cursor as the spinner
		}

		if cursor == p.spinner.View() { // Checks where cursor is selecting
			styling.ItemStyle.Foreground(lipgloss.Color("12")) // Set font color of Project Name if selected by cursor
		}

		// Returns the model as a string, starting with the cursor, the item, then description
		finalStr += cursor + " " + styling.ItemStyle.Render(item) + "  " + styling.ShortDescStyle.Render(p.descriptions[i]) + "\n\n"
	}

	// Sets a variable fullHelpView as a string to return our pages menu help view...
	// Which is managed automatically by the help package
	fullHelpView := (p.paginator.View() + "\n\n" + p.help.View(p.keys))

	// Sets the height as an int the counts all lines, even empty ones
	height := p.termHeight - strings.Count(finalStr, "\n") // Counts lines in final string

	height -= strings.Count(fullHelpView, "\n") // Counts lines in full help view

	// Counts all new lines left in terminal
	height -= strings.Count("0", "\n") + 2 // Adds buffer for full help description

	// Adds the helpview which includes the paginator to our string
	finalStr += strings.Repeat("\n", height) + styling.HelpBarStyle.Render(fullHelpView)

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(p.termWidth).Height(p.termHeight).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

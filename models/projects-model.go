// OPEN SANDBOX PROGRAM WITH THE GOALS OF ADDING ITEMS TO A LIST, AND BEING ABLE TO SELECT THEM TO PRESENT A DIFFERENT

// TASKS:
// Format height and Width better
// Spinner freezes after model switch
// Returns an error if no project selected and try to get view

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
)

// Sets the Projects as a struct
// This is our main model for the projects page
type ProjectViewModel struct {
	items, shortDesc, longDesc []string        // Each project with a short description
	cursor                     int             // Used to track the cursor's location
	keys                       keyMap          // Sets a keymap needed to use the help view
	help                       help.Model      // Sets help as a help.Model so we can add it automatically to the bottom of our model
	paginator                  paginator.Model // Adds page scrolling to bottom of page
	spinner                    spinner.Model   // Adds the spinner to be used as a cursor
	width                      int
	height                     int
	err                        error
}

// This function is run in main to start a new program
// It sets our previously defined model with values
func CreateProjectViewModel() ProjectViewModel {
	// Sets items and descriptions new so we can change them easier here, and return them later
	var items, shortDesc, longDesc []string

	longDesc = []string{"This is a dank long description my homie"}

	//items = []string{"Pirates of the Cryptobbean", "Haramgay", "Another Dank Project here", "Midget Wrestling"}
	//shortDesc = []string{"Dank Pirates", "Gay Harambe NFT's", "Description of Dank Project", "Is Badass"}

	for i := 1; i < 36; i++ {
		text := fmt.Sprintf("Project: %d", i)
		desc := fmt.Sprintf("Short Description: %d", i)
		items = append(items, text)
		shortDesc = append(shortDesc, desc)
	}

	// SPINNER -- Sets up our spinner
	s := spinner.New() // Sets s as a new spinner
	s.Tick()
	s.Tick()
	s.Spinner = spinner.Dot                                         // Sets the dots style spinner
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205")) // Sets the spinner color

	// PAGINATOR -- Initializes the page scrolling for our list of items
	p := paginator.New()    // Sets p as a new paginator we can return later
	p.Type = paginator.Dots // Renders dots for our itmes
	p.PerPage = 8           // Items per page
	p.ActiveDot = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•") // Selected page formatting
	p.InactiveDot = lipgloss.NewStyle().Foreground(
		lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•") // Non-selected pages formatting
	p.SetTotalPages(len(items)) // Sets the total number of pages
	p.KeyMap.NextPage.Unbind()  // Unbinds the Next page command so we can customize it ourselves
	p.KeyMap.PrevPage.Unbind()  // Unbinds the Prev page command so we can customize it ourselves

	// Returns our model
	return ProjectViewModel{
		items:     items,
		shortDesc: shortDesc,
		longDesc:  longDesc,
		keys:      keys,
		help:      help.New(),
		paginator: p,
		spinner:   s,
	}
}

// ********************** BUBBLE TEA BUILT IN FUNCTIONS ***********************
// Initializes the model at start of program.
// Returns a command if there is one
func (p ProjectViewModel) Init() tea.Cmd {
	return p.spinner.Tick // Sets up the spinner
}

// Runs whenever there is an update or event
func (p ProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Sets the msg to types
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		// Sets the height and width of terminal so the border shows correctly
		p.help.Width = msg.Width
		p.width = msg.Width - 2
		p.height = msg.Height - 2

	case spinner.TickMsg: // update the spinner
		s, cmd := p.spinner.Update(msg)
		s.Tick()
		p.spinner = s
		return p, cmd

	// Handles key press events
	case tea.KeyMsg:

		// Converts the messages to string so we can see which key was pressed
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return p, tea.Quit

		// Moves the cursor up
		case "up":
			// If cursor above first item
			if p.cursor > -1 {
				// If on first page and cursor on first item, do nothing
				if p.paginator.Page == 0 && p.cursor == 0 {

					// Otherwise move cursor up
				} else {
					p.cursor--
				}
			}
			// Moving cursor up whn on first item of page that is not the first, change page to previous one
			if p.cursor == -1 {
				p.paginator.PrevPage()
				p.cursor = 7
			}

		// Moves the cursor down
		case "down":
			// If cursor below last item on page
			if p.cursor < 8 {
				// Checks if on last page
				if p.paginator.OnLastPage() == true {
					// Only moves cursor down if cursor is not on last item
					if p.cursor < (len(p.items)%p.paginator.PerPage - 1) {
						p.cursor++
					}

				} else { // If not on last page move the cursor down
					p.cursor++
				}

			}

			if p.cursor == 8 { // Move cursor to next page if on last item of page
				p.paginator.NextPage()
				p.cursor = 0
			}
		// Move to Previous page
		case "left":
			p.paginator.PrevPage()

		// Moves to next page
		case "right":
			p.paginator.NextPage()

			if p.paginator.OnLastPage() && p.cursor > 3 { // Moves cursor to last item if there are item ...
				p.cursor = 2 // 						..Not enough items on last page
			}

		// Toggles the help view between mini and full view
		case "?":
			p.help.ShowAll = !p.help.ShowAll

		// Switches to description page model for the selected project if space is pressed
		case " ":
			if p.cursor > len(p.items)%p.paginator.PerPage && p.paginator.OnLastPage() { // Error handling
				return p, nil
			} else {
				return CreateDescriptionModel(p.items[p.cursor+(p.paginator.Page*p.paginator.PerPage)], p.cursor), p.spinner.Tick
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

	finalStr += "\n"

	// Renders the header
	finalStr += styling.HeaderStyle.Foreground(lipgloss.Color("#7D56F4")).Render("Projects")
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
		finalStr += cursor + " " + styling.ItemStyle.Render(item) + "  " + styling.ShortDescStyle.Render(p.shortDesc[i]) + "\n\n"
	}

	// Sets a variable fullHelpView as a string to return our pages menu help view...
	// Which is managed automatically by the help package
	fullHelpView := (p.paginator.View() + "\n\n" + p.help.View(p.keys))

	// Sets the height as an int the counts all lines, even empty ones
	height := 11 - strings.Count("0", "\n") - strings.Count(fullHelpView, "\n")

	// Adds the helpview which includes the paginator to our string
	finalStr += "\n" + strings.Repeat("\n", height) + fullHelpView

	// Runs our complete string through the border/background styling
	completeModel := styling.Background.Width(p.width).Height(p.height).Render(finalStr)

	//returns our completed model as a string
	s.WriteString(completeModel)
	return s.String()
}

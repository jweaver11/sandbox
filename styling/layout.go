package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for the Header
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	PaddingLeft(3).
	PaddingRight(8)

// Styling for the items on the list
var ItemStyle = lipgloss.NewStyle().
	Bold(true)

// Styling for the Short Descriptions of the items
var ShortDescStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true).
	PaddingLeft(6)

// Styling for full Descriptions for description model
var FullDescStyle = lipgloss.NewStyle().
	PaddingLeft(4).
	Faint(true)

// Active dot of the paginator
var ActiveDotStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"})

// Inactive dot of the paginator
var InactiveDotStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"})

// Spinner Style
var SpinnerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("12"))

// Formatting to run final model through to add padding
var Background = lipgloss.NewStyle().
	PaddingLeft(4)

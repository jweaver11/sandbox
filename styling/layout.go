package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for the Header
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4"))

// Styling for the items on the list
var ItemStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("12"))

// Styling for the Short Descriptions of the items
var ShortDescStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)

// Custom border for outside of the TUI
var FullBorder = lipgloss.Border{
	Top:         "._.:*:",
	Bottom:      "._.:*:",
	Left:        "|*",
	Right:       "|*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

var Background = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder())

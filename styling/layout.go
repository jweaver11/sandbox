package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for the Header
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	PaddingLeft(4).
	PaddingTop(1)
	//Foreground(lipgloss.Color("#7D56F4")).

// Styling for the items on the list
var ItemStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("12"))

// Styling for the Short Descriptions of the items
var ShortDescStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true).
	PaddingLeft(6)

var LongDescStyle = lipgloss.NewStyle().
	PaddingLeft(4)

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

// Border doesnt work correctly when models are re-loaded
var Background = lipgloss.NewStyle().
	//Border(lipgloss.DoubleBorder()).
	PaddingLeft(4)

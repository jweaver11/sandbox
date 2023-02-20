package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for title messages
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4"))

var ItemStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("12"))

// Styling for the Short Descriptions
var ShortDescStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)
	//Border(RightBorder)

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
	Border(lipgloss.RoundedBorder())

package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for title messages
var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	PaddingTop(1).
	PaddingLeft(1)

var ShortDescStyle = lipgloss.NewStyle().
	Italic(true).
	Faint(true)

var Border = lipgloss.Border{
	Top:         "._.:*:",
	Bottom:      "._.:*:",
	Left:        "|*",
	Right:       "|*",
	TopLeft:     "*",
	TopRight:    "*",
	BottomLeft:  "*",
	BottomRight: "*",
}

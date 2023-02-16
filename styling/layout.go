package styling

import (
	"github.com/charmbracelet/lipgloss"
)

// Styling for title messages
var TitleStyle = lipgloss.NewStyle().
	SetString("WHAT YOU WANT\n").
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	//Background(lipgloss.Color("#7D56F4")).
	PaddingTop(1).
	PaddingLeft(1).
	Width(1)

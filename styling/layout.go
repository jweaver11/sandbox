package styling

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var Title = lipgloss.NewStyle().
	SetString("WHAT YOU WANT\n").
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	//Background(lipgloss.Color("#7D56F4")).
	PaddingTop(1).
	PaddingLeft(6).
	Width(22)

func PrintTitle() {
	fmt.Println(Title)
}

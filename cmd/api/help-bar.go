package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
)

// Sets a keymap struct to store the controls and key bind variables
// So they can be called on later for the help view
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}

// Sets keys as our object using our keyMap struct from above
var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),           // Activator keys for the struct
		key.WithHelp("↑", "move up"), // What help model shows
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "page left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "page right"),
	),
	Help: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q", "quit"),
	),
}

type PageModel struct {
	paginator paginator.Model
}

func createPageScrollModel() PageModel {
	var items []string
	for i := 1; i < 101; i++ {
		text := fmt.Sprintf("Item %d", i)
		items = append(items, text)
	}

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 10
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(items))

	return PageModel{
		paginator: p,
	}
}

func (o PageModel) View() string {
	var b strings.Builder
	b.WriteString("\n  Paginator Example\n\n")
	b.WriteString("  " + o.paginator.View())
	b.WriteString("\n\n  h/l ←/→ page • q: quit\n")
	return b.String()
}

// Main
//p := tea.NewProgram(newModel())
//if _, err := p.Run(); err != nil {
//log.Fatal(err)
//}

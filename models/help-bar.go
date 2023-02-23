package main

import (
	"github.com/charmbracelet/bubbles/key"
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
	Space key.Binding
}

// Built in function from the help package that shows our mini help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Space, k.Help, k.Quit}
}

// Built in function from the help package that shows our full help view at the bottom of our active model
// It is part of the key.Map interface
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Space, k.Help, k.Quit},       // second column
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
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("q", "quit"),
	),
	Space: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("space", "description"),
	),
}

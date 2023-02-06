package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4

// Three incrementing variables for the columns
const (
	todo status = iota
	inProgress
	done
)

/* MODEL MANAGEMENT */
var models []tea.Model

// incrementing model and form as int
const (
	model status = iota
	form
)

/* STYLING */
var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

// Sets the task struct layout for the task on todo list
type Task struct {
	status      status
	title       string
	description string
}

// Function that returns the new task
func NewTask(status status, title, description string) Task {
	return Task{status: status, title: title, description: description}
}

// Function used for controlling arrow keys to navigate correctly
func (t *Task) Next() {
	if t.status == done {
		t.status = todo
	} else {
		t.status++
	}
}

// implement the list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */
type Model struct {
	loaded   bool
	focused  status
	lists    []list.Model
	err      error
	quitting bool
}

// Creates a new model
func New() *Model {
	return &Model{}
}

// Moves task using 'enter' to the right
func (m *Model) MoveToNext() tea.Msg {
	selectedItem := m.lists[m.focused].SelectedItem()
	selectedTask := selectedItem.(Task)
	m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
	selectedTask.Next()
	m.lists[selectedTask.status].InsertItem(len(m.lists[selectedTask.status].Items())-1, list.Item(selectedTask))
	return nil
}

// Helper function to delete selected task
func (m *Model) DeleteCurrent() tea.Msg {
	if len(m.lists[m.focused].VisibleItems()) > 0 {
		selectedTask := m.lists[m.focused].SelectedItem().(Task)
		m.lists[selectedTask.status].RemoveItem(m.lists[m.focused].Index())
	}
	return nil
}

// Helper Movment function if selected task is on edge of lists
func (m *Model) Next() {
	if m.focused == done {
		m.focused = todo
	} else {
		m.focused++
	}
}

// Helper Movement function if selected task is on edge of lists
func (m *Model) Prev() {
	if m.focused == todo {
		m.focused = done
	} else {
		m.focused--
	}
}

// Initializes three lists with their tasks
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// Init To Do
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "get gud at coding", description: "stop sucking"},
		Task{status: todo, title: "go to jim", description: "get huge"},
	})
	// Init in progress
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: inProgress, title: "write dank code", description: "in GoLangs"},
	})
	// Init done
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "eat lunch", description: "be a fatty"},
	})
}

// Initialize helper
func (m Model) Init() tea.Cmd {
	return nil
}

// Adds in key commands to navigate UI and update what is shown to user
// Update function used whenever an event is triggered. Scans every frame for events
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Event that starts when loaded
	case tea.WindowSizeMsg:
		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	// Event that a ke is pressed
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left":
			m.Prev()
		case "right":
			m.Next()
		case "enter":
			return m, m.MoveToNext
		case "n":
			models[model] = m // save the state of the current model
			models[form] = NewForm(m.focused)
			return models[form].Update(nil)
		case "d":
			return m, m.DeleteCurrent
		// Uses the help.go file to load help screen
		case "h":
			if os.Getenv("HELP_DEBUG") != "" {
				if f, err := tea.LogToFile("debug.log", "help"); err != nil {
					fmt.Println("Couldn't open a file for logging:", err)
					os.Exit(1)
				} else {
					defer f.Close()
				}
			}

			if _, err := tea.NewProgram(newModel()).Run(); err != nil {
				fmt.Printf("Could not start program :(\n%v\n", err)
				os.Exit(1)
			}
		}
	// Event that involves the task struct
	case Task:
		task := msg
		return m, m.lists[task.status].InsertItem(len(m.lists[task.status].Items()), task)
	}
	// Sets cmd as a tea.Cmd
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

// Renders for user to view the UI
func (m Model) View() string {
	if m.quitting {
		return "bye bye bozo \n"
	}
	if m.loaded {
		// Loads in the views of the created lists
		todoView := m.lists[todo].View()
		inProgView := m.lists[inProgress].View()
		doneView := m.lists[done].View()

		// Loads in the focused views for each list
		switch m.focused {
		case inProgress:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				focusedStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		case done:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(todoView),
				columnStyle.Render(inProgView),
				focusedStyle.Render(doneView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(todoView),
				columnStyle.Render(inProgView),
				columnStyle.Render(doneView),
			)
		}
	} else {
		return "loading..."
	}
}

/* FORM MODEL */
type Form struct {
	focused     status
	title       textinput.Model
	description textarea.Model
}

func NewForm(focused status) *Form {
	form := &Form{focused: focused}
	form.title = textinput.New()
	form.title.Focus()
	form.description = textarea.New()
	return form
}

// Uses helper function to actually create the new task?
func (m Form) CreateTask() tea.Msg {
	task := NewTask(m.focused, m.title.Value(), m.description.Value())
	return task
}

// Initialize function?
func (m Form) Init() tea.Cmd {
	return nil
}

// Updates the UI in accordance with model?
func (m Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.description.Focus()
				return m, textarea.Blink
			} else {
				models[form] = m
				return models[model], m.CreateTask
			}
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else {
		m.description, cmd = m.description.Update(msg)
		return m, cmd
	}
}

// Formats Vertical view of the tasks using lipgloss
func (m Form) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.description.View())
}

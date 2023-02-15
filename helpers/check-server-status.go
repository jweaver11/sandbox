// Checks a serve to see if it recieves a response that it is up, or if it is down returns an error.
package helpers

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh"

type CheckServerModel struct {
	status int
	err    error
}

// make a request to the server, and return the result as a Msg
func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second} // New http client with a timeout of 10 seconds
	res, err := c.Get(url)                       // Result and error variable after getting our url

	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode) //returns the status message of the result stated earlier
}

type (
	statusMsg int
	errMsg    struct{ err error }
)

// implements the error interface
func (e errMsg) Error() string { return e.err.Error() }

func (m CheckServerModel) Init() tea.Cmd {
	return checkServer
}

func (m CheckServerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) { // makes a switch on the type of the message
	case statusMsg: // Handles the statusMsg created earlier
		m.status = int(msg)
		return m, tea.Quit

	case errMsg: // Handles when there is an error
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg: // Handles when a key is pressed
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m CheckServerModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nOur servers are currently unavailable %v\n\n", m.err)
	}
	// Sets s as a string that will be our return model for the status of the server
	// Currently it is set to do nothing if it successfully connects
	var s string
	// s = fmt.Sprintf("Connecting to %s...", url) //Sets a string as the checking text

	//if m.status > 0 { //if we have a response from status
	//s += fmt.Sprintf("%d %s! \n\n", m.status, http.StatusText(m.status)) // Adds the number from the status and returned string to our string s
	//}

	return s // Returns our string

}

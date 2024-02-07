package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MrShanks/pizza-delivery/pkg/menu"
	"github.com/MrShanks/pizza-delivery/pkg/restaurant"
	tea "github.com/charmbracelet/bubbletea"
)

const url = "http://localhost:3010/kitchen"

var order = restaurant.NewOrder()

type model struct {
	choices  []string // items on the to-do list
	pizzas   []restaurant.Pizza
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	status   int
	err      error
}

func initialModel() model {
	return model{
		choices:  []string{"Margherita", "Capricciosa", "Diavola", "Calabra"},
		pizzas:   []restaurant.Pizza{menu.Margherita(), menu.Capricciosa(), menu.Diavola(), menu.Calabra()},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		// The server returned a status message. Save it to our model. Also
		// tell the Bubble Tea runtime we want to exit because we have nothing
		// else to do. We'll still be able to render a final view with our
		// status message.
		m.status = int(msg)

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.
		m.err = msg
		return m, tea.Quit

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "s":
			SendOrder(&order)
			order = restaurant.NewOrder()
			for key, _ := range m.selected {
				delete(m.selected, key)
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
				pizza := m.pizzas[m.cursor]
				restaurant.RemovePizza(&order, pizza)
			} else {
				m.selected[m.cursor] = struct{}{}
				restaurant.AddPizza(&order, m.pizzas[m.cursor])
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Welcome to pizza-delivery restaurant, what pizza do you want?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress s to send the order\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func SendOrder(order *restaurant.Order) tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	ordJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}
	res, err := c.Post(url, "application/json", bytes.NewBuffer(ordJSON))

	if err != nil {
		return errMsg{err}
	}
	return statusMsg(res.StatusCode)
}

type statusMsg int

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

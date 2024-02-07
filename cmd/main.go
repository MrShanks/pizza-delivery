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
	choices  []string
	pizzas   []restaurant.Pizza
	cursor   int
	selected map[int]struct{}
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
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		m.status = int(msg)

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

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

	return m, nil
}

func (m model) View() string {
	s := "Welcome to pizza-delivery restaurant, what pizza do you want?\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress s to send the order\nPress q to quit.\n"

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

package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	NEW_GAME int = iota
	QUIT
)

type MenuOption struct {
	option int
	label  string
}

type MenuPageModel struct {
	options []MenuOption
	cursorY int
}

func newMenuPageModel() *MenuPageModel {
	return &MenuPageModel{
		options: []MenuOption{{NEW_GAME, "New game"}, {QUIT, "Quit"}},
		cursorY: 0,
	}
}

func (m *MenuPageModel) Init() tea.Cmd { return nil }

func (m *MenuPageModel) Update(msg tea.Msg) (Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down", "j":
			if m.cursorY < len(m.options)-1 {
				m.cursorY++
			}
		case "enter":
			choice := m.cursorY
			switch choice {
			case QUIT:
				return m, tea.Quit
			case NEW_GAME:
				return newGamePageModel(), nil
			}
		}
	}
	return m, nil
}

func (m *MenuPageModel) View() string {
	var b strings.Builder
	b.WriteString("Choose an action:\n\n")
	for _, o := range m.options {
		if o.option == m.cursorY {
			b.WriteString("> ")
		}
		b.WriteString(o.label)
		b.WriteByte('\n')
	}
	b.WriteString("\n\n")
	b.WriteString("Keys: j/k, up/down: select • enter: choose\n")
	return b.String()
}

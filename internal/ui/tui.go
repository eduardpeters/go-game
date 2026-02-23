package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page interface {
	Init() tea.Cmd
	Update(tea.Msg) (Page, tea.Cmd)
	View() string
}

type model struct {
	page Page
}

func NewModel() tea.Model {
	return model{
		page: newGamePageModel(),
	}
}

func (m model) Init() tea.Cmd { return m.page.Init() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newPage, cmd := m.page.Update(msg)
	if newPage != m.page {
		m.page = newPage
		initCmd := m.page.Init()
		return m, tea.Batch(initCmd)
	}

	return m, cmd
}

func (m model) View() string {
	return m.page.View()
}

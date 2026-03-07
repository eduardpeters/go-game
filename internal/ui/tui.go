package ui

import (
	tea "charm.land/bubbletea/v2"
)

type Page interface {
	Init() tea.Cmd
	Update(tea.Msg) (Page, tea.Cmd)
	View() tea.View
}

type model struct {
	page Page
}

func NewModel() tea.Model {
	return &model{
		page: newMenuPageModel(),
	}
}

func (m *model) Init() tea.Cmd { return m.page.Init() }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	newPage, cmd := m.page.Update(msg)

	m.page = newPage
	if m.page != newPage {
		return m, m.page.Init()
	}

	return m, cmd
}

func (m *model) View() tea.View {
	return m.page.View()
}

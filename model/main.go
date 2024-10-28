package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	main  tea.Model
	pager pagerModel
}

func New(title, content string) model {
	pager := pagerModel{
		title:   title,
		content: content,
	}
	return model{
		main:  pager,
		pager: pager,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.main.Update(msg)
}

func (m model) View() string {
	return m.main.View()
}

package model

import (
	"github.com/apainintheneck/wpedia/model/list"
	"github.com/apainintheneck/wpedia/model/pager"
	"github.com/apainintheneck/wpedia/wiki"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	isList bool
	list   list.Model
	pager  pager.Model
}

func New(title string, choices []string) Model {
	return Model{
		isList: true,
		list:   list.New(title, choices),
		pager:  pager.Model{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.isList {
		pager, cmd := m.pager.Update(msg)

		if pager.Title == "" && pager.Content == "" {
			m.isList = true
			m.list.Choice = ""
			return m, nil
		}

		m.pager = pager
		return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	if title := m.list.Choice; title != "" {
		m.isList = false
		m.pager.Title = title
		m.pager.Content = wiki.FetchContent(title)
		cmd = tea.WindowSize()
	}

	return m, cmd
}

func (m Model) View() string {
	if m.isList {
		return m.list.View()
	}

	return m.pager.View()
}

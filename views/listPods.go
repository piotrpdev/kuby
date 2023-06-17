package views

import (
	"errors"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
)

type ListPodsModel struct {
	Table     table.Model
	Altscreen bool
}

func (m ListPodsModel) Init() tea.Cmd { return nil }

func (m ListPodsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "q":
			return m, utils.BackToMainMenu
		case "enter":
			return m, utils.CreateChangeModel(ConnectErrorModel{Error: errors.New("hello there")})
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m ListPodsModel) View() string {
	return utils.BaseStyle.Render(m.Table.View()) + "\n" + utils.Subtle("up/down: select") + utils.Dot + utils.Subtle("enter: choose") + utils.Dot + utils.Subtle("q: go back") + "\n"
}

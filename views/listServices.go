package views

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
)

type ListServicesModel struct {
	Table     table.Model
	Altscreen bool
}

func (m ListServicesModel) Init() tea.Cmd { return nil }

func (m ListServicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "enter":
			//return m, tea.Batch(
			//	tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			//)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m ListServicesModel) View() string {
	return utils.BaseStyle.Render(m.Table.View()) + "\n" + utils.Subtle("up/down: select") + utils.Dot + utils.Subtle("enter: choose") + utils.Dot + utils.Subtle("q: go back") + "\n"
}

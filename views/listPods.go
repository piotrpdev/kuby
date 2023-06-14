package views

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type ListPodsModel struct {
	Table table.Model
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
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m ListPodsModel) View() string {
	return baseStyle.Render(m.Table.View()) + "\n" + subtle("up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, ctrl+c: quit") + "\n"
}

package utils

import tea "github.com/charmbracelet/bubbletea"

type ChangeModelMsg tea.Model

func CreateChangeModel(model tea.Model) func() tea.Msg {
	return func() tea.Msg {
		return ChangeModelMsg(model)
	}
}

type BackToMainMenuMsg struct{}

func BackToMainMenu() tea.Msg {
	return BackToMainMenuMsg{}
}

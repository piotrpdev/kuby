package views

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
)

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m MainMenuModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if msg.String() == "enter" {
				if i, ok := m.List.SelectedItem().(MainMenuItem); ok {
					m.Chosen = true
					m.Choice = i.GetModel()
					return m, m.Choice.Init()
					//fmt.Println(i.Title())
				}
			}
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

// Update loop for the second view after a choice has been made
func updateChosen(msg tea.Msg, m MainMenuModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			m.Chosen = false
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.Choice, cmd = m.Choice.Update(msg)

	return m, cmd
}

type MainMenuItem struct {
	TitleString, DescString string
	GetModel                func() tea.Model
}

func (i MainMenuItem) Title() string       { return i.TitleString }
func (i MainMenuItem) Description() string { return i.DescString }
func (i MainMenuItem) FilterValue() string { return i.TitleString }

type MainMenuModel struct {
	List   list.Model
	Choice tea.Model
	Chosen bool
}

func (m MainMenuModel) Init() tea.Cmd {
	// TODO: Check if this causes issues
	return tea.EnterAltScreen
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := utils.AppStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.Chosen {
		m.List, _ = m.List.Update(msg)
		return updateChoices(msg, m)
	} else {
		return updateChosen(msg, m)
	}
}

func (m MainMenuModel) View() string {
	var s string

	if !m.Chosen {
		s = utils.AppStyle.Render(m.List.View())
	} else {
		s = m.Choice.View()
	}

	return s
}

func NewMainMenuModel(items *[]list.Item) MainMenuModel {
	delegateKeys := utils.DelegateKeyMap{KeyBindings: []key.Binding{
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}}

	// Setup list
	delegate := utils.NewItemDelegate(&delegateKeys)
	mainMenuList := list.New(*items, delegate, 0, 0)
	mainMenuList.Title = "Main Menu"
	mainMenuList.Styles.Title = utils.TitleStyle

	return MainMenuModel{
		List:   mainMenuList,
		Chosen: false,
	}
}

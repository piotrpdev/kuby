package views

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"kuby/utils"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type MainMenuItem struct {
	TitleString, DescString string
}

func (i MainMenuItem) Title() string       { return i.TitleString }
func (i MainMenuItem) Description() string { return i.DescString }
func (i MainMenuItem) FilterValue() string { return i.TitleString }

func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return m.NewStatusMessage(utils.StatusMessageStyle("You chose"))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type DelegateKeyMap struct {
	choose key.Binding
}

// ShortHelp Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d DelegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
	}
}

// FullHelp Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d DelegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
		},
	}
}

func NewDelegateKeyMap() *DelegateKeyMap {
	return &DelegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}
}

type MainMenuModel struct {
	List         list.Model
	DelegateKeys *DelegateKeyMap
}

func (m MainMenuModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd

	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m MainMenuModel) View() string {
	return docStyle.Render(m.List.View())
}

func NewMainMenuModel(items *[]list.Item) MainMenuModel {
	var delegateKeys = NewDelegateKeyMap()

	// Setup list
	delegate := NewItemDelegate(delegateKeys)
	mainMenuList := list.New(*items, delegate, 0, 0)
	mainMenuList.Title = "Main Menu"
	mainMenuList.Styles.Title = utils.TitleStyle

	return MainMenuModel{
		List:         mainMenuList,
		DelegateKeys: delegateKeys,
	}
}

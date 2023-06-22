package views

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
	"strings"
	"testing"
)

type testModel struct{}

func (m testModel) Init() tea.Cmd {
	return nil
}

func (m testModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m testModel) View() string {
	return "Test View"
}

var testMainMenuItems = []list.Item{
	MainMenuItem{TitleString: "List Pods", DescString: "The smallest and simplest Kubernetes object.", GetModel: func(_ *MainMenuModel) tea.Model {
		return testModel{}
	}},
	MainMenuItem{TitleString: "List Services", DescString: "A method for exposing a network application that is running as one or more Pods in your cluster.", GetModel: func(_ *MainMenuModel) tea.Model {
		return testModel{}
	}},
	MainMenuItem{TitleString: "List Endpoints", DescString: "Network endpoint, typically referenced by a Service to define which Pods the traffic can be sent to.", GetModel: func(_ *MainMenuModel) tea.Model {
		return testModel{}
	}},
}

var testMainMenuModel = MainMenuModel{
	List:        list.New(testMainMenuItems, list.DefaultDelegate{}, 100, 100),
	Choice:      nil,
	Chosen:      false,
	ViewHistory: []*tea.Model{},
	Width:       100,
	Height:      100,
}

func getTestViewHistory() []*tea.Model {
	// https://stackoverflow.com/a/23172457/19020549
	model1 := interface{}(testModel{}).(tea.Model)
	model2 := interface{}(testModel{}).(tea.Model)
	model3 := interface{}(testModel{}).(tea.Model)

	return []*tea.Model{&model1, &model2, &model3}
}

var testMainMenuModelWithChoice = MainMenuModel{
	List:        list.New(testMainMenuItems, list.DefaultDelegate{}, 100, 100),
	Choice:      testModel{},
	Chosen:      true,
	ViewHistory: getTestViewHistory(),
	Width:       100,
	Height:      100,
}

func TestMainMenuItem_Description(t *testing.T) {
	type fields struct {
		TitleString string
		DescString  string
		GetModel    func(currentModel *MainMenuModel) tea.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Returns the description string", fields{TitleString: "Test Title", DescString: "Test Description", GetModel: func(currentModel *MainMenuModel) tea.Model { return testModel{} }}, "Test Description"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MainMenuItem{
				TitleString: tt.fields.TitleString,
				DescString:  tt.fields.DescString,
				GetModel:    tt.fields.GetModel,
			}
			if got := i.Description(); got != tt.want {
				t.Errorf("Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainMenuItem_FilterValue(t *testing.T) {
	type fields struct {
		TitleString string
		DescString  string
		GetModel    func(currentModel *MainMenuModel) tea.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Returns the title string", fields{TitleString: "Test Title", DescString: "Test Description", GetModel: func(currentModel *MainMenuModel) tea.Model { return testModel{} }}, "Test Title"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MainMenuItem{
				TitleString: tt.fields.TitleString,
				DescString:  tt.fields.DescString,
				GetModel:    tt.fields.GetModel,
			}
			if got := i.FilterValue(); got != tt.want {
				t.Errorf("FilterValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainMenuItem_Title(t *testing.T) {
	type fields struct {
		TitleString string
		DescString  string
		GetModel    func(currentModel *MainMenuModel) tea.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Returns the title string", fields{TitleString: "Test Title", DescString: "Test Description", GetModel: func(currentModel *MainMenuModel) tea.Model { return testModel{} }}, "Test Title"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := MainMenuItem{
				TitleString: tt.fields.TitleString,
				DescString:  tt.fields.DescString,
				GetModel:    tt.fields.GetModel,
			}
			if got := i.Title(); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainMenuModel_Init(t *testing.T) {
	t.Run("Returns nil",
		func(t *testing.T) {
			if got := testMainMenuModel.Init(); got != nil {
				t.Errorf("Init() did not return %v", nil)
			}
		})
}

func TestMainMenuModel_Update(t *testing.T) {
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name     string
		args     args
		function func(t *testing.T, args args, want tea.Model, want1 tea.Cmd)
		want     tea.Model
		want1    tea.Cmd
	}{
		{"Return nil when nil is passed", args{nil}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			if _, cmd := testMainMenuModel.Update(args.msg); cmd != nil {
				t.Errorf("Update() didn't return null, want %v", nil)
			}
		}, nil, nil},
		{"Return BackToMainMenuMsg when Quit key passed", args{tea.KeyMsg{
			Type:  tea.KeyCtrlC,
			Runes: nil,
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testMainMenuModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(tea.QuitMsg)
			if !ok {
				t.Errorf("Update() didn't return a tea.Quit")
			}
		}, nil, nil},
		{"Return WindowSizeMsg when Enter key passed", args{tea.KeyMsg{
			Type:  tea.KeyEnter,
			Runes: nil,
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testMainMenuModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(tea.BatchMsg)
			if !ok {
				t.Errorf("Update() didn't return a tea.BatchMsg")
			}
		}, nil, nil},
		{"Change List size when window size changes is passed", args{utils.WindowSizeCmd(200, 200)()}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			m, _ := testMainMenuModel.Update(args.msg)

			modelHeight := m.(MainMenuModel).Height

			if modelHeight != 200 {
				t.Errorf("Update() didn't change model height, got %v, want %v", modelHeight, 200)
			}
		}, nil, nil},
		{"History size changes on BackInHistoryMsg", args{utils.BackInHistoryMsg{}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			m, _ := testMainMenuModelWithChoice.Update(args.msg)

			viewHistorySize := len(m.(MainMenuModel).ViewHistory)

			if viewHistorySize != 2 {
				t.Errorf("Update() didn't change view history size, got %v, want %v", viewHistorySize, 2)
			}
		}, nil, nil},
		{"History size changes on BackToMainMenuMsg", args{utils.BackToMainMenuMsg{}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			m, _ := testMainMenuModelWithChoice.Update(args.msg)

			viewHistorySize := len(m.(MainMenuModel).ViewHistory)

			if viewHistorySize != 0 {
				t.Errorf("Update() didn't change view history size, got %v, want %v", viewHistorySize, 0)
			}
		}, nil, nil},
		{"History size changes on ChangeModelMsg", args{utils.ChangeModelMsg{testMainMenuModelWithChoice.ViewHistory[0], testMainMenuModelWithChoice.ViewHistory[1]}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			m, _ := testMainMenuModelWithChoice.Update(args.msg)

			viewHistorySize := len(m.(MainMenuModel).ViewHistory)

			if viewHistorySize != 4 {
				t.Fatalf("Update() didn't change view history size, got %v, want %v", viewHistorySize, 4)
			}

			if m.(MainMenuModel).Choice != *testMainMenuModelWithChoice.ViewHistory[0] {
				t.Errorf("Update() didn't change model choice, got %v, want %v", m.(MainMenuModel).Choice, testMainMenuModelWithChoice.ViewHistory[0])
			}
		}, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				tt.function(t, tt.args, tt.want, tt.want1)
			})
	}
}

func TestMainMenuModel_View(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, viewString string, want int)
		want     int
	}{
		{"View has one 'List Pods' option", func(t *testing.T, viewString string, want int) {
			if got := strings.Count(viewString, "List Pods"); got != want {
				t.Errorf("View() has %v 'List Pods' options, want %v", got, want)
			}
		}, 1},
		{"View doesn't have a 'List Tests' option", func(t *testing.T, viewString string, want int) {
			if got := strings.Count(viewString, "List Tests"); got != want {
				t.Errorf("View() has %v 'List Tests' options, want %v", got, want)
			}
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				viewString := testMainMenuModel.View()
				tt.function(t, viewString, tt.want)
			})
	}
}

func TestNewMainMenuModel(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, mainMenuModel MainMenuModel)
	}{
		{"Choice is nil", func(t *testing.T, mainMenuModel MainMenuModel) {
			if mainMenuModel.Choice != nil {
				t.Errorf("mainMenuModel() Choice isn't nil")
			}
		}},
		{"Chosen is false", func(t *testing.T, mainMenuModel MainMenuModel) {
			if mainMenuModel.Chosen != false {
				t.Errorf("mainMenuModel() Chosen isn't false")
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				mainMenuModel := NewMainMenuModel(&testMainMenuItems)
				tt.function(t, mainMenuModel)
			})
	}
}

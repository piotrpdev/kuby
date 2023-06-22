package views

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
	"strings"
	"testing"
)

var testListPodsRows = []table.Row{
	{"0", "suntracker-deployment-79bb7888b8-62jlh", "default", "Running", "10.244.1.4", "172.18.0.2"},
	{"1", "ingress-nginx-admission-create-h4f9f", "ingress-nginx", "Succeeded", "10.244.1.3", "172.18.0.2"},
	{"2", "ingress-nginx-admission-patch-d9w9j", "ingress-nginx", "Succeeded", "10.244.1.2", "172.18.0.2"},
	{"3", "ingress-nginx-controller-5bb6b499dc-ss425", "ingress-nginx", "Running", "10.244.0.5", "172.18.0.4"},
	{"4", "coredns-5d78c9869d-hfnqn", "kube-system", "Running", "10.244.0.3", "172.18.0.4"},
}

var testListPodsModel = ListPodsModel{
	Table:     testTable,
	Help:      help.New(),
	Altscreen: true,
	Height:    30,
	Width:     30,
}

var testTable = table.New(
	table.WithColumns([]table.Column{
		{Title: "Index", Width: 50},
		{Title: "Name", Width: 50},
		{Title: "Namespace", Width: 50},
		{Title: "Phase", Width: 50},
		{Title: "Pod IP", Width: 50},
		{Title: "Host IP", Width: 50},
	}),
	table.WithRows(testListPodsRows),
	table.WithFocused(true),
	table.WithHeight(30),
)

func TestListPodsModel_Init(t *testing.T) {
	t.Run("Returns nil",
		func(t *testing.T) {
			if got := testListPodsModel.Init(); got != nil {
				t.Errorf("Init() did not return %v", nil)
			}
		})
}

func TestListPodsModel_Update(t *testing.T) {
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
			if _, cmd := testListPodsModel.Update(args.msg); cmd != nil {
				t.Errorf("Update() didn't return null, want %v", nil)
			}
		}, nil, nil},
		{"Return BackToMainMenuMsg when Quit key passed", args{tea.KeyMsg{
			Type:  tea.KeyType(-1),
			Runes: []rune{rune(113)}, // q
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testListPodsModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(utils.BackToMainMenuMsg)
			if !ok {
				t.Errorf("Update() didn't return a BackToMainMenuMsg")
			}
		}, nil, utils.BackToMainMenu},
		{"Return ChangeModelMsg when Choose key passed", args{tea.KeyMsg{
			Type:  tea.KeyEnter,
			Runes: nil,
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testListPodsModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(utils.ChangeModelMsg)
			if !ok {
				t.Errorf("Update() didn't return a BackToMainMenuMsg")
			}
		}, nil, utils.BackToMainMenu},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				tt.function(t, tt.args, tt.want, tt.want1)
			})
	}
}

func TestListPodsModel_View(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, viewString string, want int)
		want     int
	}{
		{"View has one Namespace column", func(t *testing.T, viewString string, want int) {
			if got := strings.Count(viewString, "Namespace"); got != want {
				t.Errorf("View() has %v columns, want %v", got, want)
			}
		}, 1},
		{"View doesn't have a Test column", func(t *testing.T, viewString string, want int) {
			if got := strings.Count(viewString, "Test"); got != want {
				t.Errorf("View() has %v columns, want %v", got, want)
			}
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				viewString := testListPodsModel.View()
				tt.function(t, viewString, tt.want)
			})
	}
}

//func TestPodListToRows(t *testing.T) {
//	type args struct {
//		pods *v1.PodList
//	}
//	tests := []struct {
//		name string
//		args args
//		want []table.Row
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := PodListToRows(tt.args.pods); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("PodListToRows() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestNewPodsTable(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, podsTable *table.Model, want int)
		want     int
	}{
		{"Should have 5 rows", func(t *testing.T, podsTable *table.Model, want int) {
			rowCount := len(podsTable.Rows())
			if rowCount != want {
				t.Errorf("NewPodsTable() rows count = %v, want %v", rowCount, want)
			}
		}, 5},
		{"Should have 6 columns", func(t *testing.T, podsTable *table.Model, want int) {
			columnCount := len(podsTable.Rows()[0]) // .Columns() is not a thing
			if columnCount != want {
				t.Errorf("NewPodsTable() column count = %v, want %v", columnCount, want)
			}
		}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				podsTable := NewPodsTable(testListPodsRows)
				tt.function(t, podsTable, tt.want)
			})
	}
}

package views

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
	"strings"
	"testing"
)

var testListEndpointsRows = []table.Row{
	{"0", "kubernetes", "default", "172.18.0.3", "6443"},
	{"1", "suntracker-service", "default", "10.244.1.2", "80"},
	{"2", "ingress-nginx-controller", "ingress-nginx", "10.244.0.5", "443, 80"},
	{"3", "ingress-nginx-controller-admission", "ingress-nginx", "10.244.0.5", "8443"},
	{"4", "kube-dns", "kube-system", "10.244.0.2, 10.244.0.3", "53, 53, 9153"},
}

var testListEndpointsModel = ListEndpointsModel{
	Table:     testListEndpointsTable,
	Help:      help.New(),
	Altscreen: true,
	Height:    30,
	Width:     30,
}

var testListEndpointsTable = table.New(
	table.WithColumns([]table.Column{
		{Title: "Index", Width: 50},
		{Title: "Name", Width: 50},
		{Title: "Namespace", Width: 50},
		{Title: "Addresses", Width: 50},
		{Title: "Ports", Width: 50},
	}),
	table.WithRows(testListEndpointsRows),
	table.WithFocused(true),
	table.WithHeight(30),
)

func TestListEndpointsModel_Init(t *testing.T) {
	t.Run("Returns nil",
		func(t *testing.T) {
			if got := testListEndpointsModel.Init(); got != nil {
				t.Errorf("Init() did not return %v", nil)
			}
		})
}

func TestListEndpointsModel_Update(t *testing.T) {
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
			if _, cmd := testListEndpointsModel.Update(args.msg); cmd != nil {
				t.Errorf("Update() didn't return null, want %v", nil)
			}
		}, nil, nil},
		{"Return BackToMainMenuMsg when Quit key passed", args{tea.KeyMsg{
			Type:  tea.KeyType(-1),
			Runes: []rune{rune(113)}, // q
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testListEndpointsModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(utils.BackToMainMenuMsg)
			if !ok {
				t.Errorf("Update() didn't return a BackToMainMenuMsg")
			}
		}, nil, nil},
		{"Return ChangeModelMsg when Choose key passed", args{tea.KeyMsg{
			Type:  tea.KeyEnter,
			Runes: nil,
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Model, want1 tea.Cmd) {
			_, cmd := testListEndpointsModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(utils.ChangeModelMsg)
			if !ok {
				t.Errorf("Update() didn't return a BackToMainMenuMsg")
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

func TestListEndpointsModel_View(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, viewString string, want int)
		want     int
	}{
		{"View has one Addresses column", func(t *testing.T, viewString string, want int) {
			if got := strings.Count(viewString, "Addresses"); got != want {
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
				viewString := testListEndpointsModel.View()
				tt.function(t, viewString, tt.want)
			})
	}
}

func TestNewEndpointsTable(t *testing.T) {
	tests := []struct {
		name     string
		function func(t *testing.T, endpointsTable *table.Model, want int)
		want     int
	}{
		{"Should have 5 rows", func(t *testing.T, endpointsTable *table.Model, want int) {
			rowCount := len(endpointsTable.Rows())
			if rowCount != want {
				t.Errorf("NewendpointsTable() rows count = %v, want %v", rowCount, want)
			}
		}, 5},
		{"Should have 6 columns", func(t *testing.T, endpointsTable *table.Model, want int) {
			columnCount := len(endpointsTable.Rows()[0]) // .Columns() is not a thing
			if columnCount != want {
				t.Errorf("NewendpointsTable() column count = %v, want %v", columnCount, want)
			}
		}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				endpointsTable := NewEndpointsTable(testListEndpointsRows)
				tt.function(t, endpointsTable, tt.want)
			})
	}
}

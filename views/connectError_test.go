package views

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/utils"
	"strings"
	"testing"
)

var testConnectErrorModel = ConnectErrorModel{
	Error: nil,
}

func TestConnectErrorModel_Init(t *testing.T) {
	t.Run("Returns nil",
		func(t *testing.T) {
			if got := testListPodsModel.Init(); got != nil {
				t.Errorf("Init() did not return %v", nil)
			}
		})
}

func TestConnectErrorModel_Update(t *testing.T) {
	type args struct {
		msg tea.Msg
	}
	tests := []struct {
		name     string
		args     args
		function func(t *testing.T, args args, want tea.Cmd)
		want     tea.Cmd
	}{
		{"Return nil when nil is passed", args{nil}, func(t *testing.T, args args, want tea.Cmd) {
			if _, cmd := testConnectErrorModel.Update(args.msg); cmd != nil {
				t.Errorf("Update() didn't return null, want %v", nil)
			}
		}, nil},
		{"Return BackToMainMenuMsg when Quit key passed", args{tea.KeyMsg{
			Type:  tea.KeyType(-1),
			Runes: []rune{rune(113)}, // q
			Alt:   false,
		}}, func(t *testing.T, args args, want tea.Cmd) {
			_, cmd := testConnectErrorModel.Update(args.msg)
			if cmd == nil {
				t.Fatalf("Update() cmd is nil")
			}
			msg := cmd()

			_, ok := msg.(utils.BackInHistoryMsg)
			if !ok {
				t.Errorf("Update() didn't return a BackInHistoryMsg")
			}
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name,
			func(t *testing.T) {
				tt.function(t, tt.args, tt.want)
			})
	}
}

func TestConnectErrorModel_View(t *testing.T) {
	type fields struct {
		Error error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"Contains error", fields{errors.New("Test")}, "Test"},
		{"Contains Docker notice", fields{errors.New("Test")}, "Is Docker/K8s running?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ConnectErrorModel{
				Error: tt.fields.Error,
			}
			if !strings.Contains(m.View(), tt.want) {
				t.Errorf("View() does not contain %v", tt.want)
			}
		})
	}
}

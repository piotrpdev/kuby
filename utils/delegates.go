package utils

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type DelegateKeyMap struct {
	KeyBindings []key.Binding
}

func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.ShortHelpFunc = func() []key.Binding {
		return keys.KeyBindings
	}
	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{keys.KeyBindings}
	}

	return d
}

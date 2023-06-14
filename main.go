package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/k8s"
	"kuby/tables"
	"kuby/views"
	"os"
)

func main() {
	clientset := k8s.GetClientset()

	podsTable := tables.GetPodsTable(clientset)

	m := views.ListPodsModel{Table: podsTable}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/views"
	"os"
)

func main() {
	//clientset := k8s.GetClientset()
	//
	//podsTable := tables.GetPodsTable(clientset)
	//
	//m := views.ListPodsModel{Table: podsTable, Altscreen: true}

	items := []list.Item{
		views.MainMenuItem{TitleString: "List Pods", DescString: "The smallest and simplest Kubernetes object."},
		views.MainMenuItem{TitleString: "List Services", DescString: "A method for exposing a network application that is running as one or more Pods in your cluster."},
		views.MainMenuItem{TitleString: "List Endpoints", DescString: "Network endpoint, typically referenced by a Service to define which Pods the traffic can be sent to."},
	}

	m := views.NewMainMenuModel(&items)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

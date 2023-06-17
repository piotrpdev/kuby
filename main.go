package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"kuby/k8s"
	"kuby/tables"
	"kuby/views"
	"os"
)

func main() {
	clientset := k8s.GetClientset()

	// TODO: Maybe use a custom type for top-level views to force them to use `utils.BackToMainMenu`
	items := []list.Item{
		views.MainMenuItem{TitleString: "List Pods", DescString: "The smallest and simplest Kubernetes object.", GetModel: func(m *views.MainMenuModel) tea.Model {
			podsTable, err := tables.GetPodsTable(clientset)
			if err != nil {
				// TODO: Test this works with the new router/history system
				return views.ConnectErrorModel{Error: err}
			}

			return views.ListPodsModel{Table: *podsTable, Altscreen: true, Height: m.List.Height(), Width: m.List.Width()}
		}},
		views.MainMenuItem{TitleString: "List Services", DescString: "A method for exposing a network application that is running as one or more Pods in your cluster.", GetModel: func(m *views.MainMenuModel) tea.Model {
			servicesTable, err := tables.GetServicesTable(clientset)
			if err != nil {
				return views.ConnectErrorModel{Error: err}
			}

			return views.ListServicesModel{Table: *servicesTable, Altscreen: true, Height: m.List.Height(), Width: m.List.Width()}
		}},
		views.MainMenuItem{TitleString: "List Endpoints", DescString: "Network endpoint, typically referenced by a Service to define which Pods the traffic can be sent to.", GetModel: func(m *views.MainMenuModel) tea.Model {
			endpointsTable, err := tables.GetEndpointsTable(clientset)
			if err != nil {
				return views.ConnectErrorModel{Error: err}
			}

			return views.ListEndpointsModel{Table: *endpointsTable, Altscreen: true, Height: m.List.Height(), Width: m.List.Width()}
		}},
	}

	//pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//niceJson, _ := json.MarshalIndent(pods.Items[0], "", "    ")
	//fmt.Println(string(niceJson))

	m := views.NewMainMenuModel(&items)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

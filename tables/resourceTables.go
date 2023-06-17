package tables

import (
	"context"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"strings"
)

func GetPodsTable(clientset *kubernetes.Clientset) (*table.Model, error) {
	// TODO: Maybe move the api call somewhere else and add pods as param e.g. separate concerns
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	rows := lop.Map(pods.Items, func(item v1.Pod, index int) table.Row {
		return table.Row{strconv.Itoa(index), item.ObjectMeta.Name, item.ObjectMeta.Namespace, string(item.Status.Phase), item.Status.PodIP, item.Status.HostIP}
	})

	columns := []table.Column{
		{Title: "Index", Width: 5},
		{Title: "Name", Width: longestInColumn(&rows, 1)},
		{Title: "Namespace", Width: longestInColumn(&rows, 2)},
		{Title: "Phase", Width: longestInColumn(&rows, 3)},
		{Title: "Pod IP", Width: longestInColumn(&rows, 4)},
		{Title: "Host IP", Width: longestInColumn(&rows, 5)},
		//{Title: "Created", Width: longestInColumn(&rows, 6)},
		//{Title: "Started", Width: longestInColumn(&rows, 7)},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &t, err
}

func GetServicesTable(clientset *kubernetes.Clientset) (*table.Model, error) {
	// TODO: Maybe move the api call somewhere else and add services as param e.g. separate concerns
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	rows := lop.Map(services.Items, func(item v1.Service, index int) table.Row {
		ports := lop.Map(item.Spec.Ports, func(item v1.ServicePort, _ int) string {
			return strconv.Itoa(int(item.Port))
		})
		return table.Row{strconv.Itoa(index), item.ObjectMeta.Name, item.ObjectMeta.Namespace, item.Spec.ClusterIP, strings.Join(ports, ", ")}
	})

	columns := []table.Column{
		{Title: "Index", Width: 5},
		{Title: "Name", Width: longestInColumn(&rows, 1)},
		{Title: "Namespace", Width: longestInColumn(&rows, 2)},
		{Title: "Cluster IP", Width: longestInColumn(&rows, 3)},
		{Title: "Ports", Width: longestInColumn(&rows, 4)},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &t, err
}

func GetEndpointsTable(clientset *kubernetes.Clientset) (*table.Model, error) {
	// TODO: Maybe move the api call somewhere else and add endpoints as param e.g. separate concerns
	endpoints, err := clientset.CoreV1().Endpoints("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	rows := lop.Map(endpoints.Items, func(item v1.Endpoints, index int) table.Row {
		addresses := lo.FlatMap(item.Subsets, func(item v1.EndpointSubset, _ int) []string {
			return lop.Map(item.Addresses, func(item v1.EndpointAddress, _ int) string {
				return item.IP
			})
		})
		ports := lo.FlatMap(item.Subsets, func(item v1.EndpointSubset, _ int) []string {
			return lop.Map(item.Ports, func(item v1.EndpointPort, _ int) string {
				return strconv.Itoa(int(item.Port))
			})
		})
		return table.Row{strconv.Itoa(index), item.ObjectMeta.Name, item.ObjectMeta.Namespace, strings.Join(addresses, ", "), strings.Join(ports, ", ")}
	})

	columns := []table.Column{
		{Title: "Index", Width: 5},
		{Title: "Name", Width: longestInColumn(&rows, 1)},
		{Title: "Namespace", Width: longestInColumn(&rows, 2)},
		{Title: "Adresses", Width: longestInColumn(&rows, 3)},
		{Title: "Ports", Width: longestInColumn(&rows, 4)},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &t, err
}

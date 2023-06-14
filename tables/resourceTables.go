package tables

import (
	"context"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	lop "github.com/samber/lo/parallel"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"time"
)

func GetPodsTable(clientset *kubernetes.Clientset) table.Model {
	// TODO: Maybe move the api call somewhere else and add pods as param e.g. separate concerns
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	rows := lop.Map(pods.Items, func(item v1.Pod, index int) table.Row {
		return table.Row{strconv.Itoa(index), item.ObjectMeta.Name, item.ObjectMeta.Namespace, string(item.Status.Phase), item.Status.PodIP, item.Status.HostIP, item.ObjectMeta.CreationTimestamp.Format(time.RFC3339), item.Status.StartTime.Format(time.RFC3339)}
	})

	columns := []table.Column{
		{Title: "Index", Width: 5},
		{Title: "Name", Width: longestInColumn(&rows, 1)},
		{Title: "Namespace", Width: longestInColumn(&rows, 2)},
		{Title: "Phase", Width: longestInColumn(&rows, 3)},
		{Title: "Pod IP", Width: longestInColumn(&rows, 4)},
		{Title: "Host IP", Width: longestInColumn(&rows, 5)},
		{Title: "Created", Width: longestInColumn(&rows, 6)},
		{Title: "Started", Width: longestInColumn(&rows, 7)},
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

	return t
}

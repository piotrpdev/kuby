package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func getClientset() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func longestInColumn(rows *[]table.Row, columnIndex int) int {
	columnSlice := lo.Map(*rows, func(item table.Row, _ int) string {
		return item[columnIndex]
	})

	longestColumnValue := lo.MaxBy(columnSlice, func(item string, max string) bool {
		return len(item) > len(max)
	})

	return len(longestColumnValue)
}

func getPodsTable(clientset *kubernetes.Clientset) table.Model {
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

func main() {
	clientset := getClientset()

	t := getPodsTable(clientset)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

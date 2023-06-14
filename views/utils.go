package views

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

var term = termenv.EnvColorProfile()
var subtle = makeFgStyle("241")
var dot = colorFg(" • ", "236")

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

package tables

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/samber/lo"
)

func longestInColumn(rows *[]table.Row, columnIndex int) int {
	columnSlice := lo.Map(*rows, func(item table.Row, _ int) string {
		return item[columnIndex]
	})

	longestColumnValue := lo.MaxBy(columnSlice, func(item string, max string) bool {
		return len(item) > len(max)
	})

	return len(longestColumnValue)
}

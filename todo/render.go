package todo

import (
	"fmt"

	"bitbucket.org/abdullin/olympia/forms"
)

func renderMainForm(ts []*Task) *forms.Window {
	w := forms.NewWindow("Tasks ERP")
	w.AddNav("Dashboard", false)
	w.AddNav(fmt.Sprintf("Tasks (%d)", len(ts)), true)
	w.AddNav("Audit", false)
	w.AddMenu("Admin")
	w.AddMenu("Help")
	w.Content = renderGrid(ts)
	return w
}

func renderGrid(ts []*Task) *forms.DataGrid {

	grid := forms.NewGrid()
	grid.AddTextColumn("#")
	grid.AddTextColumn("Priority")
	grid.AddTextColumn("Name")
	for _, task := range ts {
		grid.AddRow(task.ID, task.Priority, task.Name)
	}

	return grid
}

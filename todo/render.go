package todo

import "bitbucket.org/abdullin/olympia/forms"

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

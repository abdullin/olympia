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

func renderAddTaskScreen() interface{} {
	form := forms.NewForm()
	form.AddTextField("Name", "New task")
	form.AddSelectField("Priority", "Minor", "Major", "Blocker")

	w := forms.NewDialog("Add Task", form)

	args := make(map[string]interface{})
	args["name"] = "New Task"
	args["priority"] = "Major"

	w.AddAcceptButton("Add Task", "add-task-accept", args)
	w.AddCancelButton("Cancel", "add-task-cancel")

	return w

}

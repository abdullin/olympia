package todo

import (
	"fmt"

	"github.com/abdullin/olympia/forms"
)

func mainMenu(w *forms.Window) {
	w.AddMenu("Exercise", forms.NewAction(actionShowGrid))
	w.AddMenu("Add Task", forms.NewAction("show-add-task-form"))
	w.AddMenu("Admin", forms.NewAction("go-admin"))
	w.AddMenu("Help", forms.NewAction("go-help"))
}

func renderTaskWindow(ts []*Task) *forms.Window {
	w := forms.NewWindow("Tasks ERP")
	w.AddNav("Dashboard", false, forms.NewAction("go-dashboard"))
	w.AddNav(fmt.Sprintf("Tasks (%d)", len(ts)), true, forms.NewAction("go-tasks"))
	w.AddNav("Audit", false, forms.NewAction("go-audit"))

	mainMenu(w)

	w.Content = renderTable(ts)
	return w
}

func cat() *forms.Image {
	return forms.NewImage("http://thecatapi.com/api/images/get?format=src&type=gif")
}

func renderExerciseWindow() *forms.Window {
	w := forms.NewWindow("1 / 10")
	mainMenu(w)

	w.Content = renderGrid()
	return w
}

func renderGrid() *forms.Grid {
	g := forms.NewGrid()
	r1 := g.AddRow()

	b1 := forms.NewButton("99", nil)
	b1.Large = true
	b1.Fill = true

	b1.Style = "danger"

	b2 := forms.NewButton("99", nil)
	b2.Large = true
	b2.Fill = true
	b2.Style = "success"

	r1.AddCol(b1)
	r1.AddCol(b2)

	r2 := g.AddRow()
	r2.AddCol(cat())
	r2.AddCol(cat())
	r2.AddCol(cat())
	return g
}

func renderTable(ts []*Task) *forms.DataTable {
	grid := forms.NewTable()
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

package todo

import (
	"fmt"

	"bitbucket.org/abdullin/olympia/forms"
)

type Task struct {
	Name string

	Priority string
	ID       int
}

type App struct {
	Screen  string
	Tasks   []*Task
	Changed chan bool
}

func Start() *App {
	// todo - load state from the DB
	changed := make(chan bool, 100)
	return &App{
		Screen:  "main",
		Changed: changed,
	}
}

func (a *App) AddTask(name string, priority string) {
	a.Tasks = append(a.Tasks, &Task{name, priority, len(a.Tasks)})
	a.Changed <- true
}

func (a *App) GetScreen() interface{} {

	switch a.Screen {
	case "main":

		w := forms.NewWindow("Tasks ERP")
		w.AddNav("Dashboard", false)
		w.AddNav(fmt.Sprintf("Tasks (%d)", len(a.Tasks)), true)
		w.AddNav("Audit", false)
		w.AddMenu("Admin")
		w.AddMenu("Help")
		w.Content = renderGrid(a.Tasks)
		return w

	}
	panic("Unknown state")

}

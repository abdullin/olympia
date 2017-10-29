package todo

import "fmt"

type Task struct {
	Name     string
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

func (a *App) Dispatch(action string, args map[string]interface{}) {
	switch action {
	case "show-add-task-form":
		a.Screen = "add-task-form"
		a.Changed <- true
	case "add-task-accept":
		name := args["name"].(string)
		priority := args["priority"].(string)
		a.Tasks = append(a.Tasks, &Task{name, priority, len(a.Tasks)})
		a.Screen = "main"
		a.Changed <- true
	case "add-task-cancel":
		a.Screen = "main"
		a.Changed <- true

	default:
		fmt.Println("Unknown dispatch", action, "with args", args)

	}
}

func (a *App) AddTask(name string, priority string) {
}

// GetScreen returns screen UI for the running app
// DEFER: we are re-rendering screen completely for now
func (a *App) GetScreen() interface{} {
	switch a.Screen {
	case "main":
		return renderMainForm(a.Tasks)
	case "add-task-form":
		return renderAddTaskScreen()
	}
	panic("Unknown state")

}

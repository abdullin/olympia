package exam

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

package forms

type Window struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
	Title   string      `json:"title"`
	Menu    []*MenuItem `json:"menu"`
	Nav     []*NavItem  `json:"nav"`
}

func NewWindow(title string) *Window {
	return &Window{Type: "window", Title: title,
		Menu: []*MenuItem{},
		Nav:  []*NavItem{},
	}
}

func (w *Window) AddNav(text string, active bool, action *Action) {
	var state string
	if active {
		state = "active"
	}

	w.Nav = append(w.Nav, &NavItem{Text: text, State: state, Action: action})
}

func (w *Window) AddNavDisabled(text string) {
	w.Nav = append(w.Nav, &NavItem{Text: text, State: "disabled"})
}

// Disabled

func (w *Window) AddMenu(text string, action *Action) {
	w.Menu = append(w.Menu, &MenuItem{text, true, action})
}

type MenuItem struct {
	Title  string  `json:"title"`
	Active bool    `json:"active"`
	Action *Action `json:"action"`
}

type NavItem struct {
	Text   string  `json:"text"`
	State  string  `json:"state"`
	Action *Action `json:"action"`
}

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

func (w *Window) AddNav(text string, active bool) {
	w.Nav = append(w.Nav, &NavItem{text, active})
}

func (w *Window) AddMenu(text string) {
	w.Menu = append(w.Menu, &MenuItem{text, true})
}

type MenuItem struct {
	Title  string `json:"title"`
	Active bool   `json:"active"`
}

type NavItem struct {
	Text   string `json:"text"`
	Active bool   `json:"active"`
}

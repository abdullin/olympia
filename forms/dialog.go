package forms

type Dialog struct {
	Type         string      `json:"type"`
	Title        string      `json:"title"`
	Content      interface{} `json:"content"`
	CancelButton *Button     `json:"cancelButton"`
	AcceptButton *Button     `json:"acceptButton"`
}

func NewDialog(title string, content interface{}) *Dialog {
	return &Dialog{Type: "dialog", Title: title, Content: content}
}

type Button struct {
	Text   string  `json:"text"`
	Action *Action `json:"action"`
}

func (d *Dialog) AddAcceptButton(text, actionName string, args map[string]interface{}) {
	d.AcceptButton = &Button{Text: text, Action: &Action{Type: actionName, Args: args}}
}

func (d *Dialog) AddCancelButton(text, actionName string) {
	d.CancelButton = &Button{Text: text, Action: &Action{Type: actionName}}
}

type Action struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

type Image struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

func NewImage(src string) *Image {
	return &Image{Type: "image", Src: src}
}

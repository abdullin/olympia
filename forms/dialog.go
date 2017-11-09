package forms

type Dialog struct {
	Type         string        `json:"type"`
	Title        string        `json:"title"`
	Content      interface{}   `json:"content"`
	CancelButton *DialogButton `json:"cancelButton"`
	AcceptButton *DialogButton `json:"acceptButton"`
}

func NewDialog(title string, content interface{}) *Dialog {
	return &Dialog{Type: "dialog", Title: title, Content: content}
}

type DialogButton struct {
	Text   string  `json:"text"`
	Action *Action `json:"action"`
}

func (d *Dialog) AddAcceptButton(text, actionName string, args map[string]interface{}) {
	d.AcceptButton = &DialogButton{Text: text, Action: &Action{Type: actionName, Args: args}}
}

func (d *Dialog) AddCancelButton(text, actionName string) {
	d.CancelButton = &DialogButton{Text: text, Action: &Action{Type: actionName}}
}

type Image struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

func NewImage(src string) *Image {
	return &Image{Type: "image", Src: src}
}

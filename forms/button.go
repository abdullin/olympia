package forms

type Button struct {
	Type   string  `json:"type"`
	Style  string  `json:"style"`
	Text   string  `json:"text"`
	Large  bool    `json:"large"`
	Action *Action `json:"action"`
	Fill   bool    `json:"fill"`
}

func NewButton(text string, action *Action) *Button {
	return &Button{
		Type:   "button",
		Text:   text,
		Action: action,
	}
}

type Label struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Level int    `json:"level"`
}

func NewLabel(text string) *Label {
	return &Label{
		Type: "label",
		Text: text,
	}
}

func NewTitle(text string, level int) *Label {
	return &Label{
		Type:  "label",
		Text:  text,
		Level: level,
	}
}

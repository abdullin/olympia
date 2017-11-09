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

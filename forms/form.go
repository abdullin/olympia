package forms

type FormTextField struct {
	Type  string `json:"type"`
	Hint  string `json:"hint"`
	Label string `json:"label"`
}

type FormSelectField struct {
	Type    string   `json:"type"`
	Label   string   `json:"label"`
	Options []string `json:"options"`
}

type Form struct {
	Type  string        `json:"type"`
	Items []interface{} `json:"items"`
}

func NewForm() *Form {
	return &Form{Items: []interface{}{}, Type: "form"}
}

func (f *Form) AddTextField(label, hint string) {
	f.Items = append(f.Items, &FormTextField{Type: "text", Label: label, Hint: hint})
}

func (f *Form) AddSelectField(label string, options ...string) {
	f.Items = append(f.Items, &FormSelectField{Type: "select", Label: label, Options: options})
}

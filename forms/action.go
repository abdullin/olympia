package forms

type Action struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

func NewAction(t string) *Action {
	return &Action{Type: t, Args: make(map[string]interface{})}
}

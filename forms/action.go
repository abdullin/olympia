package forms

type Action struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

func NewAction(t string) *Action {
	return &Action{Type: t, Args: make(map[string]interface{})}
}

type PieChart struct {
	Type   string     `json:"type"`
	Width  int        `json:"width"`
	Height int        `json:"height"`
	Data   []PiePoint `json:"data"`
}

type PiePoint struct {
	X string  `json:"x"`
	Y float32 `json:"y"`
}

func NewPieChart(data []PiePoint) *PieChart {

	return &PieChart{
		Type:   "pie-chart",
		Height: 200,
		Width:  200,
		Data:   data,
	}
}

type LineChart struct {
	Type   string      `json:"type"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Data   []LinePoint `json:"data"`
}

type LinePoint struct {
	X string  `json:"x"`
	Y float32 `json:"y"`
}

func NewLineChart(data []LinePoint) *LineChart {

	return &LineChart{
		Type:   "line-chart",
		Height: 200,
		Width:  200,
		Data:   data,
	}
}

package forms

type Grid struct {
	Type string `json:"type"`
	Rows []*Row `json:"rows"`
}

type Row struct {
	Cols []*Col `json:"cols"`
}

type Col struct {
	Content interface{} `json:"content"`
	Steps   int         `json:"steps"`
}

func NewGrid() *Grid {
	return &Grid{Type: "grid", Rows: []*Row{}}
}

func (g *Grid) AddRow() *Row {
	r := &Row{Cols: []*Col{}}
	g.Rows = append(g.Rows, r)
	return r
}

func (g *Grid) AddRowItems(cols ...interface{}) *Row {
	r := &Row{Cols: []*Col{}}

	for _, c := range cols {
		r.AddCol(c)
	}
	g.Rows = append(g.Rows, r)
	return r
}

func (r *Row) AddCol(content interface{}) *Col {
	c := &Col{Content: content}
	r.Cols = append(r.Cols, c)
	return c
}

func (r *Row) SetSpans(spans ...int) *Row {
	for i, s := range spans {
		r.Cols[i].Steps = s
	}
	return r
}

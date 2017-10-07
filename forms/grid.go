package forms

type DataGridColumn struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type DataGridRow struct {
	Cells []*DataGridCell `json:"cells"`
}

type DataGridCell struct {
	Value interface{} `json:"val"`
}

type DataGrid struct {
	Type      string            `json:"type"`
	CanDelete bool              `json:"canDelete"`
	CanEdit   bool              `json:"canEdit"`
	Rows      []*DataGridRow    `json:"rows"`
	Columns   []*DataGridColumn `json:"columns"`
}

func NewGrid() *DataGrid {
	return &DataGrid{Type: "data-grid", Rows: []*DataGridRow{}}
}

func (d *DataGrid) AddTextColumn(text string) {
	d.Columns = append(d.Columns, &DataGridColumn{Text: text})
}

func (d *DataGrid) AddRow(args ...interface{}) {
	cells := make([]*DataGridCell, len(args), len(args))

	for i := 0; i < len(args); i++ {
		cells[i] = &DataGridCell{Value: args[i]}
	}
	row := &DataGridRow{Cells: cells}
	d.Rows = append(d.Rows, row)

}

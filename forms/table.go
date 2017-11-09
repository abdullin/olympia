package forms

type DataTableColumn struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type DataTableRow struct {
	Cells []*DataTableCell `json:"cells"`
}

type DataTableCell struct {
	Value interface{} `json:"val"`
}

type DataTable struct {
	Type      string             `json:"type"`
	CanDelete bool               `json:"canDelete"`
	CanEdit   bool               `json:"canEdit"`
	Rows      []*DataTableRow    `json:"rows"`
	Columns   []*DataTableColumn `json:"columns"`
}

func NewTable() *DataTable {
	return &DataTable{Type: "data-grid", Rows: []*DataTableRow{}}
}

func (d *DataTable) AddTextColumn(text string) {
	d.Columns = append(d.Columns, &DataTableColumn{Text: text})
}

func (d *DataTable) AddRow(args ...interface{}) {
	cells := make([]*DataTableCell, len(args), len(args))

	for i := 0; i < len(args); i++ {
		cells[i] = &DataTableCell{Value: args[i]}
	}
	row := &DataTableRow{Cells: cells}
	d.Rows = append(d.Rows, row)

}

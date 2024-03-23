package printer

import "github.com/rodaine/table"

type TablePrinter struct {
	tbl table.Table
}

func NewTablePrinter(colsHeaders ...interface{}) TablePrinter {
	tbl := table.New(colsHeaders...)
	return TablePrinter{tbl}
}

func (t TablePrinter) AddRow(vals ...any) {
	t.tbl.AddRow(vals...)
}

func (t TablePrinter) Print() {
	t.tbl.Print()
}

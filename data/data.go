package data

type Table struct {
	Headers []Column
	Rows    []Row
}

type Row struct {
	Columns []Column
}

type Column struct {
	Value string
}

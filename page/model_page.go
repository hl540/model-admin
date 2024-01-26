package page

type ModelPage interface {
	Table() *Table
	Detail() *Detail
	Form() *Form
}

type Detail struct {
}

type Form struct {
}

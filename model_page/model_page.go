package model_page

import "github.com/hl540/model-admin/model_page/table"

type ModelTablePage interface {
	Table() *table.Table
}

type ModelDetailPage interface {
	Detail() *Detail
}

type ModelEditPage interface {
	Edit() *Form
}

type ModelNewPage interface {
	Add() *Form
}

type ModelPage interface {
	ModelTablePage
	ModelDetailPage
	ModelEditPage
	ModelNewPage
}

type Detail struct {
}

type Form struct {
}

package model_page

import (
	"github.com/hl540/model-admin/model_page/detail"
	"github.com/hl540/model-admin/model_page/form"
	"github.com/hl540/model-admin/model_page/table"
)

// ModelTablePage 模型表格页面
type ModelTablePage interface {
	Table() *table.Table
}

// ModelDetailPage 模型详情页面
type ModelDetailPage interface {
	Detail() *detail.Detail
}

// ModelEditPage 模型新增页面
type ModelEditPage interface {
	Edit() *form.Form
}

// ModelNewPage 模型编辑页面
type ModelNewPage interface {
	New() *form.Form
}

type ModelPage interface {
	ModelTablePage
	ModelDetailPage
	ModelEditPage
	ModelNewPage
}

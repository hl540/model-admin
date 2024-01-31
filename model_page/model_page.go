package model_page

import (
	"github.com/hl540/model-admin/model_page/detail_page"
	"github.com/hl540/model-admin/model_page/form_page"
	"github.com/hl540/model-admin/model_page/table_page"
)

// ModelTablePage 模型表格页面
type ModelTablePage interface {
	Table() *table_page.Table
}

// ModelDetailPage 模型详情页面
type ModelDetailPage interface {
	Detail() *detail_page.Field
}

// ModelEditPage 模型新增页面
type ModelEditPage interface {
	Edit() *form_page.Form
}

// ModelNewPage 模型编辑页面
type ModelNewPage interface {
	New() *form_page.Form
}

type ModelPage interface {
	ModelTablePage
	ModelDetailPage
	ModelEditPage
	ModelNewPage
}

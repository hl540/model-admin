package model_page

import (
	"fmt"
)

// modelPage集合
var modelPageSet = make(map[string]any)

// Register 注册一个modelPage
func Register(name string, modelPage any) {
	modelPageSet[name] = modelPage
}

// GetTable 获取ModelTablePage的实现
func GetTable(name string) (*Table, error) {
	modelPage, ok := modelPageSet[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	page, ok := modelPage.(ModelTablePage)
	if !ok {
		return nil, fmt.Errorf("unrealized ModelTablePage")
	}
	return page.Table(), nil
}

// GetDetail 获取ModelDetailPage的实现
func GetDetail(name string) (*Field, error) {
	modelPage, ok := modelPageSet[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	page, ok := modelPage.(ModelDetailPage)
	if !ok {
		return nil, fmt.Errorf("unrealized ModelDetailPage")
	}
	return page.Detail(), nil
}

// GetEdit 获取ModelEditPage的实现
func GetEdit(name string) (*Form, error) {
	modelPage, ok := modelPageSet[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	page, ok := modelPage.(ModelEditPage)
	if !ok {
		return nil, fmt.Errorf("unrealized ModelEditPage")
	}
	return page.Edit(), nil
}

// GetNew 获取ModelNewPage的实现
func GetNew(name string) (*Form, error) {
	modelPage, ok := modelPageSet[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	page, ok := modelPage.(ModelNewPage)
	if !ok {
		return nil, fmt.Errorf("unrealized ModelNewPage")
	}
	return page.New(), nil
}

// ModelTablePage 模型表格页面
type ModelTablePage interface {
	Table() *Table
}

// ModelDetailPage 模型详情页面
type ModelDetailPage interface {
	Detail() *Field
}

// ModelEditPage 模型新增页面
type ModelEditPage interface {
	Edit() *Form
}

// ModelNewPage 模型编辑页面
type ModelNewPage interface {
	New() *Form
}

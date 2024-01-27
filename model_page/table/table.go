package table

import (
	"gorm.io/gorm"
)

// Table 表格
type Table struct {
	tableName string
	columns   []*Column
	columnMap map[string]*Column

	getDataFn GetDataFn
}

// SetTableName 设置数据库表名称
func (t *Table) SetTableName(tableName string) *Table {
	t.tableName = tableName
	return t
}

// GetDataFn 取表格数据方法
type GetDataFn func(db *gorm.DB, param *GetDataParam) ([]map[string]any, int64, error)

// SetGetDataFn 设置获取表格数据的方法，用于自定义数据源
func (t *Table) SetGetDataFn(fn GetDataFn) {
	t.getDataFn = fn
}

// AddColumn 添加一列
func (t *Table) AddColumn(name, title string) *Column {
	column := &Column{
		Name:  name,
		Title: title,
	}
	t.columns = append(t.columns, column)
	if t.columnMap == nil {
		t.columnMap = make(map[string]*Column)
	}
	t.columnMap[name] = column
	return column
}

// GetColumn 获取单个列
func (t *Table) GetColumn(title string) *Column {
	return t.columnMap[title]
}

// GetColumns 获取所有个列
func (t *Table) GetColumns() []*Column {
	return t.columns
}

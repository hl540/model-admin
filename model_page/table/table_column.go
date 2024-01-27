package table

import "html/template"

// Column 表格列
type Column struct {
	Name      string          // 列字段名称
	Title     string          // 列展示名称
	displayFn ColumnDisplayFn // 列值展示方法
}

// ColumnDisplayFn 列值展示方法
type ColumnDisplayFn func(value any) template.HTML

// SetDisplay 设置列值展示方法
func (c *Column) SetDisplay(fn ColumnDisplayFn) {
	c.displayFn = fn
}

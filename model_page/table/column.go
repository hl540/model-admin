package table

import (
	"github.com/spf13/cast"
	"html/template"
)

// Column 表格列
type Column struct {
	Title       string            // 列展示名称
	name        string            // 列字段名称
	displayFns  []ColumnDisplayFn // 列值展示方法
	valueMap    map[string]any    // 列值映射
	displayHtml template.HTML     // display内容
	joinName    string            // join名称
}

// ColumnDisplayFn 列值展示方法
type ColumnDisplayFn func(value map[string]any) template.HTML

// SetDisplayFn 设置列值展示方法
func (c *Column) SetDisplayFn(fn ColumnDisplayFn) {
	c.displayFns = append(c.displayFns, fn)
}

// SetValueMap 设置列值映射
func (c *Column) SetValueMap(vm map[string]any) *Column {
	c.valueMap = vm
	return c
}

// ExecuteDisplay 执行display获取列最终显示
func (c *Column) ExecuteDisplay(rowValue map[string]any) template.HTML {
	// 获取当前字段的值
	currentColValue := cast.ToString(rowValue[c.name])
	// 执行值映射
	if nweValue, ok := c.valueMap[currentColValue]; ok {
		rowValue[c.name] = nweValue
	}
	// 初始化展示内容
	c.displayHtml = template.HTML(cast.ToString(rowValue[c.name]))
	// 执行display方法获取列的展示内容
	for _, displayFn := range c.displayFns {
		c.displayHtml = displayFn(rowValue)
	}
	return c.displayHtml
}

// Join join名称
func (c *Column) Join(name string) *Column {
	c.joinName = name
	return c
}
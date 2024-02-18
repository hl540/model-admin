package table_page

import (
	"html/template"

	"github.com/spf13/cast"
)

type ShowFormatName string

// ImageFormat 图片格式
const ImageFormat ShowFormatName = "imageFormatter"

// LinkFormat 链接格式
const LinkFormat ShowFormatName = "linkFormatter"

// Column 表格列
type Column struct {
	Title         string              // 列展示名称
	Name          string              // 列字段名称
	Hide          bool                // 是否隐藏
	SortAble      bool                // 可以排序
	ShowFormat    ShowFormatName      // 内容展示格式化名称
	valueFormatFn ColumnValueFormatFn // 值格式化方法
	valueMap      map[string]any      // 列值映射
	displayHtml   template.HTML       // display内容
	joinName      string              // join名称
}

// ColumnValueFormatFn 列值格式化方法
type ColumnValueFormatFn func(value map[string]any) any

// SetHide 设置列隐藏
func (c *Column) SetHide() *Column {
	c.Hide = true
	return c
}

// SetSort 设置可以排序
func (c *Column) SetSort() *Column {
	c.SortAble = true
	return c
}

// SetShowFormatName 设置列值展示格式化方法名称
func (c *Column) SetShowFormatName(name ShowFormatName) *Column {
	c.ShowFormat = name
	return c
}

// SetValueFormatFn 设置列值格式化方法
func (c *Column) SetValueFormatFn(fn ColumnValueFormatFn) *Column {
	c.valueFormatFn = fn
	return c
}

// SetValueMap 设置列值映射
func (c *Column) SetValueMap(vm map[string]any) *Column {
	c.valueMap = vm
	return c
}

// JoinName join名称
func (c *Column) JoinName(name string) *Column {
	c.joinName = name
	return c
}

// ColumnValue 列值
type ColumnValue struct {
	RowValue    any           // 原始值
	DisplayText template.HTML // 展示内容
}

// ParseValue 解析列值
func (c *Column) ParseValue(rowValue map[string]any) any {
	newValue := rowValue[c.Name]
	// 执行值映射
	if c.valueMap != nil {
		currentColValue := cast.ToString(rowValue[c.Name])
		if _, ok := c.valueMap[currentColValue]; ok {
			newValue = c.valueMap[currentColValue]
		}
	}
	// 执行格式化
	if c.valueFormatFn != nil {
		newValue = c.valueFormatFn(rowValue)
	}
	return newValue
}

/***********************************************************************/

// FormatDateTime 展示为日期
func (c *Column) FormatDateTime(layout string) *Column {
	c.SetValueFormatFn(func(value map[string]any) any {
		valueStr := cast.ToString(value[c.Name])
		time, err := cast.StringToDate(valueStr)
		if err != nil {
			return template.HTML(err.Error())
		}
		return time.Format(layout)
	})
	return c
}

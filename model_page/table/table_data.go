package table

import (
	"github.com/hl540/model-admin/data_source"
)

// TmplData 表格模板数据结构体
type TmplData struct {
	Columns []*Column // 列配置
	Data    [][]any   // 数据
	Count   int64     // 数据总数
}

// GetTmplData 获取表格模板数据
func (t *Table) GetTmplData(param *GetDataParam) (*TmplData, error) {
	// 加载数据
	data, count, err := t.GetData(param)
	if err != nil {
		return nil, err
	}
	// 组装模板参数
	return &TmplData{
		Columns: t.GetColumns(),
		Data:    t.parseData(data),
		Count:   count,
	}, nil
}

// GetData 获取表格数据
func (t *Table) GetData(param *GetDataParam) ([]map[string]any, int64, error) {
	if t.getDataFn != nil {
		return t.getDataFn(data_source.GetDB(), param)
	}
	return nil, 0, nil
}

// 解析表格数据
func (t *Table) parseData(data []map[string]any) [][]any {
	var result [][]any
	for _, item := range data {
		var value []any
		for _, col := range t.GetColumns() {
			if col.displayFn != nil {
				value = append(value, col.displayFn(item[col.Name]))
			} else {
				value = append(value, item[col.Name])
			}
		}
		result = append(result, value)
	}
	return result
}

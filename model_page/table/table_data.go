package table

import "github.com/hl540/model-admin/data_source"

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
		return t.getDataFn(param)
	}
	// 获取数据源
	if t.dataSource == "" {
		t.dataSource = "default"
	}
	db, err := data_source.GetDB(t.dataSource)
	if err != nil {
		return nil, 0, err
	}
	// 开始查询数据
	query := db.Table(t.tableName)
	// 附件table过滤条件
	if t.dataFilterFn != nil {
		query = t.dataFilterFn(query)
	}
	// 设置请求过滤条件
	for key, value := range param.Filter {
		query.Where(key+" = ?", value)
	}
	// 查询总数
	var count int64
	if err = query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 排序
	if param.Sort != nil {
		query.Order(param.Sort[0] + " " + param.Sort[1])
	}
	// 分页
	query.Offset((param.Pagination.Page - 1) * param.Pagination.Size)
	query.Limit(param.Pagination.Size)
	// 解析数据
	var data []map[string]any
	if err = query.Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
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

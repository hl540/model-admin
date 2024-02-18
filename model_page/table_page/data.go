package table_page

import (
	"github.com/hl540/model-admin/data_source"
)

// TableData 模型表格查询结果
type TableData struct {
	Rows  []map[string]any // 数据行
	Count int64            // 总数
	Size  int              // 数量
	Page  int              // 页码
}

// GetData 获取表格模板数据
func (t *Table) GetData(param *QueryParam) (*TableData, error) {
	var result = &TableData{
		Page: param.Page,
		Size: param.Size,
		Rows: make([]map[string]any, 0),
	}
	// 请求模板的时候是否加载数据
	if t.loadTemplateData == false {
		return result, nil
	}

	var err error
	// 加载数据
	if t.customGetDataFn != nil { // 自定义数据源
		result.Rows, result.Count, err = t.customGetDataFn(param)
	} else { // db数据
		result.Rows, result.Count, err = t.GetDBData(param)
	}
	if err != nil {
		return nil, err
	}
	// 根据列配置处理最终值
	for index, rows := range result.Rows {
		for key := range rows {
			column := t.columnMap[key]
			if column == nil {
				continue
			}
			result.Rows[index][key] = column.ParseValue(rows)
		}
	}
	return result, nil
}

// GetDBData 从数据库获取数据
func (t *Table) GetDBData(param *QueryParam) ([]map[string]any, int64, error) {
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
	// 连表
	for _, joinExpression := range t.joins {
		query.Joins(joinExpression)
	}
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
	countQuery := query
	if err = countQuery.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// 排序
	if param.Sort != nil {
		query.Order(param.Sort[0] + " " + param.Sort[1])
	}
	// 分页
	query.Offset((param.Page - 1) * param.Size)
	query.Limit(param.Size)
	// 查询列
	query.Select(t.parseQueryColumnName())
	// 解析数据
	var data []map[string]any
	if err = query.Find(&data).Error; err != nil {
		return nil, 0, err
	}
	return data, count, nil
}

// 查询查询的列名
func (t *Table) parseQueryColumnName() []string {
	var columnNames []string
	for _, col := range t.Columns {
		if col.joinName == "" {
			columnNames = append(columnNames, t.tableName+"."+col.Name+" AS "+col.Name)
		} else {
			columnNames = append(columnNames, col.joinName+" AS "+col.Name)
		}
	}
	return columnNames
}

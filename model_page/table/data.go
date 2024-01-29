package table

import (
	"html/template"

	"github.com/hl540/model-admin/data_source"
)

// GetTmplData 获取表格模板数据
func (t *Table) GetTmplData(param *GetDataParam) ([][]template.HTML, int64, error) {
	// 加载数据
	data, count, err := t.GetData(param)
	if err != nil {
		return nil, 0, err
	}
	return t.parseData(data), count, nil
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
	query := db.Table(t.TableName)
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
	query.Offset((param.Pagination.Page - 1) * param.Pagination.Size)
	query.Limit(param.Pagination.Size)
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
	for _, col := range t.columns {
		if col.joinName == "" {
			columnNames = append(columnNames, t.TableName+"."+col.Name+" AS "+col.Name)
		} else {
			columnNames = append(columnNames, col.joinName+" AS "+col.Name)
		}
	}
	return columnNames
}

// 解析表格数据
func (t *Table) parseData(data []map[string]any) [][]template.HTML {
	var result [][]template.HTML
	for _, item := range data {
		var value []template.HTML
		for _, col := range t.columns {
			value = append(value, col.ExecuteDisplay(item))
		}
		result = append(result, value)
	}
	return result
}

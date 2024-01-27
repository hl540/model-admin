package table

import "gorm.io/gorm"

// Table 表格
type Table struct {
	title        string             // 表格title
	tableName    string             // 数据表名称
	dataSource   string             // 数据源名称
	columns      []*Column          // 所有列
	columnMap    map[string]*Column // 所有列
	getDataFn    GetDataFn          // 自定义数据方法
	dataFilterFn DataFilterFn       // 数据过滤方法
}

// SetTitle 设置表格title
func (t *Table) SetTitle(title string) *Table {
	t.title = title
	return t
}

// SetTableName 设置数据库表名称
func (t *Table) SetTableName(tableName string) *Table {
	t.tableName = tableName
	return t
}

// GetDataFn 格数据方法
type GetDataFn func(param *GetDataParam) ([]map[string]any, int64, error)

// SetGetDataFn 设置自定义表格数据方法
func (t *Table) SetGetDataFn(fn GetDataFn) *Table {
	t.getDataFn = fn
	return t
}

// SetDataSource 设置数据源
func (t *Table) SetDataSource(name string) *Table {
	t.dataSource = name
	return t
}

// DataFilterFn 过滤方法
type DataFilterFn func(query *gorm.DB) *gorm.DB

// SetDataFilterFn 设置数据过滤方法
func (t *Table) SetDataFilterFn(fn DataFilterFn) *Table {
	t.dataFilterFn = fn
	return t
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

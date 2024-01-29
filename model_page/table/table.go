package table

import "gorm.io/gorm"

// Table 表格
type Table struct {
	Title        string             // 表格title
	TableName    string             // 数据表名称
	dataSource   string             // 数据源名称
	joins        []string           // 连表条件，可以有多个
	columns      []*Column          // 所有列
	columnMap    map[string]*Column // 所有列
	getDataFn    GetDataFn          // 自定义数据方法
	dataFilterFn DataFilterFn       // 数据过滤方法
}

// SetTitle 设置表格title
func (t *Table) SetTitle(title string) *Table {
	t.Title = title
	return t
}

// SetTableName 设置数据库表名称
func (t *Table) SetTableName(tableName string) *Table {
	t.TableName = tableName
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

// GetColumns 获取所有列
func (t *Table) GetColumns() []*Column {
	return t.columns
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

// Join 连表
func (t *Table) Join(expression string) *Table {
	t.joins = append(t.joins, expression)
	return t
}

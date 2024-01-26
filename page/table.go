package page

// Table 表格
type Table struct {
	// 所有列
	columns []*TableColumn
	actions []*Action
}

// AddColumn 添加一列
func (t *Table) AddColumn(title string, display string) *TableColumn {
	column := &TableColumn{}
	t.columns = append(t.columns, column)
	return column
}

// AddAction 添加一个操作，位于表格每一行的末尾
func (t *Table) AddAction(title string) *Action {
	action := &Action{}
	t.actions = append(t.actions, action)
	return action
}

// TableColumn 表格列
type TableColumn struct {
	title string
}

type Action struct {
	title string
}

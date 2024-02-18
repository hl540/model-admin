package model_page

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/data_source"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Table 表格
type Table struct {
	// 表格基础信息配置
	Title            string             // 表格title
	Columns          []*Column          // 所有列
	columnMap        map[string]*Column // 所有列
	PrimaryColumn    *Column            // 唯一列
	FixedLeftNumber  int                // 左侧固定列数
	FixedRightNumber int                // 右侧固定列数
	loadTemplateData bool               // 请求模板的时候是否加载数据

	// 数据源配置
	tableName       string       // 数据表名称
	dataSource      string       // 数据源名称
	joins           []string     // 连表条件，可以有多个
	customGetDataFn GetDataFn    // 自定义数据获取方法
	dataFilterFn    DataFilterFn // 数据过滤方法
}

func NewTable(title, tableName string) *Table {
	return &Table{
		Title:            title,
		FixedLeftNumber:  2,
		FixedRightNumber: 1,
		loadTemplateData: true,
		tableName:        tableName,
		dataSource:       "default",
		PrimaryColumn:    &Column{},
	}
}

// AddColumn 添加一列
func (t *Table) AddColumn(name, title string) *Column {
	column := &Column{
		Name:  name,
		Title: title,
	}
	t.Columns = append(t.Columns, column)
	if t.columnMap == nil {
		t.columnMap = make(map[string]*Column)
	}
	t.columnMap[name] = column
	return column
}

// AddPrimaryColumn 设置唯一列
func (t *Table) AddPrimaryColumn(name, title string) *Column {
	col := t.AddColumn(name, title)
	t.PrimaryColumn = col
	return col
}

// GetColumns 获取所有列
func (t *Table) GetColumns() []*Column {
	return t.Columns
}

// FixedNumber 设置固定列
func (t *Table) FixedNumber(number ...int) *Table {
	if len(number) >= 1 {
		t.FixedLeftNumber = number[0]
	}
	if len(number) >= 2 {
		t.FixedRightNumber = number[1]
	}
	return t
}

// LoadData 是否加载模板数据
func (t *Table) LoadData(b bool) *Table {
	t.loadTemplateData = b
	return t
}

// TableName 设置数据库表名称
func (t *Table) TableName(tableName string) *Table {
	t.tableName = tableName
	return t
}

// DataSource 设置数据源
func (t *Table) DataSource(name string) *Table {
	t.dataSource = name
	return t
}

// Join 连表
func (t *Table) Join(expression string) *Table {
	t.joins = append(t.joins, expression)
	return t
}

// GetDataFn 格数据方法
type GetDataFn func(param *QueryParam) ([]map[string]any, int64, error)

// CustomDataFn 设置自定义表格数据方法
func (t *Table) CustomDataFn(fn GetDataFn) *Table {
	t.customGetDataFn = fn
	return t
}

// DataFilterFn 过滤方法
type DataFilterFn func(query *gorm.DB) *gorm.DB

// DataFilterFn 设置数据过滤方法
func (t *Table) DataFilterFn(fn DataFilterFn) *Table {
	t.dataFilterFn = fn
	return t
}

/**********************************************************************************************************************/

const DefaultSize = 10 // 单页默认数量
const SizeMax = 100    // 单页最大数量
const DefaultPage = 1  // 默认页码

// 解析分页参数
func parsePaginationParam(ctx *gin.Context) (int, int) {
	page := cast.ToInt(ctx.Query("_page"))
	size := cast.ToInt(ctx.Query("_size"))
	if page <= 0 {
		page = DefaultPage
	}
	if size <= 0 {
		size = DefaultSize
	}
	if size > SizeMax {
		size = SizeMax
	}
	return page, size
}

// 解析排序参数
func parseSortParam(ctx *gin.Context) []string {
	sortName := ctx.Query("_sort_name")
	if len(sortName) == 0 {
		return nil
	}
	sortOrder := ctx.Query("_sort_order")
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "DESC"
	}
	return []string{sortName, sortOrder}
}

// 解析筛选参数
func parseFilterParam(ctx *gin.Context) map[string]any {
	filterParam := make(map[string]any)
	for k := range ctx.Request.URL.Query() {
		if len(k) > 8 && k[:8] == "_filter_" {
			filterParam[k[8:]] = ctx.Query(k)
		}
	}
	return filterParam
}

// QueryParam 表格数据查询参数
type QueryParam struct {
	Filter map[string]any // 筛选参数
	Sort   []string       // 排序参数
	Page   int            // 数量
	Size   int            // 页码
}

// ParseQueryParam 解析查询参数
func ParseQueryParam(ctx *gin.Context) *QueryParam {
	param := &QueryParam{
		Filter: parseFilterParam(ctx), // 解析筛选参数
		Sort:   parseSortParam(ctx),   // 解析排序参数
	}
	// 解析分页参数
	param.Page, param.Size = parsePaginationParam(ctx)
	return param
}

/**********************************************************************************************************************/

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

/**********************************************************************************************************************/

type ShowFormatName string

const ImageFormat ShowFormatName = "imageFormatter" // 图片格式
const LinkFormat ShowFormatName = "linkFormatter"   // 链接格式

// Column 表格列
type Column struct {
	Title          string              // 列展示名称
	Name           string              // 列字段名称
	IsHide         bool                // 是否隐藏
	IsSort         bool                // 可以排序
	FilterOption   *FilterOption       // 可筛选
	ShowFormatName ShowFormatName      // 内容展示格式化名称
	valueFormatFn  ColumnValueFormatFn // 值格式化方法
	valueMap       map[string]any      // 列值映射
	joinName       string              // join名称
}

// ColumnValueFormatFn 列值格式化方法
type ColumnValueFormatFn func(value map[string]any) any

// Hide 设置列隐藏
func (c *Column) Hide() *Column {
	c.IsHide = true
	return c
}

// Sort 设置可以排序
func (c *Column) Sort() *Column {
	c.IsSort = true
	return c
}

// FilterType 过滤类型
type FilterType string

const FilterTypeLike FilterType = "FilterTypeLike"                   // 模糊搜索
const FilterTypeSingle FilterType = "FilterTypeSingle"               // 单选
const FilterTypeMultiple FilterType = "FilterTypeMultiple"           // 多选
const FilterTypeDatetimeRange FilterType = "FilterTypeDatetimeRange" // 时间范围

// FilterOption 过滤选项
type FilterOption struct {
	FilterType FilterType
	Options    []map[string]string
}

// Filter 设置字段过滤
func (c *Column) Filter(option *FilterOption) *Column {
	c.FilterOption = option
	return c
}

// ShowFormat 设置列值展示格式化方法名称
func (c *Column) ShowFormat(name ShowFormatName) *Column {
	c.ShowFormatName = name
	return c
}

// FormatFn 设置列值格式化方法
func (c *Column) FormatFn(fn ColumnValueFormatFn) *Column {
	c.valueFormatFn = fn
	return c
}

// ValueMap 设置列值映射
func (c *Column) ValueMap(vm map[string]any) *Column {
	c.valueMap = vm
	return c
}

// JoinName join名称
func (c *Column) JoinName(name string) *Column {
	c.joinName = name
	return c
}

// FormatDateTime 展示为日期
func (c *Column) FormatDateTime(layout string) *Column {
	c.FormatFn(func(value map[string]any) any {
		valueStr := cast.ToString(value[c.Name])
		time, err := cast.StringToDate(valueStr)
		if err != nil {
			return template.HTML(err.Error())
		}
		return time.Format(layout)
	})
	return c
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

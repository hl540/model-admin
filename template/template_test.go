package template

import (
	"fmt"
	table2 "github.com/hl540/model-admin/model_page/table"
	"gorm.io/gorm"
	"html/template"
	"testing"
)

var tableData = []map[string]any{
	{
		"name": "张三",
		"age":  25,
		"sex":  "男",
	},
	{
		"name": "王五",
		"age":  30,
		"sex":  "女",
	},
	{
		"name": "小明",
		"age":  12,
		"sex":  "男",
	},
}

func TestTableTemplate(t *testing.T) {
	table := &table2.Table{}
	table.AddColumn("name", "名称").SetDisplayFn(func(value any) template.HTML {
		return template.HTML(fmt.Sprintf("<h1>%v<h1>", value))
	})
	table.AddColumn("age", "年龄")
	table.AddColumn("sex", "性别")
	table.SetGetDataFn(func(db *gorm.DB, param *table2.GetDataParam) ([]map[string]any, int64, error) {
		return tableData, int64(len(tableData)), nil
	})
	SetTemplatePath("./tmpl")
	html, err := TableTemplate(table, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(html)
}

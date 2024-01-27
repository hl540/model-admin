package main

import (
	"fmt"
	table2 "github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func init() {
	template2.SetTemplatePath("./tmpl")
}

func main() {
	http.HandleFunc("/table", func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := getTable(request)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		writer.Write([]byte(tmpl))
	})
	http.ListenAndServe(":9696", nil)
}

func getTable(req *http.Request) (template.HTML, error) {
	table := &table2.Table{}
	table.AddColumn("id", "ID")
	table.AddColumn("name", "名称").SetDisplay(func(value any) template.HTML {
		return template.HTML(fmt.Sprintf("<h1>%v<h1>", value))
	})
	table.AddColumn("age", "年龄")
	table.AddColumn("sex", "性别")
	table.SetGetDataFn(func(db *gorm.DB, param *table2.GetDataParam) ([]map[string]any, int64, error) {
		return tableData, int64(len(tableData)), nil
	})
	return template2.TableTemplate(table, table2.ParseGetDataParam(req))
}

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

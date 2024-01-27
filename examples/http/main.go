package main

import (
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"github.com/hl540/model-admin/model_page"
	table2 "github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"log"
	"net/http"
)

func init() {
	template2.SetTemplatePath("./tmpl")
}

func main() {
	conf, err := config.LoadFromYaml("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = data_source.InitDB(conf.Databases); err != nil {
		log.Fatal(err)
	}
	template2.SetTemplatePath("../../tmpl")
	var userModel any = &UserModel{}
	if tablePage, ok := userModel.(model_page.ModelTablePage); ok {
		http.HandleFunc("/user/table", func(writer http.ResponseWriter, request *http.Request) {
			tmpl, err := template2.TableTemplate(tablePage.Table(), table2.ParseGetDataParam(request))
			if err != nil {
				writer.Write([]byte(err.Error()))
				return
			}
			writer.Write([]byte(tmpl))
		})
	}

	http.ListenAndServe(":9696", nil)
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

type UserModel struct{}

func (u *UserModel) Table() *table2.Table {
	table := &table2.Table{}
	table.AddColumn("id", "ID")
	table.AddColumn("name", "名称")
	table.AddColumn("age", "年龄")
	table.AddColumn("sex", "性别")
	table.SetTableName("user").SetTitle("用户列表")
	return table
}

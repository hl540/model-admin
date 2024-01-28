package main

import (
	modeladmin "github.com/hl540/model-admin"
	table2 "github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"log"
	"net/http"
)

func init() {
	template2.SetTemplatePath("./tmpl")
}

func main() {
	handler := modeladmin.New()
	if err := handler.Init("./conf.yaml"); err != nil {
		log.Fatal(err)
	}
	handler.AddModelPage("user", &UserModel{})
	http.ListenAndServe(":9696", handler)
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
	table.Join("LEFT JOIN role ON user.role_id = role.id")
	table.AddColumn("id", "ID").DisplayLink("https://www.baidu.com?wd={id}", "_blank")
	table.AddColumn("name", "名称")
	table.AddColumn("age", "年龄")
	table.AddColumn("sex", "性别").SetValueMap(map[string]any{
		"1": "男",
		"0": "女",
	})
	table.AddColumn("shot", "头像").DisplayImage(30, 30)
	table.AddColumn("role_name", "所属组").Join("role.name")
	table.AddColumn("created_at", "创建时间").DisplayDateTime("2006-01-02 15:04:05")
	table.SetTableName("user").SetTitle("用户列表")
	return table
}

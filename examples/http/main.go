package main

import (
	_ "github.com/hl540/model-admin/src/bootstrap_admin_ui"
	"log"
	"net/http"

	"github.com/hl540/model-admin/config"
	modeladmin "github.com/hl540/model-admin/handler"
	table2 "github.com/hl540/model-admin/model_page/table"
)

func main() {
	//if err := bootstrap_admin_ui.SetTemplatePath("../../tmpl"); err != nil {
	//	log.Fatal(err)
	//}
	conf, err := config.LoadFromYaml("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	handler := modeladmin.New()
	err = handler.SetConfig(conf).SetModelPage(models).Init()
	if err != nil {
		log.Fatal(err)
	}
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

var models = map[string]any{
	"user": new(UserModel),
	"role": new(RoleModel),
}

type UserModel struct{}

func (u *UserModel) Table() *table2.Table {
	table := new(table2.Table)
	table.Join("LEFT JOIN role ON user.role_id = role.id")
	table.AddColumn("id", "ID").SetPrimary().DisplayLink("https://www.baidu.com?wd={id}", "_blank")
	table.AddColumn("name", "名称")
	table.AddColumn("age", "年龄").SetHide()
	table.AddColumn("sex", "性别").SetValueMap(map[string]any{
		"1": "男",
		"0": "女",
	})
	table.AddColumn("shot", "头像").DisplayImage(30, 30)
	table.AddColumn("role_name", "所属组").JoinName("role.name")
	table.AddColumn("created_at", "创建时间").DisplayDateTime("2006-01-02 15:04:05")
	table.SetTableName("user").SetTitle("用户列表").SetFixedNumber(2, 1)
	return table
}

type RoleModel struct {
}

func (r *RoleModel) Table() *table2.Table {
	table := new(table2.Table)
	table.AddColumn("id", "ID").SetPrimary()
	table.AddColumn("name", "名称")
	table.SetTableName("role").SetTitle("角色列表")
	return table
}

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	modeladmin "github.com/hl540/model-admin"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page"
	table2 "github.com/hl540/model-admin/model_page/table_page"
	_ "github.com/hl540/model-admin/src/bootstrap_admin_ui"
)

func init() {
	// gin.SetMode(gin.ReleaseMode)
	model_page.Register("user", new(UserModel))
	model_page.Register("role", new(RoleModel))
}

func main() {
	conf, err := config.LoadFromYaml("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()

	admin := modeladmin.NewAdmin(app)
	if err = admin.SetConfig(conf).Init(); err != nil {
		log.Fatal(err)
	}

	app.Run(conf.ServerListen)
}

type UserModel struct{}

func (u *UserModel) Table() *table2.Table {
	table := table2.NewTable("user")
	table.SetTitle("用户列表")
	table.AddColumn("id", "ID").SetPrimary().SetFormat(table2.LinkFormat)
	table.AddColumn("name", "名称")
	table.AddColumn("age", "年龄").SetHide()
	table.AddColumn("sex", "性别").SetValueMap(map[string]any{
		"1": "男",
		"0": "女",
	})
	table.AddColumn("shot", "头像").SetFormat(table2.ImageFormat)
	table.Join("LEFT JOIN role ON user.role_id = role.id")
	table.AddColumn("role_name", "所属组").JoinName("role.name")
	table.AddColumn("created_at", "创建时间").DisplayDateTime("2006-01-02 15:04:05")
	return table
}

type RoleModel struct {
}

func (r *RoleModel) Table() *table2.Table {
	table := table2.NewTable("role")
	table.SetTitle("角色列表")
	table.AddColumn("id", "ID").SetPrimary()
	table.AddColumn("name", "名称")
	return table
}

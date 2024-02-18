package main

import (
	"log"

	"github.com/gin-gonic/gin"
	modeladmin "github.com/hl540/model-admin"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page"
	_ "github.com/hl540/model-admin/src/bootstrap_admin_ui"
)

func init() {
	// gin.SetMode(gin.ReleaseMode)
	model_page.Register("user", new(UserModel))
	model_page.Register("role", new(RoleModel))
	model_page.Register("test", new(TestModel))
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

func (u *UserModel) Table() *model_page.Table {
	table := model_page.NewTable("用户列表", "user")
	table.AddPrimaryColumn("id", "ID").ShowFormat(model_page.LinkFormat).Sort()
	table.AddColumn("name", "名称")
	table.AddColumn("age", "年龄").Hide()
	table.AddColumn("sex", "性别").ValueMap(map[string]any{
		"1": "男",
		"0": "女",
	}).Sort()
	table.AddColumn("shot", "头像").ShowFormat(model_page.ImageFormat)
	table.Join("LEFT JOIN role ON user.role_id = role.id")
	table.AddColumn("role_name", "所属组").JoinName("role.name").Sort()
	table.AddColumn("created_at", "创建时间").FormatDateTime("2006-01-02 15:04:05").Sort()
	return table
}

type RoleModel struct {
}

func (r *RoleModel) Table() *model_page.Table {
	table := model_page.NewTable("角色列表", "role")
	table.AddPrimaryColumn("id", "ID").Sort()
	table.AddColumn("name", "名称")
	return table
}

type TestModel struct{}

func (t *TestModel) Table() *model_page.Table {
	table := model_page.NewTable("sqlite测试", "test").DataSource("sqlite")
	table.AddPrimaryColumn("code", "ID").Sort()
	table.AddColumn("name", "名称").Sort()
	return table
}

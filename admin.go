package model_admin

import (
	_ "embed"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"github.com/hl540/model-admin/handler"
	template2 "github.com/hl540/model-admin/template"
	"github.com/hl540/model-admin/utils"
)

type Admin struct {
	app        *gin.Engine
	pathPrefix string         // 路由前缀
	conf       *config.Config // 配置
}

func NewAdmin(app *gin.Engine) *Admin {
	return &Admin{
		app:        app,
		pathPrefix: "/admin",
	}
}

// SetConfig 设置配置
func (a *Admin) SetConfig(conf *config.Config) *Admin {
	a.conf = conf
	return a
}

// Init 初始化
func (a *Admin) Init() error {
	// 初始化db
	if err := data_source.InitDB(a.conf.Databases); err != nil {
		return err
	}
	// 初始化模板渲染器
	if err := a.initHTMLTemplate(); err != nil {
		return err
	}
	// 路由前缀
	if a.conf.RouterPrefix != "" {
		a.pathPrefix = "/" + a.conf.RouterPrefix
	}
	// 初始化路由由组
	a.initRouter()
	a.printStartLog()
	return nil
}

// 初始化路由
func (a *Admin) initRouter() {
	adminRouter := a.app.Group(a.pathPrefix)
	// 首页路由
	adminRouter.GET("/", handler.AdminPageHandler())
	// 模型列表页面路由
	adminRouter.GET("/model/:mode_name/table", handler.TablePageHandler())
	// 模型详情页面路由
	adminRouter.GET("/model/:mode_name/detail", handler.DetailPageHandler())
	// 模型新增路由
	adminRouter.GET("/model/:mode_name/new", handler.NewPageHandler())
	adminRouter.POST("/model/:mode_name/new", handler.NewPageHandlerPost())
	// 模型编辑路由
	adminRouter.GET("/model/:mode_name/edit", handler.EditPageHandler())
	adminRouter.POST("/model/:mode_name/edit", handler.EditPageHandlerPost())
	// 自定义页面路由
	adminRouter.GET("/page/:page_name", handler.CustomPageHandler())
}

func (a *Admin) initHTMLTemplate() error {
	// 如果配置了模板路径
	if a.conf.TemplatePath != "" {
		files, err := utils.GetFilesName(a.conf.TemplatePath)
		if err != nil {
			return err
		}
		a.app.LoadHTMLFiles(files...)
		return nil
	}

	// 获取配置模板名称
	tmpl, err := template2.GetTemplate(a.conf.TemplateName)
	if err != nil {
		return err
	}

	// debug模式
	if gin.IsDebugging() {
		a.app.LoadHTMLFiles(tmpl.GetTemplateFiles()...)
		return nil
	}
	a.app.SetHTMLTemplate(tmpl.GetTemplate())
	return nil
}

//go:embed ascii_logo.txt
var asciiLogo string

func (a *Admin) printStartLog() {
	// a.app
	log.Println(asciiLogo)
}

package model_admin

import (
	_ "embed"
	"github.com/hl540/model-admin/handler"
	"github.com/hl540/model-admin/model_page"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
)

// Handler 模型页面处理器
type Handler struct {
	mux         *mux.Router
	adminRouter *mux.Router
	pathPrefix  string         // 路由前缀
	conf        *config.Config // 配置
	modelPages  map[string]any // 已注册模型页面
}

func New() *Handler {
	return &Handler{
		pathPrefix: "/admin",
		mux:        mux.NewRouter(),
	}
}

// SetConfig 设置配置
func (h *Handler) SetConfig(conf *config.Config) *Handler {
	h.conf = conf
	return h
}

// Init 初始化
func (h *Handler) Init() error {
	// 初始化db
	if err := data_source.InitDB(h.conf.Databases); err != nil {
		return err
	}

	// 初始化应用配置
	if h.conf.RouterPrefix != "" {
		h.pathPrefix = "/" + h.conf.RouterPrefix
	}
	// admin路由组
	h.adminRouter = h.mux.PathPrefix(h.pathPrefix).Subrouter()

	// 初始化默认路由
	handler.AdminPageHandler(h.adminRouter)

	// 初始化模型路由
	h.initModelPageRouter()

	// 输出路由注册log
	h.printRouterLogs()
	return nil
}

// 实现ServeHTTP
func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.mux.ServeHTTP(writer, request)
}

// SetModelPage 设置模型页面
func (h *Handler) SetModelPage(models map[string]any) *Handler {
	h.modelPages = models
	return h
}

// 初始化模型路由
func (h *Handler) initModelPageRouter() {
	for name, model := range h.modelPages {
		h.AddModelPage(name, model)
	}
}

// AddModelPage 注册单个模型页面
func (h *Handler) AddModelPage(name string, modelPage any) {
	// 注册模型表格页面处理器
	if tablePage, ok := modelPage.(model_page.ModelTablePage); ok {
		handler.TablePageHandler(h.adminRouter, name, tablePage.Table())
	}
	// 注册模型详情页面处理器
	if detailPage, ok := modelPage.(model_page.ModelDetailPage); ok {
		handler.DetailPageHandler(h.adminRouter, name, detailPage.Detail())
	}
	// 注册新建和编辑页面处理器，新建和编辑可以互相复用
	// 新建和编辑可以复用，如果其中一个有实现另一个没有就复用实现的那个
	newPage, npOk := modelPage.(model_page.ModelNewPage)
	editPage, epOk := modelPage.(model_page.ModelEditPage)
	// 如果新增页面有实现而编辑页面没有实现，则编辑复用新增的
	if npOk {
		handler.NewPageHandler(h.adminRouter, name, newPage.New())
		if !epOk {
			handler.EditPageHandler(h.adminRouter, name, newPage.New())
		}
	}
	// 如果编辑页面有实现而新增页面没有实现，则编辑复用编辑的
	if epOk {
		handler.EditPageHandler(h.adminRouter, name, editPage.Edit())
		if !npOk {
			handler.NewPageHandler(h.adminRouter, name, editPage.Edit())
		}
	}
	// 添加到已注册模型页面
	h.modelPages[name] = modelPage
}

//go:embed ascii_logo.txt
var asciiLogo string

// 打印l路由
func (h *Handler) printRouterLogs() {
	log.Println(asciiLogo)
	log.Printf("服务在%s上监听", h.conf.ServerListen)
	// 遍历所有路由并打印信息
	h.mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("%-30s %-5s\n", path, strings.Join(methods, ","))
		return nil
	})
}

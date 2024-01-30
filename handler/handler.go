package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"github.com/hl540/model-admin/model_page"
)

// Handler 模型页面处理器
type Handler struct {
	router         *mux.Router
	routerPrefix   string         // 路由前缀
	pageRouterLogs []string       // 已经注册的模型路由
	conf           *config.Config // 配置
	modelPages     map[string]any // 已注册模型页面
}

func New() *Handler {
	return &Handler{
		routerPrefix: "admin",
		router:       mux.NewRouter(),
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
		h.routerPrefix = h.conf.RouterPrefix
	}

	// 初始化默认路由
	h.initDefaultRouter()

	// 初始化模型路由
	h.initModelPageRouter()

	// 添加资源路由
	for path, handler := range extraHandler {
		h.router.PathPrefix(path).Handler(handler).Methods("GET")
	}

	// 输出路由注册log
	h.printPageRouterLogs()
	return nil
}

// 实现ServeHTTP
func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}

// AddHandleFunc 添加路由处理器
func (h *Handler) AddHandleFunc(name, path string, handler http.HandlerFunc, methods ...string) {
	h.router.HandleFunc(path, handler).Methods(methods...)
	h.addPageRouterLogs(name, path, methods...)
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
		h.addModelTablePageHandler(name, tablePage.Table())
	}
	// 注册模型详情页面处理器
	if detailPage, ok := modelPage.(model_page.ModelDetailPage); ok {
		h.addModelDetailPageHandler(name, detailPage.Detail())
	}
	// 注册新建和编辑页面处理器，新建和编辑可以互相复用
	// 新建和编辑可以复用，如果其中一个有实现另一个没有就复用实现的那个
	newPage, npOk := modelPage.(model_page.ModelNewPage)
	editPage, epOk := modelPage.(model_page.ModelEditPage)
	// 如果新增页面有实现而编辑页面没有实现，则编辑复用新增的
	if npOk {
		h.addModelNewPageHandler(name, newPage.New())
		if !epOk {
			h.addModelEditPageHandler(name, newPage.New())
		}
	}
	// 如果编辑页面有实现而新增页面没有实现，则编辑复用编辑的
	if epOk {
		h.addModelNewPageHandler(name, editPage.Edit())
		if !npOk {
			h.addModelEditPageHandler(name, editPage.Edit())
		}
	}
	// 添加到已注册模型页面
	h.modelPages[name] = modelPage
}

// 添加路由注册日志
func (h *Handler) addPageRouterLogs(name string, path string, methods ...string) {
	for _, method := range methods {
		log := fmt.Sprintf("%s\t\t%s\t%s", name, method, path)
		h.pageRouterLogs = append(h.pageRouterLogs, log)
	}
}

// 打印日子
func (h *Handler) printPageRouterLogs() {
	for _, logStr := range h.pageRouterLogs {
		log.Println(logStr)
	}
}

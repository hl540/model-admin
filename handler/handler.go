package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"github.com/hl540/model-admin/model_page"
	"github.com/hl540/model-admin/template"
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

func (h *Handler) Init() error {
	// 初始化db
	if err := data_source.InitDB(h.conf.Databases); err != nil {
		return err
	}

	// 初始化应用配置
	h.routerPrefix = h.conf.RouterPrefix

	// 添加资源路由
	for path, handler := range extraHandler {
		h.router.PathPrefix(path).Handler(handler).Methods("GET")
	}

	// 初始化默认路由
	h.initDefaultRouter()

	// 初始化模型路由
	h.initModelPageRouter()

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

// AddModelPage 添加单个模型页面
func (h *Handler) AddModelPage(name string, modelPage any) {
	// 添加到已注册模型页面
	h.modelPages[name] = modelPage
	// 注册路由
	if tablePage, ok := modelPage.(model_page.ModelTablePage); ok {
		path := fmt.Sprintf("/%s/%s/list", h.routerPrefix, name)
		h.AddHandleFunc(name, path, func(writer http.ResponseWriter, request *http.Request) {
			tmpl := template.TablePageRender(tablePage.Table(), request)
			writer.Write([]byte(tmpl))
		}, "GET")
	}
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

package model_admin

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/data_source"
	"github.com/hl540/model-admin/model_page"
	"github.com/hl540/model-admin/template"
	"log"
	"net/http"
)

// Handler 模型页面处理器
type Handler struct {
	router       *mux.Router
	routerPrefix string // 路由前缀
}

func New() *Handler {
	return &Handler{
		routerPrefix: "admin",
		router:       mux.NewRouter(),
	}
}

func (h *Handler) Init(confFile string) error {
	// 加载配置
	conf, err := config.LoadFromYaml(confFile)
	if err != nil {
		return err
	}
	// 初始化db
	if err = data_source.InitDB(conf.Databases); err != nil {
		return err
	}
	// 设置模板路径
	template.SetTemplatePath(conf.TemplatePath)

	// 初始化应用配置
	h.routerPrefix = conf.RouterPrefix
	return nil
}

// 实现ServeHTTP
func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.router.ServeHTTP(writer, request)
}

// AddModelPage 添加模型页面
func (h *Handler) AddModelPage(name string, modelPage any) {
	// 注册路由
	if tablePage, ok := modelPage.(model_page.ModelTablePage); ok {
		path := fmt.Sprintf("/%s/%s/list", h.routerPrefix, name)
		h.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
			tmpl := template.TableTemplate(tablePage.Table(), request)
			writer.Write([]byte(tmpl))
		}).Methods("GET")
		log.Printf("[%s] %s", name, path)
	}
}

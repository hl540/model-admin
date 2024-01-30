package handler

import (
	"net/http"

	"github.com/hl540/model-admin/template"
)

// 初始化默认路由
func (h *Handler) initDefaultRouter() {
	// 添加默认path路由
	h.AddHandleFunc("admin", "/"+h.routerPrefix, func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.LayoutPageRender(request)
		writer.Write([]byte(tmpl))
	}, "GET")

	// 欢迎路由
	h.AddHandleFunc("welcome", "/"+h.routerPrefix+"/welcome", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("welcome to admin"))
	}, "GET")
}

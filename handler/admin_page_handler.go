package handler

import (
	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/template"
	"net/http"
)

// AdminPageHandler 后台页面处理器
func AdminPageHandler(router *mux.Router) {
	// 添加默认path路由
	router.Path("/index").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.LayoutPageRender(request)
		writer.Write([]byte(tmpl))
	}).Methods("GET")

	// 欢迎路由
	router.Path("/welcome").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("welcome to admin"))
	}).Methods("GET")
}

package handler

import (
	"github.com/gorilla/mux"
	"github.com/hl540/model-admin/model_page/table_page"
	"github.com/hl540/model-admin/template"
	"net/http"
)

// TablePageHandler 列表页面处理器
func TablePageHandler(router *mux.Router, name string, table *table_page.Table) {
	// 列表路由处理器
	router.Path("/" + name + "/list").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.TablePageRender(table, request)
		writer.Write([]byte(tmpl))
	}).Methods("GET")
}

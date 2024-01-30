package handler

import (
	"fmt"
	"github.com/hl540/model-admin/model_page/detail"
	"github.com/hl540/model-admin/model_page/form"
	"github.com/hl540/model-admin/model_page/table"
	"github.com/hl540/model-admin/template"
	"net/http"
)

// 添加模型表格页面处理器
func (h *Handler) addModelTablePageHandler(name string, table *table.Table) {
	path := fmt.Sprintf("/%s/%s/list", h.routerPrefix, name)
	h.AddHandleFunc(name, path, func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.TablePageRender(table, request)
		writer.Write([]byte(tmpl))
	}, "GET")
}

// 添加模型详情页面处理器
func (h *Handler) addModelDetailPageHandler(name string, tablePage *detail.Detail) {
}

// 添加模型新增页面处理器
func (h *Handler) addModelNewPageHandler(name string, tablePage *form.Form) {
}

// 添加模型编辑页面处理器
func (h *Handler) addModelEditPageHandler(name string, tablePage *form.Form) {
}

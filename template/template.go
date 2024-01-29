package template

import (
	"html/template"
	"net/http"

	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page/table"
)

// LayoutPageRender 首页页面模板渲染，它是界面的最外层，包裹其他页面
func LayoutPageRender(req *http.Request) template.HTML {
	render := GetRender(config.GetTemplateName())
	return render.LayoutPageRender(req)
}

// TablePageRender 表格页面模板渲染
func TablePageRender(tableModel *table.Table, req *http.Request) template.HTML {
	render := GetRender(config.GetTemplateName())
	return render.TablePageRender(tableModel, req)
}

// ErrorPageTemplate 错误页面模板渲染
func ErrorPageTemplate(err error, req *http.Request) template.HTML {
	render := GetRender(config.GetTemplateName())
	return render.ErrorPageRender(err, req)
}

// 模板渲染器集合
var templateRender = map[string]Render{}

// AddRender 设置一个模板渲染器
func AddRender(name string, render Render) {
	templateRender[name] = render
}

// GetRender 获取一个模板渲染器
func GetRender(name string) Render {
	return templateRender[name]
}

// Render 模板渲染器
type Render interface {
	// LayoutPageRender 首页页面模板渲染，它是界面的最外层，包裹其他页面
	LayoutPageRender(req *http.Request) template.HTML
	// TablePageRender 表格页面模板渲染
	TablePageRender(tableModel *table.Table, req *http.Request) template.HTML
	// ErrorPageRender 错误页面模板渲染
	ErrorPageRender(err error, req *http.Request) template.HTML
}

package template

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page/table_page"
)

// 模板渲染器集合
var templateSet = make(map[string]HTMLTemplateRender)

// AddTemplate 新增一个模板
func AddTemplate(name string, tmpl HTMLTemplateRender) {
	templateSet[name] = tmpl
}

// GetTemplate 获取一个模板
func GetTemplate(name string) (HTMLTemplateRender, error) {
	tmpl, ok := templateSet[name]
	if !ok {
		return nil, fmt.Errorf("the %s template does not exist", name)
	}
	return tmpl, nil
}

// GetDefaultTemplate 获取配置中的模板名称
func GetDefaultTemplate() (HTMLTemplateRender, error) {
	return GetTemplate(config.GetTemplateName())
}

// HTMLTemplateRender 模板渲染器
type HTMLTemplateRender interface {
	// GetTemplate 获取模板实例
	GetTemplate() *template.Template
	// GetTemplateFiles 获取区别模板文件名称
	GetTemplateFiles() []string
	// LayoutPageRender 首页渲染
	LayoutPageRender(ctx *gin.Context)
	// TablePageRender 表格页面渲染
	TablePageRender(ctx *gin.Context, tableModel *table_page.Table, data *table_page.TableData)
}

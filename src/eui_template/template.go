package eui_template

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/handler"
	"github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"github.com/hl540/model-admin/tools"
	"github.com/pkg/errors"
)

// EUITemplateRender eui模板
type EUITemplateRender struct {
	templatePath string
}

const Name = "eui_template"

func init() {
	_, file, _, _ := runtime.Caller(0)
	absPath := filepath.Dir(file)
	// 注册渲染器
	template2.AddRender(Name, &EUITemplateRender{
		templatePath: absPath + "/template",
	})
	// 注册附加资源路由
	fileServer := http.FileServer(http.Dir(filepath.Dir(absPath + "/template/static/")))
	handler.AddExtraHandler("/eui_static/", http.StripPrefix("/eui_static/", fileServer))
}

// LayoutPageRender 首页页面模板渲染
func (t *EUITemplateRender) LayoutPageRender(req *http.Request) template.HTML {
	// 编译模板内容
	templateFiles := []string{t.templatePath + "/layout.tmpl"}
	htmlStr, err := tools.ExecuteTemplateFile("layout", templateFiles, nil)
	if err != nil {
		return t.ErrorPageRender(err, req)
	}
	return htmlStr
}

// TablePageTemplateData 表格页面模板参数
type TablePageTemplateData struct {
	TablePageKey string            // 当前模型标记
	Title        string            // 标题
	Columns      []map[string]any  // 列
	Data         [][]template.HTML // 数据
	Count        int64             // 总数
}

// TablePageRender 表格页面模板渲染
func (t *EUITemplateRender) TablePageRender(tableModel *table.Table, req *http.Request) template.HTML {
	// 解析参数
	param := table.ParseGetDataParam(req)
	// 加载模板数据
	data, count, err := tableModel.GetTmplData(param)
	if err != nil {
		return t.ErrorPageRender(err, req)
	}
	// 模板数据
	tmplData := TablePageTemplateData{
		TablePageKey: tools.MD5(tableModel.TableName),
		Title:        tableModel.Title,
		Columns:      make([]map[string]any, 0),
		Data:         data,
		Count:        count,
	}
	for _, col := range tableModel.GetColumns() {
		tmplData.Columns = append(tmplData.Columns, map[string]any{
			"name":     col.Title,
			"field":    col.Name,
			"checked":  true,
			"disabled": col.Primary,
		})
	}
	// 编译模板内容
	templateFiles := []string{
		t.templatePath + "/model_table.tmpl",
		t.templatePath + "/table_search.tmpl",
	}
	htmlStr, err := tools.ExecuteTemplateFile("model_table", templateFiles, tmplData)
	if err != nil {
		return t.ErrorPageRender(err, req)
	}
	return htmlStr
}

// ErrorPageRender 表格页面模板生成
func (t *EUITemplateRender) ErrorPageRender(err error, req *http.Request) template.HTML {
	// 加载模板数据
	tmplData := make(map[string]any)
	tmplData["debug"] = config.GetDebugConf().Enable
	// debug模式
	if config.GetDebugConf().Enable {
		// 获取错误堆栈信息
		stack := fmt.Sprintf("%+v", errors.WithStack(err))
		tmplData["error_stacks"] = strings.Split(stack, "\n")
	}
	tmplData["error_message"] = err.Error()
	// 编译模板内容
	templateFiles := []string{t.templatePath + "/error.tmpl"}
	return tools.ExecuteTemplateFileNoError("error", templateFiles, tmplData)
}

package bootstrap_admin_ui

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"github.com/hl540/model-admin/tools"
	"github.com/pkg/errors"
)

//go:embed templates/*.tmpl
var templateFs embed.FS

// 初始化Render
var render = &BootstrapAdminRender{}

// DEV 开发模式，可以热加载模板，方便调试
var DEV = true

func init() {
	// 加载模板
	template, err := template.ParseFS(templateFs, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}
	render.template = template
	// 注册渲染器
	template2.AddRender("bootstrap_admin_ui", render)
}

// BootstrapAdminRender bootstrap-admin模板
// https://www.bootstrap-admin.top/
type BootstrapAdminRender struct {
	template *template.Template
}

// LoadTemplate 加载模板
func (r *BootstrapAdminRender) LoadTemplate() (*template.Template, error) {
	if DEV {
		// 获取绝对路径
		_, file, _, _ := runtime.Caller(0)
		fs := os.DirFS(filepath.Dir(file))
		return template.ParseFS(fs, "templates/*.tmpl")
	}
	return r.template, nil
}

// LayoutPageRender 首页页面模板渲染
func (r *BootstrapAdminRender) LayoutPageRender(req *http.Request) template.HTML {
	// 编译模板内容
	template := template.Must(r.LoadTemplate())
	htmlStr, err := tools.ExecuteTemplateFile(template, "layout", nil)
	if err != nil {
		return r.ErrorPageRender(err, req)
	}
	return htmlStr
}

// TablePageTemplateData 表格页面模板参数
type TablePageTemplateData struct {
	Title            string           // 标题
	Columns          []map[string]any // 列
	Data             []map[string]any // 数据
	Count            int64            // 总数
	FixedLeftNumber  int              // 左侧固定列数
	FixedRightNumber int              // 右侧固定列数
}

// TablePageRender 表格页面模板渲染
func (r *BootstrapAdminRender) TablePageRender(tableModel *table.Table, req *http.Request) template.HTML {
	// 解析参数
	param := table.ParseGetDataParam(req)
	// 加载模板数据
	data, count, err := tableModel.GetTmplData(param)
	if err != nil {
		return r.ErrorPageRender(err, req)
	}
	// 模板数据
	tmplData := TablePageTemplateData{
		Title:            tableModel.Title,
		Columns:          make([]map[string]any, 0, len(tableModel.GetColumns())),
		Data:             make([]map[string]any, 0, len(data)),
		Count:            count,
		FixedLeftNumber:  tableModel.FixedLeftNumber,
		FixedRightNumber: tableModel.FixedRightNumber,
	}
	for _, row := range data {
		rowData := make(map[string]any)
		for k, v := range row {
			rowData[k] = v.RowValue
		}
		tmplData.Data = append(tmplData.Data, rowData)
	}
	for _, col := range tableModel.GetColumns() {
		tmplData.Columns = append(tmplData.Columns, map[string]any{
			"title":      col.Title,
			"field":      col.Name,
			"align":      "center",
			"switchable": !col.Primary,
			"visible":    !col.Hide, // 是否显示该列
		})
	}
	// 编译模板内容
	template := template.Must(r.LoadTemplate())
	htmlStr, err := tools.ExecuteTemplateFile(template, "model_table", tmplData)
	if err != nil {
		return r.ErrorPageRender(err, req)
	}
	return htmlStr
}

// ErrorPageRender 表格页面模板生成
func (r *BootstrapAdminRender) ErrorPageRender(err error, req *http.Request) template.HTML {
	// 加载模板数据
	tmplData := make(map[string]any)
	tmplData["debug"] = config.GetDebug().Enable
	// debug模式
	if config.GetDebug().Enable {
		// 获取错误堆栈信息
		stack := fmt.Sprintf("%+v", errors.WithStack(err))
		tmplData["error_stacks"] = strings.Split(stack, "\n")
	}
	tmplData["error_message"] = err.Error()
	// 编译模板内容
	template := template.Must(r.LoadTemplate())
	return tools.ExecuteTemplateFileNoError(template, "error", tmplData)
}

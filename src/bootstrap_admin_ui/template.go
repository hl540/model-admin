package bootstrap_admin_ui

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page/table"
	template2 "github.com/hl540/model-admin/template"
	"github.com/hl540/model-admin/tools"
	"github.com/pkg/errors"
)

const Name = "bootstrap_admin_ui"

func init() {
	_, file, _, _ := runtime.Caller(0)
	absPath := filepath.Dir(file)
	// 注册渲染器
	template2.AddRender(Name, &BootstrapAdminRender{
		templatePath: absPath + "/template",
	})
}

// BootstrapAdminRender bootstrap-admin模板
// https://www.bootstrap-admin.top/
type BootstrapAdminRender struct {
	templatePath string
}

// LayoutPageRender 首页页面模板渲染
func (r *BootstrapAdminRender) LayoutPageRender(req *http.Request) template.HTML {
	// 编译模板内容
	templateFiles := []string{
		r.templatePath + "/common.tmpl",
		r.templatePath + "/layout.tmpl",
	}
	htmlStr, err := tools.ExecuteTemplateFile("layout", templateFiles, nil)
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
	templateFiles := []string{r.templatePath + "/model_table.tmpl"}
	htmlStr, err := tools.ExecuteTemplateFile("model_table", templateFiles, tmplData)
	if err != nil {
		return r.ErrorPageRender(err, req)
	}
	return htmlStr
}

// ErrorPageRender 表格页面模板生成
func (r *BootstrapAdminRender) ErrorPageRender(err error, req *http.Request) template.HTML {
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
	templateFiles := []string{r.templatePath + "/error.tmpl"}
	return tools.ExecuteTemplateFileNoError("error", templateFiles, tmplData)
}

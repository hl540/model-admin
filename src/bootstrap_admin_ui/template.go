package bootstrap_admin_ui

import (
	"embed"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/model_page"
	template2 "github.com/hl540/model-admin/template"
	"github.com/hl540/model-admin/utils"
)

//go:embed templates/*.tmpl
var templateFs embed.FS

func init() {
	// 加载模板
	temp, err := template.ParseFS(templateFs, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	// 获取所有模板文件名称
	_, runtimeFile, _, _ := runtime.Caller(0)
	absPath := filepath.Join(filepath.Dir(runtimeFile), "templates")
	fileNames, err := utils.GetFilesName(absPath)
	if err != nil {
		panic(err)
	}

	// 注册模板
	template2.AddTemplate("bootstrap_admin_ui", &BootstrapAdminUI{
		template:      temp,
		templateFiles: fileNames,
	})
}

// ErrorPageRender 表格页面模板生成
// func (r *BootstrapAdminRender) ErrorPageRender(err error, req *http.Request) template.HTML {
// 	// 加载模板数据
// 	tmplData := make(map[string]any)
// 	tmplData["debug"] = config.GetDebug().Enable
// 	// debug模式
// 	if config.GetDebug().Enable {
// 		// 获取错误堆栈信息
// 		stack := fmt.Sprintf("%+v", errors.WithStack(err))
// 		tmplData["error_stacks"] = strings.Split(stack, "\n")
// 	}
// 	tmplData["error_message"] = err.Error()
// 	// 编译模板内容
// 	template := template.Must(r.LoadTemplate())
// 	return utils.ExecuteTemplateFileNoError(template, "error", tmplData)
// }

type BootstrapAdminUI struct {
	template      *template.Template
	templateFiles []string
}

// GetTemplate 获取模板实例
func (b *BootstrapAdminUI) GetTemplate() *template.Template {
	return b.template
}

// GetTemplateFiles 获取区别模板文件名称
func (b *BootstrapAdminUI) GetTemplateFiles() []string {
	return b.templateFiles
}

// LayoutPageRender 首页渲染
func (b *BootstrapAdminUI) LayoutPageRender(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "layout.tmpl", nil)
}

// TablePageTemplateData 表格页面模板参数
type TablePageTemplateData struct {
	Title            string           // 标题
	Columns          []map[string]any // 列
	PrimaryColumn    string           // 唯一列
	Data             []map[string]any // 数据
	Count            int64            // 总数
	Size             int              // 数量
	Page             int              // 页码
	SortName         string           // 排序
	SortOrder        string           // 排序
	FixedLeftNumber  int              // 左侧固定列数
	FixedRightNumber int              // 右侧固定列数
}

// TablePageRender 表格页面渲染
func (b *BootstrapAdminUI) TablePageRender(ctx *gin.Context, tableModel *model_page.Table, tableData *model_page.TableData) {
	tmplData := TablePageTemplateData{
		Title:            tableModel.Title,
		Columns:          make([]map[string]any, 0, len(tableModel.GetColumns())),
		PrimaryColumn:    tableModel.PrimaryColumn.Name,
		Data:             tableData.Rows,
		Count:            tableData.Count,
		Size:             tableData.Size,
		Page:             tableData.Page,
		SortName:         ctx.Query("_sort_name"),
		SortOrder:        ctx.Query("_sort_order"),
		FixedLeftNumber:  tableModel.FixedLeftNumber,
		FixedRightNumber: tableModel.FixedRightNumber,
	}
	for _, col := range tableModel.GetColumns() {
		tmplData.Columns = append(tmplData.Columns, map[string]any{
			"title":     col.Title,
			"field":     col.Name,
			"align":     "center",
			"visible":   !col.IsHide, // 是否显示该列
			"formatter": col.ShowFormatName,
			"sortable":  col.IsSort,
		})
	}
	ctx.HTML(http.StatusOK, "model_table_new.tmpl", tmplData)
}

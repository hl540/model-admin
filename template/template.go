package template

import (
	"fmt"
	"github.com/hl540/model-admin/config"
	"github.com/hl540/model-admin/model_page/table"
	"github.com/hl540/model-admin/tools"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"net/http"
	"strings"
)

var templatePath = "./template"

func SetTemplatePath(path string) {
	templatePath = path
}

// TableTemplate 表格页面模板生成
func TableTemplate(tableModel *table.Table, req *http.Request) template.HTML {
	// 解析参数
	param := table.ParseGetDataParam(req)
	// 加载表格模板文件
	tmpl, err := template.New("tableModel.tmpl").ParseFiles(templatePath + "/tableModel.tmpl")
	if err != nil {
		return ErrorTemplate(err, req)
	}
	// 加载模板值
	tmplData, err := tableModel.GetTmplData(param)
	if err != nil {
		return ErrorTemplate(err, req)
	}
	//编译模板
	htmlStr, err := tools.ExecuteTemplate(tmpl, tmplData)
	if err != nil {
		return ErrorTemplate(err, req)
	}
	return htmlStr
}

type ErrorTmplData struct {
	Debug  bool
	Error  string
	Url    string
	Method string
	Header http.Header
	Body   string
	Stacks []string
}

// ErrorTemplate 错误页面模板生成
func ErrorTemplate(err error, req *http.Request) template.HTML {
	// 加载表格模板文件
	tmpl, tmplErr := template.New("error.tmpl").ParseFiles(templatePath + "/error.tmpl")
	if tmplErr != nil {
		return template.HTML(err.Error())
	}
	//编译模板
	body, _ := io.ReadAll(req.Body)
	stack := fmt.Sprintf("%+v", errors.WithStack(err))
	stacks := strings.Split(stack, "\n")
	return tools.ExecuteTemplateNoError(tmpl, ErrorTmplData{
		Debug:  config.GetDebugConf().Enable,
		Error:  err.Error(),
		Url:    req.URL.String(),
		Method: req.Method,
		Header: req.Header,
		Body:   string(body),
		Stacks: stacks,
	})
}

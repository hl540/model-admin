package template

import (
	"bytes"
	"github.com/hl540/model-admin/model_page/table"
	"html/template"
)

var templatePath = "./template"

func SetTemplatePath(path string) {
	templatePath = path
}

// TableTemplate 表格模板生成
func TableTemplate(table *table.Table, param *table.GetDataParam) (template.HTML, error) {
	// 加载表格模板文件
	tmpl, err := template.New("table.tmpl").ParseFiles(templatePath + "/table.tmpl")
	if err != nil {
		return "", err
	}
	// 加载模板值
	tmplData, err := table.GetTmplData(param)
	if err != nil {
		return "", err
	}
	//编译模板
	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, tmplData); err != nil {
		return "", err
	}
	return template.HTML(buf.String()), nil
}

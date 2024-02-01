package utils

import (
	"html/template"
	"strings"
)

// ExecuteTemplateString 从字符串中加载并执行模板
func ExecuteTemplateString(name, text string, data interface{}) template.HTML {
	//  加载模板
	tmpl := template.Must(template.New(name).Parse(text))
	// 编译模板
	var resultBuilder strings.Builder
	if err := tmpl.Execute(&resultBuilder, data); err != nil {
		return template.HTML(err.Error())
	}
	htmlStr := resultBuilder.String()
	return template.HTML(htmlStr)
}

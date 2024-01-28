package tools

import (
	"html/template"
	"strings"
)

// ExecuteTemplate 执行模板，并返回结果字符串
func ExecuteTemplate(tmpl *template.Template, data interface{}) (template.HTML, error) {
	var resultString string
	var resultBuilder strings.Builder
	err := tmpl.Execute(&resultBuilder, data)
	if err != nil {
		return "", err
	}
	resultString = resultBuilder.String()
	return template.HTML(resultString), nil
}

// ExecuteTemplateNoError 执行模板，并返回结果字符串，如果有错误将错误内容作为结果返回
func ExecuteTemplateNoError(tmpl *template.Template, data interface{}) template.HTML {
	htmlStr, err := ExecuteTemplate(tmpl, data)
	if err != nil {
		return template.HTML(err.Error())
	}
	return htmlStr
}

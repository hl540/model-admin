package tools

import (
	"html/template"
	"strings"
)

// ExecuteTemplateFile 加载并执行模板，并返回结果字符串和错误
func ExecuteTemplateFile(name string, tmplFiles []string, data interface{}) (template.HTML, error) {
	//  加载模板
	tmpl, err := template.New(name).ParseFiles(tmplFiles...)
	if err != nil {
		return "", err
	}
	// 编译模板
	var resultBuilder strings.Builder
	if err = tmpl.Execute(&resultBuilder, data); err != nil {
		return "", err
	}
	htmlStr := resultBuilder.String()
	return template.HTML(htmlStr), nil
}

// ExecuteTemplateFileNoError 加载并执行模板，并返回结果字符串，如果有错误将错误内容作为结果返回
func ExecuteTemplateFileNoError(name string, tmplFiles []string, data interface{}) template.HTML {
	htmlStr, err := ExecuteTemplateFile(name, tmplFiles, data)
	if err != nil {
		return template.HTML(err.Error())
	}
	return htmlStr
}

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

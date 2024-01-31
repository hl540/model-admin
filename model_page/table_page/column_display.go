package table_page

import (
	"fmt"
	"html/template"
	url2 "net/url"

	"github.com/hl540/model-admin/tools"
	"github.com/spf13/cast"
)

// DisplayImage 展示为图片
func (c *Column) DisplayImage(width, height int) {
	if width == 0 {
		width = 50
	}
	if height == 0 {
		height = width
	}
	c.SetDisplayFn(func(value map[string]any) template.HTML {
		text := `<img src="{{.src}}" alt="{{.alt}}" width="{{.width}}px" height="{{.height}}px">`
		return tools.ExecuteTemplateString("ColumnDisplayImage", text, map[string]any{
			"src":    cast.ToString(value[c.Name]),
			"alt":    value[c.Name],
			"width":  width,
			"height": height,
		})
	})
}

// DisplayDateTime 展示为日期
func (c *Column) DisplayDateTime(layout string) {
	c.SetDisplayFn(func(value map[string]any) template.HTML {
		valueStr := cast.ToString(value[c.Name])
		time, err := cast.StringToDate(valueStr)
		if err != nil {
			return template.HTML(err.Error())
		}
		htmlStr := time.Format(layout)
		return template.HTML(htmlStr)
	})
}

// DisplayLink 展示为链接
func (c *Column) DisplayLink(urlFormat string, target string) {
	c.SetDisplayFn(func(rowValue map[string]any) template.HTML {
		url, _ := url2.Parse(urlFormat)
		query := url.Query()
		for k := range query {
			v := query.Get(k)
			vName := v[1 : len(v)-1]
			query.Set(k, cast.ToString(rowValue[vName]))
		}
		url.RawQuery = query.Encode()
		htmlStr := fmt.Sprintf(`<a href="%s" target="%s">%s</a>`, url.String(), target, c.displayHtml)
		return template.HTML(htmlStr)
	})
}

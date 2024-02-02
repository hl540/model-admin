package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/template"
)

// AdminPageHandler 后台页面处理器
func AdminPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tmpl, err := template.GetDefaultTemplate()
		if err != nil {
			ctx.String(500, err.Error())
		}
		tmpl.LayoutPageRender(ctx)
	}
}

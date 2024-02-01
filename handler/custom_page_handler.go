package handler

import "github.com/gin-gonic/gin"

// CustomPageHandler 自定义页面处理
func CustomPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("page_name")
		ctx.String(200, name)
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/model_page"
	"github.com/hl540/model-admin/template"
)

// TablePageHandler 列表页面处理器
func TablePageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取model_table
		table, err := model_page.GetTable(ctx.Param("mode_name"))
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		result, err := table.GetData(model_page.ParseQueryParam(ctx))
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		tmpl, err := template.GetDefaultTemplate()
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		tmpl.TablePageRender(ctx, table, result)
	}
}

// TablePageHandlerPost 列表页面处理器
func TablePageHandlerPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取model_table
		table, err := model_page.GetTable(ctx.Param("mode_name"))
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		result, err := table.GetData(model_page.ParseQueryParam(ctx))
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		ctx.JSON(200, result)
	}
}

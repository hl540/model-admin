package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hl540/model-admin/template"
)

// func AdminPageHandler(router *mux.Router) {
// 	// 添加默认path路由
// 	router.Path("/index").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		tmpl := template.LayoutPageRender(request)
// 		writer.Write([]byte(tmpl))
// 	}).Methods("GET")
//
// 	// 欢迎路由
// 	router.Path("/welcome").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		writer.Write([]byte("welcome to admin"))
// 	}).Methods("GET")
// }

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

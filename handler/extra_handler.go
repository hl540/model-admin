package handler

import "net/http"

// 额外的附加路由，可能是状态模板自身需要的静态资源
var extraHandler = make(map[string]http.Handler)

// AddExtraHandler 附加一个路由
func AddExtraHandler(path string, handler http.Handler) {
	extraHandler[path] = handler
}

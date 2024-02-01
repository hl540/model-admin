package table_page

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const DefaultSize = 10 // 单页默认数量
const SizeMax = 100    // 单页最大数量
const DefaultPage = 1  // 默认页码

// pagination 解析分页参数
func pagination(page int, size int) (int, int) {
	if page <= 0 {
		page = DefaultPage
	}
	if size <= 0 {
		size = DefaultSize
	}
	if size > SizeMax {
		size = SizeMax
	}
	return page, size
}

// 解析排序参数
func sort(sortStr string) []string {
	if len(sortStr) == 0 {
		return nil
	}
	sortArr := strings.Split(sortStr, ".")
	if len(sortArr) == 1 {
		return append(sortArr, "DESC")
	}
	if sortArr[1] != "DESC" && sortArr[1] != "ASC" {
		sortArr[1] = "DESC"
	}
	return sortArr
}

// QueryParam 表格数据查询参数
type QueryParam struct {
	Filter map[string]any // 筛选参数
	Sort   []string       // 排序参数
	Page   int            // 数量
	Size   int            // 页码
}

// ParseQueryParam 解析查询参数
func ParseQueryParam(ctx *gin.Context) *QueryParam {
	param := &QueryParam{
		Filter: make(map[string]any),
	}
	// 解析分页参数
	page := cast.ToInt(ctx.Query("_page"))
	size := cast.ToInt(ctx.Query("_size"))
	param.Page, param.Size = pagination(page, size)

	// 解析排序参数
	param.Sort = sort(ctx.Query("_sort"))

	// 解析筛选参数
	for k := range ctx.Request.URL.Query() {
		if len(k) > 8 && k[:8] == "_filter_" {
			param.Filter[k[8:]] = ctx.Query(k)
		}
	}
	return param
}

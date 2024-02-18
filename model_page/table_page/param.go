package table_page

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const DefaultSize = 10 // 单页默认数量
const SizeMax = 100    // 单页最大数量
const DefaultPage = 1  // 默认页码

// 解析分页参数
func parsePaginationParam(ctx *gin.Context) (int, int) {
	page := cast.ToInt(ctx.Query("_page"))
	size := cast.ToInt(ctx.Query("_size"))
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
func parseSortParam(ctx *gin.Context) []string {
	sortName := ctx.Query("_sort_name")
	if len(sortName) == 0 {
		return nil
	}
	sortOrder := ctx.Query("_sort_order")
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "DESC"
	}
	return []string{sortName, sortOrder}
}

// 解析筛选参数
func parseFilterParam(ctx *gin.Context) map[string]any {
	filterParam := make(map[string]any)
	for k := range ctx.Request.URL.Query() {
		if len(k) > 8 && k[:8] == "_filter_" {
			filterParam[k[8:]] = ctx.Query(k)
		}
	}
	return filterParam
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
		Filter: parseFilterParam(ctx), // 解析筛选参数
		Sort:   parseSortParam(ctx),   // 解析排序参数
	}
	// 解析分页参数
	param.Page, param.Size = parsePaginationParam(ctx)
	return param
}

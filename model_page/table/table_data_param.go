package table

import (
	"net/http"
	"strconv"
	"strings"
)

const PaginationSize = 10     // 单页默认数量
const PaginationSizeMax = 100 // 单页最大数量
const PaginationPage = 1      // 默认页码

// Pagination 分页参数
type Pagination struct {
	Size int64 // 数量
	Page int64 // 页码
}

// NewPagination 创建分页参数
func NewPagination(page int64, size int64) *Pagination {
	if page <= 0 {
		page = PaginationPage
	}
	if size <= 0 {
		size = PaginationSize
	}
	if size > PaginationSizeMax {
		size = PaginationSizeMax
	}
	return &Pagination{
		Size: size,
		Page: page,
	}
}

type GetDataParam struct {
	Filter     map[string]any // 筛选参数
	Sort       []string       // 排序参数
	Pagination *Pagination    // 分页参数
}

func ParseGetDataParam(req *http.Request) *GetDataParam {
	param := &GetDataParam{}
	query := req.URL.Query()
	// 解析分页参数
	page, _ := strconv.ParseInt(query.Get("_page"), 10, 64)
	size, _ := strconv.ParseInt(query.Get("_size"), 10, 64)
	param.Pagination = NewPagination(page, size)

	// 解析排序参数
	sort := query.Get("_sort")
	if sort != "" {
		param.Sort = strings.Split(sort, ".")
		if len(param.Sort) == 1 {
			param.Sort = append(param.Sort, "DESC")
		}
	}

	// 解析筛选参数
	param.Filter = make(map[string]any)
	for k := range query {
		if len(k) > 8 && k[:8] == "_filter_" {
			param.Filter[k[8:]] = query.Get(k)
		}
	}
	return param
}

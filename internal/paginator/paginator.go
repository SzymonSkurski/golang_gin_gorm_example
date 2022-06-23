package paginator

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

const MAX_PAGE_SIZE = 100
const MIN_PAGE_SIZE = 10

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := getPageNo(r)
		pageSize := getPageSize(r)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateInfo(r *http.Request) map[string]int {
	p := map[string]int{
		"pageSize": getPageSize(r),
		"page":     getPageNo(r),
	}
	return p
}

func getPageSize(r *http.Request) int {
	q := r.URL.Query()
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > MAX_PAGE_SIZE:
		pageSize = MAX_PAGE_SIZE
	case pageSize <= 0:
		pageSize = MIN_PAGE_SIZE
	}
	return pageSize
}

func getPageNo(r *http.Request) int {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}
	return page
}

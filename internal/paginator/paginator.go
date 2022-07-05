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

func PaginateInfo(r *http.Request, db *gorm.DB) map[string]int {
	size := getPageSize(r)
	p := map[string]int{
		"pageSize": size,
		"pageNo":   getPageNo(r),
		"pages":    getPagesNumber(size, int(countPages(db))),
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

func getPagesNumber(size int, count int) int {
	if size < 1 || count < 1 {
		return 0
	}
	c := count / size
	if size%count != 0 {
		c++
	}
	return c
}

func countPages(db *gorm.DB) int64 {
	var count int64
	db.Count(&count)
	return count
}

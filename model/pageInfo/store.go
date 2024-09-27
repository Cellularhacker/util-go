package pageInfo

import (
	"net/http"
	"strconv"
)

var (
	QueryCurrentPage = "current_page"
	QueryPageSize    = "page_size"
)

func (pr *Request) ToPageInfo(totalCount int64) *Response {
	if totalCount == -1 || pr == nil {
		return nil
	}

	return &Response{
		TotalPage:   (totalCount-1)/pr.Size + 1,
		TotalCount:  totalCount,
		CurrentPage: pr.Current,
		PageSize:    pr.Size,
	}
}

func Parse(r *http.Request) (*Request, *string) {
	// MARK: Preparing variables...
	q := r.URL.Query()

	// MARK: Query - current_page
	currentPage, err := strconv.ParseInt(q.Get(QueryCurrentPage), 10, 64)
	if err != nil || currentPage <= 0 {
		return nil, &QueryCurrentPage
	}
	// MARK: Query - page_size
	pageSize, err := strconv.ParseInt(q.Get(QueryPageSize), 10, 64)
	if err != nil || pageSize < 0 {
		return nil, &QueryPageSize
	}

	return &Request{Current: currentPage, Size: pageSize, Skip: pageSize * (currentPage - 1)}, nil
}

package util

import (
	"net/http"
	"strconv"
)

var (
	// QueryCurrentPage - PageInfo
	QueryCurrentPage = "current_page"
	// QueryPageSize - PageInfo
	QueryPageSize = "page_size"
)

// Request - PageInfo
type Request struct {
	Current int64 `json:"current" bson:"current"`
	Size    int64 `json:"size" bson:"size"`
	Skip    int64 `json:"skip" bson:"skip"`
}

// Response - PageInfo
type Response struct {
	TotalPage   int64 `json:"total_page" bson:"total_page"`
	TotalCount  int64 `json:"total_count" bson:"total_count"`
	CurrentPage int64 `json:"current_page" bson:"current_page"`
	PageSize    int64 `json:"page_size" bson:"page_size"`
}

// MARK: Methods and utility functions.

// ToPageInfo - PageInfo
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

// Parse - PageInfo
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

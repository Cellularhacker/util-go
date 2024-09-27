package pageInfo

type Request struct {
	Current int64 `json:"current" bson:"current"`
	Size    int64 `json:"size" bson:"size"`
	Skip    int64 `json:"skip" bson:"skip"`
}

type Response struct {
	TotalPage   int64 `json:"total_page" bson:"total_page"`
	TotalCount  int64 `json:"total_count" bson:"total_count"`
	CurrentPage int64 `json:"current_page" bson:"current_page"`
	PageSize    int64 `json:"page_size" bson:"page_size"`
}

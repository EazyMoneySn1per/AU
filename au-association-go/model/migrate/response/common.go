package response

type PageResult struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

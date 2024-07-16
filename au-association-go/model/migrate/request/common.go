package request

// Paging common input parameter structure
type PageInfo struct {
	Page  int `json:"page" form:"page"`   // 页码
	Limit int `json:"limit" form:"limit"` // 每页大小
}

// Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

type Empty struct{}

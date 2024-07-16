package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `uri:"page" form:"page" json:"page"`       // 页码
	PageSize int `uri:"limit" form:"pageSize" json:"limit"` // 每页大小
	//Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

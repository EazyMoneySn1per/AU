package request

type MPModule struct {
	Name  string `json:"name" form:"click"`
	Click string `json:"click" form:"click"`
}

package Request

type ActivityInfo struct {
	Description string `json:"description" form:"description"`
	Date        string `json:"date" form:"date"`
	Name        string `json:"name" form:"name"`
	Filed       int    `json:"filed" form:"filed"`
	Assid       string `json:"assid" form:"assid"`
	Step        int    `json:"step" form:"step"`
}

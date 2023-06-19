package schemas

type ListCaseIn struct {
	//CategoryId string `uri:"category_id"` //只能用ShouldBindUri获取
	PageQuery
	Keyword string `json:"keyword" form:"keyword"`
}

type CreateCaseIn struct {
	Title   string `json:"title" form:"title" binding:"required,min=3,max=75"`
	Content string `json:"content" form:"content" binding:"required"`
	//CategoryId  string `json:"category_id" uri:"category_id"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	Images      string `json:"images" form:"images"`
	Sort        int    `json:"sort" form:"sort"`
	Url         string `json:"url" form:"url"`
	Status      int    `json:"status" form:"status" binding:"oneof=0 1"`
}

type DetailCaseIn struct {
	CaseId int `json:"case_id" uri:"case_id" binding:"required,gt=0"`
}

type ChangeCaseHeaderIn struct {
	CurrentCase
	CurrentCategory
}

type ChangeCaseIn struct {
	Title   string `json:"title" form:"title" binding:"omitempty,min=3,max=75"`
	Content string `json:"content" form:"content"`
	//CategoryId  string `json:"category_id" uri:"category_id"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	Images      string `json:"images" form:"images"`
	Sort        int    `json:"sort" form:"sort"`
	Url         string `json:"url" form:"url"`
	Status      int    `json:"status" form:"status" binding:"oneof=0 1"`
}

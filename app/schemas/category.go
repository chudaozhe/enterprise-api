package schemas

type ListCategoryIn struct {
	Type int `json:"type" form:"type"`
}

type CreateCategoryIn struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Type     int    `json:"type" form:"type"`
	Memo     string `json:"memo" form:"memo"`
	ParentId int    `json:"parent_id" form:"parent_id"`
	Sort     int    `json:"sort" form:"sort"`
}

type ChangeCategoryIn struct {
	Name     string `json:"name" form:"name"`
	Type     int    `json:"type" form:"type"`
	Memo     string `json:"memo" form:"memo"`
	ParentId int    `json:"parent_id" form:"parent_id"`
	Sort     int    `json:"sort" form:"sort"`
}

package schemas

type ListFlashIn struct {
	PageQuery
}

type CreateFlashIn struct {
	Image  string `json:"image" form:"image" binding:"required"`
	Title  string `json:"title" form:"title"`
	Url    string `json:"url" form:"url"`
	Sort   int    `json:"sort" form:"sort"`
	Status int    `json:"status" form:"status" binding:"oneof=0 1"`
}

type ChangeFlashIn struct {
	Image  string `json:"image" form:"image" binding:"required"`
	Title  string `json:"title" form:"title"`
	Url    string `json:"url" form:"url"`
	Sort   int    `json:"sort" form:"sort"`
	Status int    `json:"status" form:"status" binding:"oneof=0 1"`
}

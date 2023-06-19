package schemas

type ListShortcutIn struct {
	PageQuery
}

type CreateShortcutIn struct {
	Title  string `json:"title" form:"title" binding:"required"`
	Image  string `json:"image" form:"image"`
	Type   int    `json:"type" form:"type"`
	Url    string `json:"url" form:"url"`
	Sort   int    `json:"sort" form:"sort"`
	Status int    `json:"status" form:"status" binding:"oneof=0 1"`
}

type ChangeShortcutIn struct {
	Title  string `json:"title" form:"title" binding:"required"`
	Image  string `json:"image" form:"image"`
	Type   int    `json:"type" form:"type"`
	Url    string `json:"url" form:"url"`
	Sort   int    `json:"sort" form:"sort"`
	Status int    `json:"status" form:"status" binding:"oneof=0 1"`
}

package schemas

type ListPageIn struct {
	PageQuery
}

type CreatePageIn struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Image   string `json:"image" form:"image"`
}

type ChangePageHeaderIn struct {
	CurrentCategory
	CurrentPage
}
type ChangePageIn struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Image   string `json:"image" form:"image"`
}

type DetailPageIn struct {
	CurrentPage
}

package schemas

type CreateGroupIn struct {
	Name string `json:"name" form:"name" binding:"required"`
	Sort int    `json:"sort" form:"sort"`
}

type ChangeGroupIn struct {
	Name string `json:"name" form:"name"`
	Sort int    `json:"sort" form:"sort"`
}

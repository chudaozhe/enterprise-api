package schemas

type ListFileIn struct {
	PageQuery
}
type CreateFileHeaderIn struct {
	CurrentAdmin
	CurrentGroup
}

type CreateFileIn struct {
	Content string `json:"content" form:"content" binding:"required"`
	Title   string `json:"title" form:"title"`
	Type    string `json:"type" form:"type"`
	Size    int    `json:"size" form:"size"`
	Width   int    `json:"width" form:"width"`
	Height  int    `json:"height" form:"height"`
}

type ChangeFileIn struct {
	FileIds string `json:"file_ids" form:"file_ids" binding:"required"`
}

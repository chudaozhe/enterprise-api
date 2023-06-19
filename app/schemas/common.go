package schemas

type PageQuery struct {
	Page int `json:"page" form:"page"`
	Max  int `json:"max" form:"max"`
}

type CurrentArticle struct {
	ArticleId int `json:"article_id" uri:"article_id" binding:"required,gt=0"`
}

type CurrentCase struct {
	CaseId int `json:"case_id" uri:"case_id" binding:"required,gt=0"`
}

type CurrentCategory struct {
	CategoryId int `json:"category_id" uri:"category_id"`
}

type CurrentGroup struct {
	GroupId int `json:"group_id" uri:"group_id"`
}

type CurrentFile struct {
	FileId int `json:"file_id" uri:"file_id" binding:"required,gt=0"`
}

type CurrentFlash struct {
	FlashId int `json:"flash_id" uri:"flash_id" binding:"required,gt=0"`
}

type CurrentPage struct {
	PageId int `json:"page_id" uri:"page_id" binding:"required,gt=0"`
}

type CurrentShortcut struct {
	ShortcutId int `json:"shortcut_id" uri:"shortcut_id" binding:"required,gt=0"`
}

type CurrentAdmin struct {
	AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
}

type TargetAdmin struct {
	ToAid int `json:"to_aid" uri:"to_aid" binding:"required,gt=0"`
}

package schemas

type LoginIn struct {
	Username string `json:"username" form:"username" binding:"required,alphanum,min=3,max=12"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=12"`
}

type FindPasswdIn struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

type DetailAdminIn struct {
	AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
}

type ChangeAdminIn struct {
	Nickname string `json:"nickname" form:"nickname"`
	Email    string `json:"email" form:"email" binding:"omitempty,email"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type ChangePasswdIn struct {
	//AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
	Password    string `json:"password" form:"password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type AvatarIn struct {
	//AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
	Content string `json:"content" form:"content" binding:"required,base64"`
}

type LogoutIn struct {
	AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
}

/**
admin2
*/

type ListAdminIn struct {
	PageQuery
	Keyword string `json:"keyword" form:"keyword"`
}

type CreateAdminIn struct {
	Username string `json:"username" form:"username" binding:"required,alphanum,min=3,max=12"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Nickname string `json:"nickname" form:"nickname"`
	Mobile   string `json:"mobile" form:"mobile"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type DetailAdmin2In struct {
	TargetAdmin
}

type DeleteAdminIn struct {
	CurrentAdmin
	TargetAdmin
}

type ResetPasswdIn struct {
	AdminId int `json:"admin_id" uri:"admin_id" binding:"required,gt=0"`
	ToAid   int `json:"to_aid" uri:"to_aid" binding:"required,gt=0"`
}

type ChangeAuthIn struct {
	Rule string `json:"rule" form:"rule" binding:"required"`
}

type DisableAdminIn struct {
	ToAid int `json:"to_aid" uri:"to_aid" binding:"required,gt=0"`
}

type EnableAdminIn struct {
	ToAid int `json:"to_aid" uri:"to_aid" binding:"required,gt=0"`
}

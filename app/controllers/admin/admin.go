package admin

import (
	"enterprise-api/app/models"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginIn schemas.LoginIn
	if err := c.ShouldBind(&loginIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err := models.LoginByName(loginIn.Username, loginIn.Password)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		token, err := models.SetToken(admin.Id, true, 0, "admin")
		/* jwt
		claims := models.MyCustomClaims{
			Id:       admin.Id,
			Username: admin.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), // 过期时间1星期
				Issuer:    admin.Username,                            // 签发人
			},
		}
		token, err := models.CreateToken(claims)
		*/
		if err != nil {
			core.Error(c, 1, "token设置失败")
			return
		}
		core.Success(c, 0, struct {
			*models.Admin
			Token string `json:"token"`
		}{
			&admin,
			token,
		})
	}
}

func FindPasswd(c *gin.Context) {
	var findPasswdIn schemas.FindPasswdIn
	if err := c.ShouldBind(&findPasswdIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err0 := models.FindByEmail(findPasswdIn.Email)
	if err0 != nil {
		core.Error(c, 404, "对应账号不存在")
		return
	}
	if admin.Status == 1 {
		err := models.ResetPasswd(admin)
		if err != nil {
			core.Error(c, 500, err.Error())
			return
		}
		core.Success(c, 0, gin.H{"reset": true})
	} else {
		core.Error(c, 405, "账号异常, 禁止操作")
	}
}

func Detail(c *gin.Context) {
	var detailAdminIn schemas.DetailAdminIn
	if err := c.ShouldBindUri(&detailAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err := models.FindById(detailAdminIn.AdminId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, admin)
	}
}

func Change(c *gin.Context) {
	var currentAdmin schemas.CurrentAdmin
	if err := c.ShouldBindUri(&currentAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeAdminIn schemas.ChangeAdminIn
	if err := c.ShouldBind(&changeAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err := models.FindById(currentAdmin.AdminId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}
	admin.Nickname = changeAdminIn.Nickname
	admin.Email = changeAdminIn.Email
	admin.Mobile = changeAdminIn.Mobile
	err2 := admin.UpdateAdmin()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func ChangePasswd(c *gin.Context) {
	var currentAdmin schemas.CurrentAdmin
	if err := c.ShouldBindUri(&currentAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changePasswdIn schemas.ChangePasswdIn
	if err := c.ShouldBind(&changePasswdIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err0 := models.FindById(currentAdmin.AdminId)
	if err0 != nil {
		core.Error(c, 1, err0.Error())
	} else {
		err := admin.ChangePasswd(changePasswdIn.Password, changePasswdIn.NewPassword)
		if err != nil {
			core.Error(c, 1, err.Error())
			return
		}
		core.Success(c, 0, gin.H{"update": true})
	}
}

func Avatar(c *gin.Context) {
	var currentAdmin schemas.CurrentAdmin
	if err := c.ShouldBindUri(&currentAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var avatarIn schemas.AvatarIn
	if err := c.ShouldBind(&avatarIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err0 := models.FindById(currentAdmin.AdminId)
	if err0 != nil {
		core.Error(c, 1, "管理员不存在")
	} else {
		admin.Avatar = models.SaveAdminAvatar(admin.Id, avatarIn.Content)
		err := admin.UpdateAdmin()
		if err != nil {
			core.Error(c, 1, "修改失败")
			return
		}
		core.Success(c, 0, gin.H{"update": true})
	}
}

func Logout(c *gin.Context) {
	var logoutIn schemas.LogoutIn
	if err := c.ShouldBindUri(&logoutIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	_, err := models.GetToken(logoutIn.AdminId, "admin")
	if err == nil {
		models.DelToken(logoutIn.AdminId, "admin")
	}
	core.Success(c, 0, gin.H{"delete": true})
}

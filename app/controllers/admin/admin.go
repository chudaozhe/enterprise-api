package admin

import (
	"enterprise-api/app/models"
	"enterprise-api/core"
	"enterprise-api/core/helper"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if len(username) > 0 && len(password) > 0 {
		admin, err := models.LoginByName(username, password)
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
	} else {
		core.Error(c, 400, "参数错误")
	}
}

func Findpasswd(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	if len(email) > 0 && helper.VerifyEmail(email) {
		admin, err0 := models.FindByEmail(email)
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
			core.Success(c, 0, gin.H{"reset": "true"})
		} else {
			core.Error(c, 405, "账号异常, 禁止操作")
		}
	} else {
		core.Error(c, 400, "参数错误")
	}
}

func Detail(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	admin, err := models.FindById(core.ToInt(id))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, admin)
	}
}

func Change(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	adminUpdate := models.Admin{
		Nickname: c.DefaultPostForm("nickname", "-isnil-"),
		Email:    c.DefaultPostForm("email", "-isnil-"),
		Mobile:   c.DefaultPostForm("mobile", "-isnil-"),
	}
	admin, err := models.FindById(core.ToInt(id))
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}

	if adminUpdate.Nickname != "-isnil-" && adminUpdate.Nickname != admin.Nickname {
		admin.Nickname = adminUpdate.Nickname
	}
	if adminUpdate.Email != "-isnil-" && adminUpdate.Email != admin.Email {
		admin.Email = adminUpdate.Email
	}
	if adminUpdate.Mobile != "-isnil-" && adminUpdate.Mobile != admin.Mobile {
		admin.Mobile = adminUpdate.Mobile
	}

	err2 := admin.UpdateAdmin()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{
		"update": "true",
	})
}

func Changepasswd(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	password := c.Request.FormValue("password")
	newPassword := c.Request.FormValue("new_password")
	if len(password) > 0 && len(newPassword) > 0 {
		admin, err0 := models.FindById(core.ToInt(id))
		if err0 != nil {
			core.Error(c, 1, err0.Error())
		} else {
			err := admin.ChangePasswd(password, newPassword)
			if err != nil {
				core.Error(c, 1, err.Error())
				return
			}
			core.Success(c, 0, gin.H{"update": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func Avatar(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	content := c.Request.FormValue("content")
	if len(content) > 0 {
		admin, err0 := models.FindById(core.ToInt(id))
		if err0 != nil {
			core.Error(c, 1, "管理员不存在")
		} else {
			admin.Avatar = models.SaveAdminAvatar(admin.Id, content)
			err := admin.UpdateAdmin()
			if err != nil {
				core.Error(c, 1, "修改失败")
				return
			}
			core.Success(c, 0, gin.H{"update": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func Logout(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	_, err := models.GetToken(core.ToInt(id), "admin")
	if err != nil {
		core.Success(c, 0, gin.H{"delete": true})
	} else {
		models.DelToken(core.ToInt(id), "admin")
		core.Success(c, 0, gin.H{"delete": true})
	}
}

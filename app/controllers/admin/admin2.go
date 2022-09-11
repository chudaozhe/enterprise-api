package admin

import (
	"enterprise-api/app/models"
	"enterprise-api/core"
	"enterprise-api/core/helper"
	"github.com/gin-gonic/gin"
)

func ListAdmin(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "10")
	keyword := c.DefaultQuery("keyword", "")
	count, result, err1 := models.ListOfAdmins(core.ToInt(page), core.ToInt(max), keyword)
	if err1 != nil {
		core.Error(c, 500, err1.Error())
	} else {
		core.Success(c, 0, gin.H{
			"count": count,
			"list":  result,
		})
	}
}

func CreateAdmin(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	email := c.DefaultPostForm("email", "")
	if len(username) > 0 && len(email) > 0 {
		if !helper.VerifyEmail(email) {
			core.Error(c, 400, "E-mail格式不正确")
		} else {
			count := models.FindByNameOrEmail(username, email)
			if count == 0 {
				admin := models.Admin{
					Username: username,
					Nickname: c.DefaultPostForm("nickname", ""),
					Email:    email,
					Mobile:   c.DefaultPostForm("mobile", ""),
					Avatar:   c.DefaultPostForm("avatar", ""),
				}
				id, err := admin.CreateNoPwd()
				if err != nil {
					core.Error(c, 1, err.Error())
					return
				}
				core.Success(c, 0, gin.H{
					"id": id,
				})
			} else {
				core.Error(c, 405, "用户名或E-mail已存在")
			}
		}
	} else {
		core.Success(c, 400, "参数错误")
	}
}

func DetailAdmin(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	if core.ToInt(toAid) > 0 {
		admin, err := models.FindById(core.ToInt(toAid))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, admin)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeAdmin(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	if core.ToInt(toAid) > 0 {
		adminUpdate := models.Admin{
			Nickname: c.DefaultPostForm("nickname", "-isnil-"),
			Email:    c.DefaultPostForm("email", "-isnil-"),
			Mobile:   c.DefaultPostForm("mobile", "-isnil-"),
			Avatar:   c.DefaultPostForm("avatar", "-isnil-"),
		}
		admin, err := models.FindById(core.ToInt(toAid))
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
		if adminUpdate.Avatar != "-isnil-" && adminUpdate.Avatar != admin.Avatar {
			admin.Avatar = adminUpdate.Avatar
		}

		err2 := admin.UpdateAdmin()
		if err2 != nil {
			core.Error(c, 1, err2.Error())
			return
		}
		core.Success(c, 0, gin.H{
			"update": "true",
		})
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func DeleteAdmin(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	if core.ToInt(toAid) > 0 {
		err := models.DeleteAdminById(core.ToInt(toAid))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func Resetpasswd(c *gin.Context) {
	id, _ := c.Params.Get("admin_id")
	toAid, _ := c.Params.Get("to_aid")
	if id == toAid {
		core.Error(c, 405, "不能操作自己账号")
	} else {
		admin, err0 := models.FindById(core.ToInt(toAid))
		if err0 != nil {
			core.Error(c, 404, "对应账号不存在")
			return
		}
		if admin.Status == 1 {
			if !helper.VerifyEmail(admin.Email) {
				core.Error(c, 400, "E-mail格式有误")
				return
			}
			err := models.ResetPasswd(admin)
			if err != nil {
				core.Error(c, 500, err.Error())
				return
			}
			core.Success(c, 0, gin.H{"reset": "true"})
		} else {
			core.Error(c, 405, "账号异常, 禁止操作")
		}
	}
}

func ChangeAuth(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	rule := c.DefaultPostForm("rule", "")
	if core.ToInt(toAid) > 0 && len(rule) > 0 {
		admin := &models.Admin{Id: core.ToInt(toAid)}
		err := admin.ChangeAuth(rule)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"update": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func DisableAdmin(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	if core.ToInt(toAid) > 0 {
		admin := &models.Admin{Id: core.ToInt(toAid)}
		err := admin.ChangeAdminState(0)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"disable": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func EnableAdmin(c *gin.Context) {
	toAid, _ := c.Params.Get("to_aid")
	if core.ToInt(toAid) > 0 {
		admin := &models.Admin{Id: core.ToInt(toAid)}
		err := admin.ChangeAdminState(1)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"enable": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

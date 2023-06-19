package admin

import (
	"enterprise-api/app/models"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"enterprise-api/core/helper"
	"github.com/gin-gonic/gin"
)

func ListAdmin(c *gin.Context) {
	var listAdminIn schemas.ListAdminIn
	if err := c.ShouldBindQuery(&listAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	count, result, err1 := models.ListOfAdmins(listAdminIn.Page, listAdminIn.Max, listAdminIn.Keyword)
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
	var createAdminIn schemas.CreateAdminIn
	if err := c.ShouldBind(&createAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	count := models.FindByNameOrEmail(createAdminIn.Username, createAdminIn.Email)
	if count == 0 {
		admin := models.Admin{
			Username: createAdminIn.Username,
			Nickname: createAdminIn.Nickname,
			Email:    createAdminIn.Email,
			Mobile:   createAdminIn.Mobile,
			Avatar:   createAdminIn.Avatar,
		}
		id, err := admin.CreateNoPwd()
		if err != nil {
			core.Error(c, 1, err.Error())
			return
		}
		core.Success(c, 0, gin.H{"id": id})
	} else {
		core.Error(c, 405, "用户名或E-mail已存在")
	}
}

func DetailAdmin(c *gin.Context) {
	var targetAdmin schemas.TargetAdmin
	if err := c.ShouldBindUri(&targetAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err := models.FindById(targetAdmin.ToAid)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, admin)
	}
}

func ChangeAdmin(c *gin.Context) {
	var targetAdmin schemas.TargetAdmin
	if err := c.ShouldBindUri(&targetAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeAdminIn schemas.ChangeAdminIn
	if err := c.ShouldBind(&changeAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin, err := models.FindById(targetAdmin.ToAid)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}

	admin.Nickname = changeAdminIn.Nickname
	admin.Email = changeAdminIn.Email
	admin.Mobile = changeAdminIn.Mobile
	admin.Avatar = changeAdminIn.Avatar
	err2 := admin.UpdateAdmin()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DeleteAdmin(c *gin.Context) {
	var deleteAdminIn schemas.DeleteAdminIn
	if err := c.ShouldBindUri(&deleteAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := models.DeleteAdminById(deleteAdminIn.ToAid)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"delete": true})
	}
}

func ResetPasswd(c *gin.Context) {
	var resetPasswdIn schemas.ResetPasswdIn
	if err := c.ShouldBindUri(&resetPasswdIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	if resetPasswdIn.AdminId == resetPasswdIn.ToAid {
		core.Error(c, 405, "不能操作自己账号")
		return
	}
	admin, err0 := models.FindById(resetPasswdIn.ToAid)
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
		core.Success(c, 0, gin.H{"reset": true})
	} else {
		core.Error(c, 405, "账号异常, 禁止操作")
	}
}

func ChangeAuth(c *gin.Context) {
	var targetAdmin schemas.TargetAdmin
	if err := c.ShouldBindUri(&targetAdmin); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeAuthIn schemas.ChangeAuthIn
	if err := c.ShouldBind(&changeAuthIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin := &models.Admin{Id: targetAdmin.ToAid}
	err := admin.ChangeAuth(changeAuthIn.Rule)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"update": true})
	}
}

func DisableAdmin(c *gin.Context) {
	var disableAdminIn schemas.DisableAdminIn
	if err := c.ShouldBindUri(&disableAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin := &models.Admin{Id: disableAdminIn.ToAid}
	err := admin.ChangeAdminState(0)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"disable": true})
	}
}

func EnableAdmin(c *gin.Context) {
	var enableAdminIn schemas.EnableAdminIn
	if err := c.ShouldBindUri(&enableAdminIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	admin := &models.Admin{Id: enableAdminIn.ToAid}
	err := admin.ChangeAdminState(1)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"enable": true})
	}
}

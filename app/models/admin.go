package models

import (
	"enterprise-api/app/config"
	customError "enterprise-api/app/models/errors"
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Admin struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Avatar     string `json:"avatar"`
	Init       int    `json:"init"`
	Status     int    `json:"status"`
	Auth       string `json:"auth"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// 设置表名
func (Admin) TableName() string { //默认为结构体名称的复数，即admins
	return "cw_admin"
}

// 添加
func (admin Admin) CreateNoPwd() (id int, err error) {
	//添加数据
	admin.Status = 1
	admin.CreateTime = helper.GetUnix()
	password, salt, err0 := InitPasswd(admin)
	if err0 != nil {
		err = errors.New(err.Error())
		return
	} else {
		admin.Password = password
		admin.Salt = salt
		admin.Init = 1

		result := orm.Db.Create(&admin)
		if result.Error != nil {
			err = result.Error
			return
		}
		return admin.Id, nil
	}
}

// 生成并发送初始密码
func InitPasswd(admin Admin) (password string, salt string, err error) {
	salt = helper.Random(6, "")
	password = helper.Random(6, "")
	//发送email
	conf := config.GetConfig().MailConfig
	replacer := strings.NewReplacer("{username}", admin.Username, "{passwd}", password, "{url}", conf.Url)
	err0 := SendEmail(admin.Email, conf.Title, replacer.Replace(conf.Content))
	if err0 != nil {
		fmt.Println(err0.Error())
		err = errors.New("密码发送失败,请重试")
		return
	}
	//fmt.Println(password, salt)
	return helper.Md5(helper.Md5(password) + salt), salt, nil
}

// 重置密码
func ResetPasswd(admin Admin) (err error) {
	password, salt, err0 := InitPasswd(admin)
	if err0 != nil {
		err = errors.New(err0.Error())
		return
	} else {
		result := orm.Db.Model(&admin).Updates(Admin{Password: password, Salt: salt, Init: 1, UpdateTime: helper.GetUnix()})
		if result.Error != nil {
			err = errors.New("服务器错误")
			return
		}
		return
	}
}

// 详情
func FindById(id int) (admin Admin, err error) {
	result := orm.Db.First(&admin, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("not found")
		return
	}
	return
}

// 管理员详情
func FindByName(username string) (admin Admin, err error) {
	result := orm.Db.First(&admin, "username = ?", username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("not found")
		return
	}
	return
}

// 列表
func ListOfAdmins(page, max int, keyword string) (count int64, admins []*Admin, err error) { //, search Admin
	obj := orm.Db.Offset((page - 1) * max).Limit(max)
	if len(keyword) > 0 {
		obj = obj.Where("CONCAT(`cw_admin`.`username`, `cw_admin`.`nickname`, `cw_admin`.`email`, `cw_admin`.`mobile`) LIKE ?", "%"+keyword+"%")
	}
	err = obj.Select([]string{"id", "username", "nickname", "status", "email", "mobile", "avatar", "init", "auth as rules", "create_time"}).
		Order("cw_admin.id desc").
		Find(&admins).Offset(-1).Limit(-1).Count(&count).Error
	return
}

// 改变管理员状态
func (admin Admin) ChangeAdminState(status int) (err error) {
	if status == 1 || status == 0 {
		err = orm.Db.Model(&admin).Update("status", status).Error
		return
	} else {
		err = errors.New("禁止更新")
		return
	}
}

// 更新
func (admin Admin) UpdateAdmin() (err error) {
	admin.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(admin).Error
	return
}

// 删除
func DeleteAdminById(id, toId int) (err error) {
	if id == toId {
		err = customError.New(5, "无法删除自己账号")
		return
	}
	res := orm.Db.Where("id=?", toId).Delete(&Admin{})
	if res.Error != nil {
		err = customError.New(5, res.Error.Error())
	}
	return
}

func LoginByName(username string, password string) (admin Admin, err error) {
	if len(password) == 0 {
		err = errors.New("密码为空")
		return
	}
	admin, err0 := FindByName(username)
	if err0 != nil {
		err = errors.New("此用户不存在")
		return
	}
	if admin.Status == 0 {
		err = errors.New("此用户被禁用")
		return
	}
	if len(password) != 32 {
		password = helper.Md5(password)
	}
	if admin.Password == helper.Md5(password+admin.Salt) {
		return
	} else {
		err = errors.New("用户名和密码不匹配")
		return
	}

}

// 修改密码
func (admin Admin) ChangePasswd(passwd string, newPasswd string) (err error) {
	if admin.Status == 0 {
		err = errors.New("用户被禁用")
		return
	}
	if len(passwd) != 32 {
		passwd = helper.Md5(passwd)
	}
	if len(newPasswd) != 32 {
		newPasswd = helper.Md5(newPasswd)
	}
	if admin.Password == helper.Md5(passwd+admin.Salt) {
		result := orm.Db.Model(&admin).Updates(Admin{
			Password: helper.Md5(newPasswd + admin.Salt),
			Init:     0, UpdateTime: helper.GetUnix(),
		})
		if result.Error != nil {
			err = errors.New("服务器错误")
			return
		}
		return
	} else {
		err = errors.New("用户名或密码错误")
		return
	}
}

// 按email查找用户
func FindByEmail(email string) (admin Admin, err error) {
	result := orm.Db.Where("email = ?", email).First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("not found")
		return
	}
	return
}

// 查找用户名或email是否存在
func FindByNameOrEmail(username string, email string) (count int64) {
	orm.Db.Model(&Admin{}).Where("username = ? OR email = ?", username, email).Count(&count)
	return
}

// 修改管理员权限
func (admin Admin) ChangeAuth(auth string) (err error) {
	err = orm.Db.Model(&admin).Update("auth", auth).Error
	return
}

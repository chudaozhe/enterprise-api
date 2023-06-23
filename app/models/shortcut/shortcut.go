package shortcut

import (
	customError "enterprise-api/app/models/errors"
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"gorm.io/gorm"
)

// 快捷方式
type Shortcut struct {
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Url        string `json:"url"`
	Status     int    `json:"status"`
	Sort       int    `json:"sort"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// 设置表名
func (Shortcut) TableName() string {
	return "cw_shortcut"
}

func FindById(id int) (shortcut Shortcut, err error) {
	result := orm.Db.Model(&shortcut).
		Select("cw_shortcut.*").
		Where("cw_shortcut.id = ?", id).
		First(&shortcut)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(user bool, page int, max int) (shortcuts []*Shortcut, err error) {
	queryDB := orm.Db.Model(&shortcuts).
		Select("cw_shortcut.*").
		Order("cw_shortcut.id desc").
		Offset((page - 1) * max).Limit(max)
	if user {
		queryDB.Where("`cw_shortcut`.`status`=?", 1)
	}
	result := queryDB.Find(&shortcuts)
	if result.Error != nil {
		err = customError.New(4, result.Error.Error())
		return
	}
	return
}

// 添加
func (shortcut Shortcut) CreateShortcut() (id int, err error) {
	shortcut.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&shortcut)
	if result.Error != nil {
		err = result.Error
		return
	}
	return shortcut.Id, nil
}

// 改变显示状态
func (shortcut Shortcut) ChangeState(status int) (err error) {
	if status == 1 || status == 0 {
		err = orm.Db.Model(&shortcut).Update("status", status).Error
		return
	} else {
		err = errors.New("禁止更新")
		return
	}
}

// 更新
func (shortcut Shortcut) UpdateShortcut() (err error) {
	shortcut.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(shortcut).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Shortcut{}).Error
	return
}

package flash

import (
	customError "enterprise-api/app/models/errors"
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"gorm.io/gorm"
)

// 轮播图
type Flash struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Url        string `json:"url"`
	Status     int    `json:"status"`
	Sort       int    `json:"sort"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// 设置表名
func (Flash) TableName() string {
	return "cw_flash"
}

func FindById(id int) (flash Flash, err error) {
	result := orm.Db.Model(&flash).
		Select("cw_flash.*").
		Where("cw_flash.id = ?", id).
		First(&flash)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(user bool, page int, max int) (flashs []*Flash, err error) {
	queryDB := orm.Db.Model(&flashs).
		Select("cw_flash.*").
		Order("cw_flash.id desc").
		Offset((page - 1) * max).Limit(max)
	if user {
		queryDB.Where("`cw_flash`.`status`=?", 1)
	}
	result := queryDB.Find(&flashs)
	if result.Error != nil {
		err = customError.New(4, result.Error.Error())
		return
	}
	return
}

// 添加
func (flash Flash) CreateFlash() (id int, err error) {
	flash.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&flash)
	if result.Error != nil {
		err = result.Error
		return
	}
	return flash.Id, nil
}

// 改变显示状态
func (flash Flash) ChangeState(status int) (err error) {
	if status == 1 || status == 0 {
		err = orm.Db.Model(&flash).Update("status", status).Error
		return
	} else {
		err = errors.New("禁止更新")
		return
	}
}

// 更新
func (flash Flash) UpdateFlash() (err error) {
	flash.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(flash).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Flash{}).Error
	return
}

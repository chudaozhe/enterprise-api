package cases

import (
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// cases
type Case struct {
	Id          int    `json:"id"`
	CategoryId  int    `json:"category_id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Images      string `json:"images"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
	Sort        int    `json:"sort"`
	Url         string `json:"url"`
	Views       int    `json:"views"`
	Status      int    `json:"status"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}
type CaseResult struct {
	Case
	CategoryName string `json:"category_name"`
}

// 设置表名
func (Case) TableName() string {
	return "cw_case"
}
func FindById(id int) (cases Case, err error) {
	result := orm.Db.Model(&cases).
		Select("cw_case.*").
		Where("cw_case.id = ?", id).
		First(&cases)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(categoryId int, keyword string, user bool, page int, max int) (count int64, casess []*CaseResult, err error) {
	queryDB := orm.Db.Model(&casess).
		Select("cw_case.*, cw_category.name category_name").
		Joins("left join cw_category on cw_category.id = cw_case.category_id").
		Order("cw_case.id desc").
		Offset((page - 1) * max).Limit(max)
	if user {
		queryDB.Where("`cw_case`.`status`=?", 1)
	}
	if categoryId > 0 {
		queryDB.Where("`cw_case`.`category_id`=?", categoryId)
	}
	if len(keyword) > 0 {
		queryDB.Where("CONCAT(`cw_case`.`title`, `cw_case`.`content`) LIKE ?", "%"+keyword+"%")
	}
	result := queryDB.Find(&casess).Offset(-1).Limit(-1).Count(&count)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加
func (cases Case) CreateCase() (id int, err error) {
	cases.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&cases)
	if result.Error != nil {
		err = result.Error
		return
	}
	return cases.Id, nil
}

// 改变显示状态
func (cases Case) ChangeState(status int) (err error) {
	if status == 1 || status == 0 {
		err = orm.Db.Model(&cases).Update("status", status).Error
		return
	} else {
		err = errors.New("禁止更新")
		return
	}
}

// 更新
func (cases Case) UpdateCase() (err error) {
	cases.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(cases).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Case{}).Error
	return
}

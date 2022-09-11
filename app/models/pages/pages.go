package pages

import (
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Page struct {
	Id         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Content    string `json:"content,omitempty"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}
type PageResult struct {
	Page
	CategoryName string `json:"category_name"`
}

// 设置表名
func (Page) TableName() string {
	return "cw_pages"
}

func FindById(id int) (page Page, err error) {
	result := orm.Db.Model(&page).
		Select("cw_pages.*").
		Where("cw_pages.id = ?", id).
		First(&page)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(categoryId int, page int, max int) (count int64, pages []*PageResult, err error) {
	queryDB := orm.Db.Model(&pages).
		Select("cw_pages.*, cw_category.name category_name").
		Joins("left join cw_category on cw_category.id = cw_pages.category_id").
		Order("cw_pages.id desc").
		Offset((page - 1) * max).Limit(max)

	if categoryId > 0 {
		queryDB.Where("`cw_page`.`category_id`=?", categoryId)
	}

	result := queryDB.Find(&pages).Offset(-1).Limit(-1).Count(&count)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加
func (page Page) CreatePage() (id int, err error) {
	page.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&page)
	if result.Error != nil {
		err = result.Error
		return
	}
	return page.Id, nil
}

// 更新
func (page Page) UpdatePage() (err error) {
	page.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(page).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Page{}).Error
	return
}

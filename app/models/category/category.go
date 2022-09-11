package category

import (
	"enterprise-api/app/models/files"
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// 分类
type Category struct {
	Id         int    `json:"id"`
	Type       int    `json:"type"`
	ParentId   int    `json:"parent_id"`
	Name       string `json:"name"`
	Sort       int    `json:"sort"`
	Count      int    `json:"count"`
	Memo       string `json:"memo"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

// 设置表名
func (Category) TableName() string {
	return "cw_category"
}

func FindById(id int) (category Category, err error) {
	result := orm.Db.Model(&category).
		Select("cw_category.*").
		Where("cw_category.id = ?", id).
		First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}

func FindByName(name string, type0 int) (category Category, err error) {
	result := orm.Db.Model(&category).
		Select("cw_category.*").
		Where("cw_category.name = ? AND cw_category.type = ?", name, type0).
		First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}

func List(type0 int) (categorys []*Category, err error) {
	result := orm.Db.Model(&categorys).
		Select("cw_category.*").
		Where("cw_category.type = ?", type0).
		Order("cw_category.id desc").
		Find(&categorys)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

func List2() (categorys []*Category, err error) {
	result := orm.Db.Model(&categorys).
		Select("cw_category.*").
		Where("cw_category.type in ?", []int{2, 5}).
		Find(&categorys)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加
func (category Category) CreateCategory() (id int, err error) {
	category.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&category)
	if result.Error != nil {
		err = result.Error
		return
	}
	return category.Id, nil
}

// 更新
func (category Category) UpdateCategory() (err error) {
	category.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(category).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Category{}).Error
	return
}

// ////////////////////////////文件分组//////////////////////////////
// 分组列表
func GetFilegroup() (categorys []*Category, err error) {
	result := orm.Db.Model(&categorys).Select("cw_category.*").Where("cw_category.type = ?", 1).Find(&categorys)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加文件分组
func (category Category) CreateFilegroup() (id int, err error) {
	category.CreateTime = helper.GetUnix()
	category.Type = 1

	result := orm.Db.Create(&category)
	if result.Error != nil {
		err = result.Error
		return
	}
	return category.Id, nil
}

// 获取一条记录
func GetFilegroupById(id int) (category Category, err error) {
	result := orm.Db.Model(&category).Select("cw_category.*").Where("cw_category.id = ?", id).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}

// 修改文件分组
func (category Category) UpdateFilegroup() (err error) {
	category.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(category).Error
	return
}

// 删除分组（文件标记未分类
func DelFilegroup(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Category{}).Error
	err2 := files.Updatefile2(id)
	if err2 != nil {
		return err2
	}
	return
}

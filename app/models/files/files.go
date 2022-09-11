package files

import (
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type File struct {
	Id         int    `json:"id"`
	AdminId    int    `json:"admin_id"`
	CategoryId int    `json:"category_id"`
	Url        string `json:"url"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Size       int    `json:"size"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
}

// 设置表名
func (File) TableName() string {
	return "cw_files"
}

func FindById(id int) (file File, err error) {
	result := orm.Db.Model(&file).
		Select("cw_files.*").
		Where("cw_files.id = ?", id).
		First(&file)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(groupId int, page int, max int) (count int64, files []*File, err error) {
	queryDB := orm.Db.Model(&files).
		Select("cw_files.*").
		Order("cw_files.id desc").
		Offset((page - 1) * max).Limit(max)

	if groupId > 0 {
		queryDB.Where("`cw_file`.`category_id`=?", groupId)
	}

	result := queryDB.Find(&files).Select("").Offset(-1).Limit(-1).Count(&count)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加
func (file File) Createfile() (id int, err error) {
	file.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&file)
	if result.Error != nil {
		err = result.Error
		return
	}
	return file.Id, nil
}

// 修改文件所属分组
func Updatefile(fileIds []int, groupId int) (err error) {
	if len(fileIds) > 0 {
		err = orm.Db.Model(File{}).Where("id in ?", fileIds).Update("category_id", groupId).Error
		return
	} else {
		return errors.New("id为空")
	}
}

// 取消某个分组
func Updatefile2(groupId int) (err error) {
	err = orm.Db.Model(File{}).Where("category_id = ?", groupId).Update("category_id", 0).Error
	return
}

// 删除
func DeleteById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&File{}).Error
	return
}

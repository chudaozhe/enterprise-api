package article

import (
	orm "enterprise-api/core/db"
	"enterprise-api/core/helper"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	Id          int    `json:"id"`
	CategoryId  int    `json:"category_id"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Content     string `json:"content,omitempty"`
	Image       string `json:"image"`
	Images      string `json:"images"`
	Source      string `json:"source"`
	Author      string `json:"author"`
	Sort        int    `json:"sort"`
	Url         string `json:"url"`
	Views       int    `json:"views"`
	Status      int    `json:"status"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}
type ArticleResult struct {
	Article
	CategoryName string `json:"category_name"`
	Type         int    `json:"type"`
}

// 设置表名
func (Article) TableName() string {
	return "cw_article"
}

func FindById(id int) (article ArticleResult, err error) {
	//sql := "SELECT SQL_CALC_FOUND_ROWS `cw_article`.*, `cw_category`.`name` `category_name`, `cw_category`.`type` `type` FROM `cw_article` LEFT JOIN `cw_category` ON `cw_category`.`id` = `cw_article`.`category_id` WHERE `cw_article`.`id` = ?"
	//err = orm.Db.Raw(sql, id).Scan(&article).Error
	result := orm.Db.Model(&article).
		Select("cw_article.*, cw_category.name category_name, type").
		Joins("left join cw_category on cw_category.id = cw_article.category_id").
		Where("cw_article.id = ?", id).
		First(&article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	return
}
func List(category_ids []int, keyword string, user bool, page int, max int) (count int64, articles []*ArticleResult, err error) {
	//var articles []*ArticleResult
	queryDB := orm.Db.Model(&articles).
		Select("cw_article.*, cw_category.name category_name, type").
		Joins("left join cw_category on cw_category.id = cw_article.category_id").
		Order("cw_article.id desc").
		Offset((page - 1) * max).Limit(max)
	if user {
		queryDB.Where("`cw_article`.`status`=?", 1)
	}
	//if category_id > 0 {
	//	queryDB.Where("`cw_article`.`category_id`=?", category_id)
	//}
	if len(category_ids) > 0 {
		queryDB.Where("`cw_article`.`category_id` in ?", category_ids)
	}
	if len(keyword) > 0 {
		queryDB.Where("CONCAT(`cw_article`.`title`, `cw_article`.`content`) LIKE ?", "%"+keyword+"%")
	}
	result := queryDB.Find(&articles).Offset(-1).Limit(-1).Count(&count)
	if result.Error != nil {
		fmt.Println("not found")
		return
	}
	return
}

// 添加
func (article Article) CreateArticle() (id int, err error) {
	article.CreateTime = helper.GetUnix()

	result := orm.Db.Create(&article)
	if result.Error != nil {
		err = result.Error
		return
	}
	return article.Id, nil
}

// 改变显示状态
func (article Article) ChangeArticleState(status int) (err error) {
	if status == 1 || status == 0 {
		err = orm.Db.Model(&article).Update("status", status).Error
		return
	} else {
		err = errors.New("禁止更新")
		return
	}
}

// 更新
func (article Article) UpdateArticle() (err error) {
	article.UpdateTime = helper.GetUnix()
	err = orm.Db.Save(article).Error
	return
}

// 删除
func DeleteArticleById(id int) (err error) {
	err = orm.Db.Where("id=?", id).Delete(&Article{}).Error
	return
}

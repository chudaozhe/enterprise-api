package models

import (
	orm "enterprise-api/core/db"
	"fmt"
	"gorm.io/plugin/soft_delete"
)

type Test struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Memo       string `json:"memo"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	//DeleteTime gorm.DeletedAt `json:"delete_time"` //逻辑删除时将该字段更新为Y-m-d H:i:s
	DeleteTime soft_delete.DeletedAt `json:"delete_time"` // 逻辑删除时将该字段更新为时间戳
}

// 设置表名
func (Test) TableName() string { //默认为结构体名称的复数，即admins
	return "cw_test"
}

// 列表
func List(page, max int, keyword string) (count int64, admins []*Admin, err error) { //, search Admin
	//var tests []*Test
	//fmt.Println(helper.GetUnix())
	//test := Test{Id: 19}
	result := orm.Db.Where("name = ?", "jinzhu1").Delete(&Test{})

	fmt.Println(result.RowsAffected) //返回找到的记录总数
	if result.Error != nil {
		fmt.Println("del error")
		return
	}
	//
	//data, _ := json.Marshal(&test)
	//fmt.Printf("%s\n", data)

	//err = orm.Db.Offset((page - 1) * max).Limit(max).Find(&admins).Offset(-1).Limit(-1).Count(&count).Error
	return
}

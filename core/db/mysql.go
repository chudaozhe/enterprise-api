package db

import (
	"enterprise-api/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	var err error
	Db, err = gorm.Open(mysql.Open(config.GetConfig().Database.DSN), &gorm.Config{
		CreateBatchSize: 1000, //批量创建时，每批的数量
	})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
}

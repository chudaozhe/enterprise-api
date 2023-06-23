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
		panic(fmt.Errorf("mysql connect error: %w", err))
	}

	if Db.Error != nil {
		panic(fmt.Errorf("database error: %w", Db.Error))
	}
}

package main

import (
	_ "embed"
	"enterprise-api/app/config"
	"enterprise-api/app/routes"
	"enterprise-api/core/cache"
	"enterprise-api/core/db"
	"github.com/gin-gonic/gin"
)

//go:embed app/config/env
var mode string //debug, release, test

func main() {
	if mode == "test" {
		gin.SetMode(gin.TestMode)
	} else if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	config.ParseConfig()
	db.Init()
	//初始化redis
	cache.InitRedis()
	//注册路由
	router := routes.InitRouter()
	router.Run(config.GetConfig().AppPort)
}

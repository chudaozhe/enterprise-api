package routes

import (
	"enterprise-api/app/controllers/admin"
	"enterprise-api/app/controllers/user"
	"enterprise-api/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Error(), middlewares.FixParam(), middlewares.Cors(), middlewares.LoggerToFile())
	v1 := router.Group("v1")

	userRouter := v1.Group("user")
	userRouter.GET("/test", user.CreateTest) //
	userRouter.Use(middlewares.Auth("user")) //middlewares.JWTAuth("user")
	{
		// 轮播图
		userRouter.GET("/:user_id/flash", user.ListFlash)

		// 分类
		userRouter.GET("/:user_id/category", user.ListCategory)

		// 快捷方式
		userRouter.GET("/:user_id/shortcut", user.ListShortcut)

		// 案例
		userRouter.GET("/:user_id/category/:category_id/cases", user.ListCase)
		userRouter.GET("/:user_id/cases/:case_id", user.DetailCase)

		// 文章
		userRouter.GET("/:user_id/category/:category_id/article", user.ListArticle)
		userRouter.GET("/:user_id/article/:article_id", user.DetailArticle)

		// 单页
		userRouter.GET("/:user_id/page/:page_id", user.DetailPage)
	}

	adminRouter := v1.Group("admin")
	adminRouter.POST("/login", admin.Login)          //登录
	adminRouter.POST("/forget", admin.FindPasswd)    //找回密码
	adminRouter.POST("/editor/upload", admin.Upload) //编辑器文件上传
	adminRouter.Use(middlewares.Auth("admin"))       //middlewares.JWTAuth("admin")
	{
		//admin
		adminRouter.GET("/:admin_id", admin.Detail)                //管理员详情
		adminRouter.PUT("/:admin_id", admin.Change)                //修改账号信息
		adminRouter.PUT("/:admin_id/password", admin.ChangePasswd) //修改密码
		adminRouter.PUT("/:admin_id/avatar", admin.Avatar)         //上传头像
		adminRouter.DELETE("/:admin_id/logout", admin.Logout)      //退出

		//other admin
		adminRouter.POST("/:admin_id/admin", admin.CreateAdmin)                 //添加管理员
		adminRouter.GET("/:admin_id/admin", admin.ListAdmin)                    //管理员列表
		adminRouter.GET("/:admin_id/admin/:to_aid", admin.DetailAdmin)          //管理员详情
		adminRouter.PUT("/:admin_id/admin/:to_aid", admin.ChangeAdmin)          //修改账号信息
		adminRouter.DELETE("/:admin_id/admin/:to_aid", admin.DeleteAdmin)       //删除账号
		adminRouter.PUT("/:admin_id/admin/:to_aid/reset", admin.ResetPasswd)    //重置密码
		adminRouter.PUT("/:admin_id/admin/:to_aid/auth", admin.ChangeAuth)      //设置权限
		adminRouter.PUT("/:admin_id/admin/:to_aid/disable", admin.DisableAdmin) //禁止登陆
		adminRouter.PUT("/:admin_id/admin/:to_aid/enable", admin.EnableAdmin)   //解禁

		//filegroup
		adminRouter.GET("/:admin_id/filegroup", admin.ListGroup)                //
		adminRouter.POST("/:admin_id/filegroup", admin.CreateGroup)             //
		adminRouter.GET("/:admin_id/filegroup/:group_id", admin.DetailGroup)    //
		adminRouter.PUT("/:admin_id/filegroup/:group_id", admin.ChangeGroup)    //
		adminRouter.DELETE("/:admin_id/filegroup/:group_id", admin.DeleteGroup) //

		//file
		adminRouter.GET("/:admin_id/filegroup/:group_id/file", admin.ListFile)    //
		adminRouter.POST("/:admin_id/filegroup/:group_id/file", admin.CreateFile) //上传文件
		adminRouter.GET("/:admin_id/file/:file_id", admin.DetailFile)             //
		adminRouter.PUT("/:admin_id/filegroup/:group_id/file", admin.ChangeFile)  //修改文件所属组
		adminRouter.DELETE("/:admin_id/file/:file_id", admin.DeleteFile)          //

		//category
		adminRouter.GET("/:admin_id/category", admin.ListCategory)                   //
		adminRouter.GET("/:admin_id/category2", admin.List2Category)                 //
		adminRouter.POST("/:admin_id/category", admin.CreateCategory)                //
		adminRouter.GET("/:admin_id/category/:category_id", admin.DetailCategory)    //
		adminRouter.PUT("/:admin_id/category/:category_id", admin.ChangeCategory)    //
		adminRouter.DELETE("/:admin_id/category/:category_id", admin.DeleteCategory) //

		//article
		adminRouter.GET("/:admin_id/article", admin.List2Article)                                    //
		adminRouter.GET("/:admin_id/category/:category_id/article", admin.ListArticle)               //
		adminRouter.POST("/:admin_id/category/:category_id/article", admin.CreateArticle)            //
		adminRouter.GET("/:admin_id/article/:article_id", admin.DetailArticle)                       //
		adminRouter.PUT("/:admin_id/category/:category_id/article/:article_id", admin.ChangeArticle) //
		adminRouter.PUT("/:admin_id/article/:article_id/display", admin.DisplayArticle)              //
		adminRouter.PUT("/:admin_id/article/:article_id/hidden", admin.HiddenArticle)                //
		adminRouter.DELETE("/:admin_id/article/:article_id", admin.DeleteArticle)                    //

		//cases
		adminRouter.GET("/:admin_id/category/:category_id/case", admin.ListCase)            //
		adminRouter.POST("/:admin_id/category/:category_id/case", admin.CreateCase)         //
		adminRouter.GET("/:admin_id/case/:case_id", admin.DetailCase)                       //
		adminRouter.PUT("/:admin_id/category/:category_id/case/:case_id", admin.ChangeCase) //
		adminRouter.PUT("/:admin_id/case/:case_id/display", admin.DisplayCase)              //
		adminRouter.PUT("/:admin_id/case/:case_id/hidden", admin.HiddenCase)                //
		adminRouter.DELETE("/:admin_id/case/:case_id", admin.DeleteCase)                    //

		//page
		adminRouter.GET("/:admin_id/category/:category_id/page", admin.ListPage)            //
		adminRouter.POST("/:admin_id/category/:category_id/page", admin.CreatePage)         //
		adminRouter.GET("/:admin_id/page/:page_id", admin.DetailPage)                       //
		adminRouter.PUT("/:admin_id/category/:category_id/page/:page_id", admin.ChangePage) //
		adminRouter.DELETE("/:admin_id/page/:page_id", admin.DeletePage)                    //

		//flash
		adminRouter.GET("/:admin_id/flash", admin.ListFlash)                      //
		adminRouter.POST("/:admin_id/flash", admin.CreateFlash)                   //
		adminRouter.GET("/:admin_id/flash/:flash_id", admin.DetailFlash)          //
		adminRouter.PUT("/:admin_id/flash/:flash_id", admin.ChangeFlash)          //
		adminRouter.PUT("/:admin_id/flash/:flash_id/display", admin.DisplayFlash) //
		adminRouter.PUT("/:admin_id/flash/:flash_id/hidden", admin.HiddenFlash)   //
		adminRouter.DELETE("/:admin_id/flash/:flash_id", admin.DeleteFlash)       //

		//shortcut
		adminRouter.GET("/:admin_id/shortcut", admin.ListShortcut)                         //
		adminRouter.POST("/:admin_id/shortcut", admin.CreateShortcut)                      //
		adminRouter.GET("/:admin_id/shortcut/:shortcut_id", admin.DetailShortcut)          //
		adminRouter.PUT("/:admin_id/shortcut/:shortcut_id", admin.ChangeShortcut)          //
		adminRouter.PUT("/:admin_id/shortcut/:shortcut_id/display", admin.DisplayShortcut) //
		adminRouter.PUT("/:admin_id/shortcut/:shortcut_id/hidden", admin.HiddenShortcut)   //
		adminRouter.DELETE("/:admin_id/shortcut/:shortcut_id", admin.DeleteShortcut)       //

	}
	return router
}

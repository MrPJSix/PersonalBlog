package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//User模块的路由接口
		auth.POST("user", v1.AddUser)

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category", v1.AddCategory)

		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.GET("users", v1.GetUsers)

		// 登录控制模块
		router.POST("login", v1.Login)

		// 分类信息模块
		router.GET("categorys", v1.GetCategorys)

		// 文章信息模块
		router.GET("article/info/:id", v1.GetArticle)
		router.GET("articles", v1.GetArticles)
		router.GET("articles/list/:cid", v1.GetCateArt)
	}
	r.Run(utils.HttpPort)
}

package routers

import (
	v1 "Blog/api/v1"
	"Blog/middleware"
	"Blog/utils"

	_ "net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Loggoer())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
		auth.Use(middleware.JwtToken())
	{
		//user模块路由接口
		
		
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		//文章模块的路由接口
		auth.POST("article/add", v1.AddArt)
		
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		//上传文件
		auth.POST("upload",v1.Upload1)

	}
	routerv1:=r.Group("api/v1")
	{
	//user模块路由接口
	routerv1.GET("users", v1.GetUsers)
	routerv1.POST("user/add", v1.AddUser)
	//分类模块的路由接口
	routerv1.GET("category", v1.GetCategorys)
	//文章模块的路由接口
	routerv1.GET("article", v1.GetArts)
	routerv1.GET("article/list/:id", v1.GetCateArt)
	routerv1.GET("article/info/:id", v1.GetArtInfo)
	routerv1.POST("login",v1.Login)
	}
	r.Run(utils.HttpPort)
	
}
// 了解一下gin validator
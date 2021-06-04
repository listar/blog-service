package routes

import (
	"github.com/gin-gonic/gin"
	"qqfav-service/controllers"
	"qqfav-service/filters/auth"
)

func RegisterApiRouter(router *gin.Engine) {
	//apiRouter := router.Group("api")
	//{
	//	apiRouter.GET("/test/index", controllers.IndexApi)
	//}

	api := router.Group("/api")
	api.GET("/index", controllers.IndexApi)
	api.POST("/login/account", controllers.Account)
	//api.GET("/cookie/set/:userid", controllers.CookieSetExample)
	// qqfav
	api.POST("/article/list", controllers.ArticleList)
	api.POST("/article/detail", controllers.ArticleDetail)

	api.POST("/poetry/list", controllers.PoetryList)
	api.POST("/poetry/detail", controllers.PoetryDetail)

	api.POST("/saying/list", controllers.SayingList)
	api.POST("/saying/detail", controllers.SayingDetail)


	//// cookie auth middleware
	//api.Use(auth.Middleware(auth.CookieAuthDriverKey))
	//{
	//	api.GET("/orm", controllers.OrmExample)
	//	api.GET("/store", controllers.StoreExample)
	//	api.GET("/db", controllers.DBExample)
	//	api.GET("/cookie/get", controllers.CookieGetExample)
	//}

	jwtApi := router.Group("/api")
	//jwtApi.GET("/jwt/set/:userid", controllers.JwtSetExample)

	// jwt auth middleware
	jwtApi.Use(auth.Middleware(auth.JwtAuthDriverKey))
	{
		jwtApi.POST("/article/action", controllers.ArticleAction)
		jwtApi.POST("/article/del", controllers.ArticleDel)

		jwtApi.POST("/poetry/action", controllers.PoetryAction)
		jwtApi.POST("/poetry/del", controllers.PoetryDel)

		jwtApi.POST("/saying/action", controllers.SayingAction)
		jwtApi.POST("/saying/del", controllers.SayingDel)

		//jwtApi.GET("/jwt/get", controllers.JwtGetExample)
	}


}

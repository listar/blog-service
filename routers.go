package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"qqfav-service/config"
	"qqfav-service/filters"
	"qqfav-service/filters/auth"
	routeRegister "qqfav-service/routes"
	"net/http"
	//proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob(config.GetEnv().TemplatePath + "/*") // html模板

	if config.GetEnv().Debug {
		pprof.Register(router) // 性能分析工具
	}

	//router.Use(Cors()) //跨域
	router.Use(gin.Logger())

	router.Use(handleErrors())            // 错误处理
	router.Use(filters.RegisterSession()) // 全局session
	router.Use(filters.RegisterCache())   // 全局cache

	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth")) // 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))    // 全局auth jwt

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
	})

	routeRegister.RegisterApiRouter(router)

	// ReverseProxy
	//router.Use(proxy.ReverseProxy(map[string] string {
	//	"http://www.qqfav.com:10070" : "localhost:8001",
	//	"http://localhost:10070":"localhost:8002",
	//}))

	return router
}


func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}


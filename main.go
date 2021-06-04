package main

import (
	"github.com/gin-gonic/gin"
	_ "qqfav-service/modules/log" // 日志
	// _ "qqfav-service/modules/schedule" // 定时任务
	"runtime"
	"qqfav-service/config"
	"qqfav-service/modules/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter()

	server.Run(router)
}

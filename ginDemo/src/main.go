package main

import (
	"fmt"
	"ginDemo/src/config"
	"ginDemo/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	fmt.Println("started on " + config.PORT)
	engine.Run(config.PORT)
}
package main

import (
	"github.com/gin-gonic/gin"
	"go-reggie/api/route"
	"go-reggie/bootstrap"
	"time"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	// 创建 Gin 引擎
	r := gin.Default()

	// 允许控制台带颜色
	gin.ForceConsoleColor()

	// 定义路由
	route.Setup(env, timeout, r)

	// 启动 HTTP 服务器
	r.Run(env.ServerAddress)
}

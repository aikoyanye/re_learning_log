package main

import (
	"./handler"
	"./tool"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(handler.Cors())
	router.Use(handler.CheckLoginStatus())
	router.Static("/static", "./static")
	router.POST("/user/signup", handler.SignUpHandler)
	router.POST("/user/login", handler.LoginHandler)
	router.POST("/user/logout", handler.LogoutHandler)

	tool.CheckError(router.Run(":8001"), "开启服务失败")
}

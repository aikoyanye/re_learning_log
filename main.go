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
	router.Static("/resource", "./static")
	router.POST("/user/signup", handler.SignUpHandler)
	router.POST("/user/login", handler.LoginHandler)
	router.POST("/user/logout", handler.LogoutHandler)
	router.POST("/user/changeInfo", handler.ChangeUserInfoHJandler)
	router.POST("/home", handler.HomeHandler)
	router.POST("/home/notice", handler.AddNoticeHandler)
	router.POST("/home/ulist", handler.AddUListHandler)
	router.POST("/home/bg", handler.UploadBgImgHandler)
	router.POST("/banip", handler.AddBanIpHandler)
	tool.CheckError(router.Run(":8001"), "开启服务失败")
}

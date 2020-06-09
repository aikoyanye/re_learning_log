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
	router.Use(handler.ColReIp())
	router.Use(handler.BanIp())
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
	router.POST("/title", handler.AllTitleHandler)
	router.POST("/contents", handler.AllContentHandler)
	router.POST("/content", handler.GetContentHandler)
	router.POST("/content/edit", handler.EditContentHandler)
	router.POST("/content/del", handler.DelContentHandler)
	router.POST("/comment/add", handler.AddCommentHandler)
	router.POST("/title/Alltitles", handler.AllTitleWhenAddContentHandler)
	router.POST("/title/add", handler.AddTitleHandler)
	router.POST("/content/add", handler.AddContent)
	router.POST("/title/edit", handler.EditTitleHandler)
	router.POST("/title/del", handler.DelTitleHandler)
	router.POST("/content/uploadPic", handler.UploadContentPicHandler)

	tool.SetBanIps()

	tool.CheckError(router.Run(":8001"), "开启服务失败")
}

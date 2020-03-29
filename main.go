package main

import (
	"./handler"
	"./tool"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main(){
	router := gin.Default()
	router.Static("/static", "./static")
	//router.Static("/view", "./view")
	//router.LoadHTMLGlob("view/*")
	i := func(c *gin.Context) {
		data,_ := ioutil.ReadFile("./static/index.html")
		c.Data(200,"text/html",data)
	}
	router.GET("", i)
	router.POST("/user/signup", handler.SignUpHandler)
	router.POST("/user/login", handler.LoginHandler)

	tool.CheckError(router.Run(":8001"), "开启服务失败")
}
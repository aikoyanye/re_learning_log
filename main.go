package main

import (
	"./handler"
	"./tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}


func main(){
	router := gin.Default()
	router.Use(Cors())
	router.Static("/static", "./static")
	//router.Static("/view", "./view")
	//router.LoadHTMLGlob("view/*")
	//i := func(c *gin.Context) {
	//	data,_ := ioutil.ReadFile("./static/index.html")
	//	c.Data(200,"text/html",data)
	//}
	//router.GET("", i)
	router.POST("/user/signup", handler.SignUpHandler)
	router.POST("/user/login", handler.LoginHandler)

	tool.CheckError(router.Run(":8001"), "开启服务失败")
}

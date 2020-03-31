package main

import (
	"./handler"
	"./tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	router.Use(Cors())
	router.Static("/static", "./static")
	router.POST("/user/signup", handler.SignUpHandler)
	router.POST("/user/login", handler.LoginHandler)

	tool.CheckError(router.Run(":8001"), "开启服务失败")
}

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8888")
		c.Header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

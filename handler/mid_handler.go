package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 检查登录状态是否过期
func CheckLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Request.Cookie("Username")
		tool.CheckError(err, "中间件：获取Cookie失败")
		if len(username.Value) <= 9 {
			c.JSON(http.StatusForbidden, gin.H{"msg": "登录状态失效"})
		} else {
			c.Next()
		}
	}
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

package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 检查登录状态是否过期
func CheckLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 敏感操作才需要判断登录状态
		blackList := map[string]bool{
			"/user/logout": true,
			"/user/changeInfo": true,
			"/home/notice": true,
			"/home/ulist": true,
			"/banip": true,
		}
		if blackList[c.Request.URL.Path]{
			username, err := c.Request.Cookie("Username")
			if !tool.CheckError(err, "中间件：获取Cookie失败") {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "登录状态失效"})
				return
			}
			if len(username.Value) <= 0{
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "登录状态失效"})
				return
			}
			c.Next()
		}
	}
}

// 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "http://120.77.153.248:8888")
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

// 收集访问IP
func ColReIp() gin.HandlerFunc{
	return func(c *gin.Context) {
		if !tool.AddReIp(c.ClientIP()){

		}
		c.Next()
	}
}

// 指定ip无法访问
func BanIp() gin.HandlerFunc{
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		for _, ip := range tool.BanIps{
			if strings.Contains(clientIp, ip){
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "bye"})
			}
		}
		c.Next()
	}
}
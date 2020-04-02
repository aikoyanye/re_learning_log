package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var signupParam struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
		Username string `json:"Username"`
	}
	tool.CheckError(c.Bind(&signupParam), "注册数据有误")
	tool.SignUp(signupParam.Username, signupParam.Password, signupParam.Email)
	c.JSON(http.StatusOK, nil)
}

func LoginHandler(c *gin.Context) {
	var loginParam struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	tool.CheckError(c.Bind(&loginParam), "登录数据有误")
	result := tool.Login(loginParam.Email, loginParam.Password)
	if result.Id == "" {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.SetCookie("Username", result.Username, 86400, "/", "127.0.0.1", false, false)
		c.SetCookie("Id", result.Id, 86400, "/", "127.0.0.1", false, false)
		c.JSON(http.StatusOK, result)
	}
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("Username", "", 86400, "/", "127.0.0.1", false, false)
	c.SetCookie("Id", "", 86400, "/", "127.0.0.1", false, false)
	c.JSON(http.StatusOK, nil)
}

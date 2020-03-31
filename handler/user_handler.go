package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context){
	tool.SignUp(c.PostForm("Username"), c.PostForm("Password"), c.PostForm("Email"))
	c.JSON(http.StatusOK, nil)
}

func LoginHandler(c *gin.Context){
	var loginParam struct{
		Email 		string	`json:"Email"`
		Password 	string	`json:"Password"`
	}
	tool.CheckError(c.Bind(&loginParam), "登录数据有误")
	result := tool.Login(loginParam.Email, loginParam.Password)
	if result.Id == ""{
		c.JSON(http.StatusBadRequest, nil)
	}else{
		c.JSON(http.StatusOK, result)
	}
}
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
	result := tool.Login(c.PostForm("Email"), c.PostForm("Password"))
	if result.Id == ""{
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, result)
}
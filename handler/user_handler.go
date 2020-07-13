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
	if !tool.CheckError(c.Bind(&signupParam), "注册数据有误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "注册失败"})
		return
	}
	if !tool.SignUp(signupParam.Username, signupParam.Password, signupParam.Email){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func LoginHandler(c *gin.Context) {
	var loginParam struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	if !tool.CheckError(c.Bind(&loginParam), "登录数据有误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
		return
	}
	err, result := tool.Login(loginParam.Email, loginParam.Password)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "登录失败"})
		return
	}
	if result.Id == "" {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.SetCookie("Username", result.Username, 86400, "/", "120.77.153.248", false, false)
		c.SetCookie("Id", result.Id, 86400, "/", "120.77.153.248", false, false)
		c.SetCookie("Type", result.Type, 86400, "/", "120.77.153.248", false, false)
		c.JSON(http.StatusOK, result)
	}
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("Username", "", 86400, "/", "120.77.153.248", false, false)
	c.SetCookie("Id", "", 86400, "/", "120.77.153.248", false, false)
	c.JSON(http.StatusOK, nil)
}

func ChangeUserInfoHJandler(c *gin.Context){
	var userInfoParam struct {
		Id    string `json:"Id"`
		Username string `json:"Username"`
		OPassword string `json:"OPassword"`
		NPassword string `json:"NPassword"`
	}
	if !tool.CheckError(c.Bind(&userInfoParam), "修改用户信息数据有误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "修改信息失败"})
		return
	}
	err, result := tool.ChangeUserInfo(userInfoParam.Id, userInfoParam.Username, userInfoParam.OPassword, userInfoParam.NPassword)
	if !err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "修改信息失败"})
		return
	}
	if result >= 1{
		c.SetCookie("Username", userInfoParam.Username, 86400, "/", "127.0.0.1", false, false)
		c.JSON(http.StatusOK, nil)
	}else{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "修改信息失败，可能是旧密码输入错误"})
		return
	}
}
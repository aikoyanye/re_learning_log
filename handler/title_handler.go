package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllTitleHandler(c *gin.Context){
	var user struct{
		Id string `json:"Id"`
	}
	if !tool.CheckError(c.Bind(&user), "获取title错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取title错误"})
	}
	c.JSON(http.StatusOK, tool.AllTitle(user.Id))
}
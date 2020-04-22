package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllContentHandler(c *gin.Context){
	var info struct{
		UserId string `json:"UserId"`
		TitleId string `json:"TitleId"`
	}
	if !tool.CheckError(c.Bind(&info), "获取AllContent错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取AllContent错误"})
	}
	c.JSON(http.StatusOK, tool.AllContent(info.TitleId, info.UserId))
}

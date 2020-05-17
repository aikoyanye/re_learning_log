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

func GetContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&content), "获取Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取文章错误"})
	}
	c.JSON(http.StatusOK, tool.GetContentById(content.ContentId))
}

func EditContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
		Value string `json:"Value"`
	}
	if !tool.CheckError(c.Bind(&content), "更新Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更新文章错误"})
	}
	c.JSON(http.StatusOK, tool.EditContentById(content.ContentId, content.Value))
}

func DelContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&content), "隐藏Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文章错误"})
	}
	if tool.DelContentById(content.ContentId){
		c.JSON(http.StatusOK, nil)
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "删除文章出错"})
	}
}
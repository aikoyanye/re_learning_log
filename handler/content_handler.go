package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllContentHandler(c *gin.Context){
	var info struct{
		UserId 		string `json:"UserId"`
		TitleId 	string `json:"TitleId"`
		ContentId 	string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&info), "获取AllContent错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取AllContent错误"})
		return
	}
	c.JSON(http.StatusOK, tool.AllContent(info.TitleId, info.UserId))
}

func GetContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&content), "获取Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取文章错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"content": tool.GetContentById(content.ContentId),
							"comments": tool.AllCommentByCId(content.ContentId)})
}

func EditContentHandler(c *gin.Context){
	var content struct{
		ContentId 	string `json:"ContentId"`
		Value 		string `json:"Value"`
		Head 		string `json:"Head"`
	}
	if !tool.CheckError(c.Bind(&content), "更新Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更新文章错误"})
		return
	}
	c.JSON(http.StatusOK, tool.EditContentById(content.ContentId, content.Value, content.Head))
}

func DelContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&content), "隐藏Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文章错误"})
		return
	}
	if tool.DelContentById(content.ContentId){
		c.JSON(http.StatusOK, nil)
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "删除文章出错"})
	}
}

func AddContent(c *gin.Context){
	var content struct{
		Content 	string `json:"Content"`
		Head 		string `json:"Head"`
		TitleId 	string `json:"TitleId"`
		Hidden 		string `json:"Hidden"`
	}
	if !tool.CheckError(c.Bind(&content), "隐藏Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文章错误"})
		return
	}
	result := tool.AddContent(content.Head, content.Content, content.TitleId, content.Hidden)
	if result != "-1"{
		c.JSON(http.StatusOK, gin.H{"ContentId": result})
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "添加文章出错"})
	}
}
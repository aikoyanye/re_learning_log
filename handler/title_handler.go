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
		return
	}
	err, results := tool.AllTitle(user.Id)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取title错误"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func AllTitleWhenAddContentHandler(c *gin.Context){
	var user struct{
		UserId string `json:"UserId"`
	}
	if !tool.CheckError(c.Bind(&user), "获取title错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取title错误"})
		return
	}
	err, result := tool.AllTitleWhenAddContent(user.UserId)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取title错误"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func AddTitleHandler(c *gin.Context){
	var title struct{
		Title string `json:"Title"`
		UserId string `json:"UserId"`
		Hidden string `json:"Hidden"`
	}
	if !tool.CheckError(c.Bind(&title), "添加title错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加title错误"})
		return
	}
	if !tool.AddTitle(title.Title, title.Hidden, title.UserId){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加title错误"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func EditTitleHandler(c *gin.Context){
	var title struct{
		Title string `json:"Title"`
		TitleId string `json:"TitleId"`
	}
	if !tool.CheckError(c.Bind(&title), "编辑title错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "编辑title错误"})
		return
	}
	if !tool.EditTitle(title.TitleId, title.Title){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "编辑title错误"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func DelTitleHandler(c *gin.Context){
	var title struct{
		TitleId string `json:"TitleId"`
	}
	if !tool.CheckError(c.Bind(&title), "删除title错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除title错误"})
		return
	}
	if !tool.DelTitle(title.TitleId){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除title错误"})
		return
	}
	c.JSON(http.StatusOK, nil)
}
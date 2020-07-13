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
	err, results := tool.AllContent(info.TitleId, info.UserId)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取AllContent错误"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetContentHandler(c *gin.Context){
	var content struct{
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&content), "获取Content错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取文章错误"})
		return
	}
	err, comments := tool.AllCommentByCId(content.ContentId)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取文章评论错误"})
		return
	}
	err, contents := tool.GetContentById(content.ContentId)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取文章评论错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"content": contents, "comments": comments})
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
	err, result := tool.EditContentById(content.ContentId, content.Value, content.Head)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更新文章错误"})
		return
	}
	c.JSON(http.StatusOK, result)
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

func UploadContentPicHandler(c *gin.Context){
	file, err := c.FormFile("File")
	username := c.PostForm("Username")
	if !tool.CheckError(err, "上传content pic出错_"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "上传图片失败_"})
		return
	}
	pic := "ContentPic/" + username + "_" + tool.Now_() + ".png"
	if !tool.CheckError(c.SaveUploadedFile(file, "static/" + pic), "上传content pic出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "上传图片失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pic": "http://120.77.153.248:8001/resource/" + pic})
}
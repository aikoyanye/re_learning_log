package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCommentHandler(c *gin.Context){
	var info struct{
		Email string `json:"Email"`
		Comment string `json:"Comment"`
		ContentId string `json:"ContentId"`
	}
	if !tool.CheckError(c.Bind(&info), "更新comment错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加评论错误"})
		return
	}
	err, result := tool.AddComment(info.Email, info.Comment, info.ContentId)
	if !err{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加评论错误"})
		return
	}
	c.JSON(http.StatusOK, result)
}

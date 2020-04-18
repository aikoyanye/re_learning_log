package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context){
	notice := tool.GetReleaseNotice()
	ulist := tool.GetUpdateList()
	if notice.Content == "" || len(ulist) == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "获取主页信息失败"})
	}
	c.JSON(http.StatusOK, gin.H{"Notice": notice, "UList": ulist})
}

func AddNoticeHandler(c *gin.Context){
	var notice struct{
		Content string `json:"Content"`
	}
	if !tool.CheckError(c.Bind(&notice), "添加通知错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加通知错误"})
	}
	tool.PostNotice(notice.Content)
	c.JSON(http.StatusOK, nil)
}

func AddUListHandler(c *gin.Context){
	var ulist struct{
		Content string `json:"Content"`
		Version string `json:"Version"`
	}
	if !tool.CheckError(c.Bind(&ulist), "添加更新公布错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加更新公布错误"})
	}
	tool.PostUList(ulist.Version, ulist.Content)
	result := tool.GetUpdateList()
	c.JSON(http.StatusOK, gin.H{"ulist": result})
}

func UploadBgImgHandler(c *gin.Context){
	file, err := c.FormFile("BgImg")
	if tool.CheckError(err, "上传首页png出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更换首页图片失败"})
	}
	if tool.CheckError(c.SaveUploadedFile(file, "static/bg.png"), "保存首页png出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更换首页图片失败"})
	}
	c.JSON(http.StatusOK, nil)
}

func AddBanIpHandler(c *gin.Context){
	var banIp struct{
		BanIp string `json:"BanIp"`
	}
	if !tool.CheckError(c.Bind(&banIp), "添加ban ip错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加ban ip错误"})
	}
	tool.AddBanIp(banIp.BanIp)
	c.JSON(http.StatusOK, nil)
}
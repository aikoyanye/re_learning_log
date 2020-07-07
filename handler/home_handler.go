package handler

import (
	"../tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context){
	notice := tool.GetReleaseNotice()
	ulist := tool.GetUpdateList()
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
	if !tool.CheckError(err, "上传首页png出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更换首页图片失败"})
		return
	}
	if !tool.CheckError(c.SaveUploadedFile(file, "static/bg.png"), "保存首页png出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "更换首页图片失败"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func AddBanIpHandler(c *gin.Context){
	var banIp struct{
		BanIp string `json:"BanIp"`
	}
	if !tool.CheckError(c.Bind(&banIp), "添加ban ip错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "添加ban ip错误"})
		return
	}
	tool.AddBanIp(banIp.BanIp)
	c.JSON(http.StatusOK, nil)
}

// 假盘
// 返回指定文件夹的文件目录
func PanHandler(c *gin.Context){
	var currentDir struct{
		CurrentDir string `json:"CurrentDir"`
	}
	if !tool.CheckError(c.Bind(&currentDir), "网盘id错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "网盘id错误"})
		return
	}
	results := tool.PathDirFileList("./static/Pan/" + currentDir.CurrentDir)
	if results == nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "网盘id错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": results})
}

// 假盘
// 指定目录和文件夹名，新建一个文件夹
func CreateDirPanHandler(c *gin.Context){
	var currentDir struct{
		CurrentDir string `json:"CurrentDir"`
		CreateDir string `json:"CreateDir"`
	}
	if !tool.CheckError(c.Bind(&currentDir), "创建文件夹错误"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "创建文件夹错误"})
		return
	}
	if tool.CreateDir("./static/Pan/" + currentDir.CurrentDir + "/" + currentDir.CreateDir){
		results := tool.PathDirFileList("./static/Pan/" + currentDir.CurrentDir)
		if results == nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "创建文件夹失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"files": results})
	}else{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "创建文件夹错误"})
		return
	}
}

// 假盘
// 上传文件
func PanUploadFileHandler(c *gin.Context){
	file, err := c.FormFile("File")
	if !tool.CheckError(err, "上传文件出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "上传文件出错"})
		return
	}
	var currentDir struct{
		CurrentDir string `json:"CurrentDir"`
	}
	if !tool.CheckError(c.Bind(&currentDir), "上传文件出错_"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "上传文件出错_"})
		return
	}
	if !tool.CheckError(c.SaveUploadedFile(file, "./static/Pan/" + currentDir.CurrentDir + file.Filename), "上传文件出错_"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "上传文件出错_"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

//假盘
// 删除选中文件or文件夹
func DeleteSelectionHandler(c *gin.Context){
	var currentDir struct{
		CurrentDir string `json:"CurrentDir"`
		Selection []string `json:"Selection"`
	}
	if !tool.CheckError(c.Bind(&currentDir), "删除文件Or文件夹出错"){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文件Or文件夹出错"})
		return
	}
	if !tool.DeleteSelection(currentDir.CurrentDir, currentDir.Selection){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文件Or文件夹出错"})
		return
	}
	results := tool.PathDirFileList("./static/Pan/" + currentDir.CurrentDir)
	if results == nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "删除文件Or文件夹出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": results})
}
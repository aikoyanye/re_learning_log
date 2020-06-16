package tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var BanIps = []string{}


// 检查err是否为空
// true：err为空，没有错误
// false：err不为空，有错
func CheckError(err error, str string) bool {
	if err != nil{
		fmt.Println(str)
		fmt.Println(err)
		return false
	}
	return true
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Now_() string {
	return time.Now().Format("2006_01_02_15_04_05")
}

func SetBanIps(){
	BanIps = AllBanIp()
}

// 创建文件夹
// true：创建成功
// false：创建失败
func CreateDir(path string) bool {
	path = strings.Replace(path, "//", "/", -1)
	err := os.Mkdir(path, os.ModePerm)
	if !CheckError(err, "创建文件夹失败，文件夹已存在"){
		return false
	}
	return true
}

// 文件夹类型，文件名、文件大小
type F struct{
	FileName string
	FileSize string
}

// 获取指定文件夹下的所有文件
// results：文件列表
func PathDirFileList(path string) []F {
	path = strings.Replace(path, "//", "/", -1)
	results := []F{}
	files, err := ioutil.ReadDir(path)
	if !CheckError(err, "获取" + path + "目录失败"){
		return nil
	}
	for _, file := range files{
		result := F{}
		result.FileName = file.Name()
		if file.IsDir(){
			result.FileSize = "文件夹"
		}else{
			size := strconv.FormatInt(file.Size()/1024,10)
			result.FileSize = size
		}
		results = append(results, result)
	}
	return results
}

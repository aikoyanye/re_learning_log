package tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func CheckError(err error, str string){
	if err != nil{
		fmt.Println(str)
		fmt.Println(err)
	}
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
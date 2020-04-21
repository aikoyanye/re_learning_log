package tool

import "fmt"

type Title struct {
	Id 			string
	Title 		string
	UserId 		string
	Username	string
	Created 	string
	Hidden 		string
}

func AllTitle(id string) []Title {
	sql := ""
	if id == ""{
		sql = "SELECT * FROM title WHERE hidden = \"0\""
	}else{
		sql = "SELECT * FROM title WHERE userId = \"" + id + "\""
	}
	fmt.Println(id)
	fmt.Println(sql)
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "查找Title错误")
	rows, err := stmt.Query()
	CheckError(err, "查找Title错误")
	results := []Title{}
	for rows.Next(){
		result := Title{}
		CheckError(rows.Scan(&result.Id, &result.Title, &result.UserId, &result.Created, &result.Hidden), "导出Title数据错误")
		results = append(results, result)
	}
	return results
}
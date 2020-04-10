package tool

type Notice struct {
	Id 			string
	Content 	string
	Created 	string
}

type UList struct {
	Id 			string
	Version 	string
	Content 	string
	Created 	string
}

func GetReleaseNotice() Notice {
	sql := "SELECT content FROM notice ORDER BY id DESC LIMIT 1"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取公告SQL语句错误")
	rows, err := stmt.Query()
	CheckError(err, "获取公告SQL语句错误")
	result := Notice{}
	if !rows.Next(){
		return result
	}
	CheckError(rows.Scan(&result.Content), "生成Notice错误")
	defer rows.Close()
	defer stmt.Close()
	return result
}

func GetUpdateList() []UList {
	sql := "SELECT version, content, created FROM update_list ORDER BY id DESC"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取更新列表SQL语句错误")
	rows, err := stmt.Query()
	CheckError(err, "获取更新列表SQL语句错误")
	results := []UList{}
	for rows.Next(){
		result := UList{}
		CheckError(rows.Scan(&result.Version, &result.Content, &result.Created), "生成ULists失败")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return results
}

func PostNotice(content string){
	sql := "INSERT INTO notice (content, created) VALUES (?, ?)"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "添加noticeSQL语句错误")
	_, err = stmt.Exec(content, Now())
	CheckError(err, "添加notice错误")
	defer stmt.Close()
}

func PostUList(version, content string){
	sql := "INSERT INTO update_list (version, content, created) VALUES (?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "添加update_listSQL语句错误")
	_, err = stmt.Exec(version, content, Now())
	CheckError(err, "添加update_list错误")
	defer stmt.Close()
}
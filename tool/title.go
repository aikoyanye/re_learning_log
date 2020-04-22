package tool

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
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u WHERE hidden = \"0\" AND t.userId = u.id"
	}else{
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u WHERE t.userId = \"" + id + "\" AND t.userId = u.id"
	}
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "查找Title错误")
	rows, err := stmt.Query()
	CheckError(err, "查找Title错误")
	results := []Title{}
	for rows.Next(){
		result := Title{}
		CheckError(rows.Scan(&result.Id, &result.Title, &result.UserId, &result.Created, &result.Hidden, &result.Username), "导出Title数据错误")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return results
}

func IsTitleIdEqUserId(titleId, userId string) bool {
	sql := "SELECT id FROM title WHERE id = ? AND userId = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "查询titleId与userId匹配错误")
	rows, err := stmt.Query(titleId, userId)
	CheckError(err, "查询titleId与userId匹配错误")
	defer rows.Close()
	defer stmt.Close()
	return rows.Next()
}
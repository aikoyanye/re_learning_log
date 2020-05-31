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
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u " +
			"WHERE hidden = \"0\" AND t.userId = u.id"
	}else{
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u " +
			"WHERE (t.userId = \"" + id + "\" OR hidden = \"0\") AND t.userId = u.id AND hidden != \"2\""
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

func AllTitleWhenAddContent(id string) []Title {
	sql := "SELECT t.id, t.title FROM title t, user u WHERE t.userId = ? AND t.userId = u.id"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "查找Title错误")
	rows, err := stmt.Query(id)
	CheckError(err, "查找Title错误")
	results := []Title{}
	for rows.Next(){
		result := Title{}
		CheckError(rows.Scan(&result.Id, &result.Title), "导出Title数据错误")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return results
}

func AddTitle(title, hidden, userId string){
	sql := "INSERT INTO title (title, userId, created, hidden) VALUES (?, ?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "添加标题错误啦")
	stmt.Exec(title, userId, Now(), hidden)
	defer stmt.Close()
}

func EditTitle(titleId, title string){
	sql := "UPDATE title SET title = ?, created = ? WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "修改标题错误啦")
	stmt.Exec(title, Now(), titleId)
	defer stmt.Close()
}

func DelTitle(titleId string){
	sql := "UPDATE title SET hidden = 2 WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "删除标题错误啦")
	stmt.Exec(titleId)
	defer stmt.Close()
}
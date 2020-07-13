package tool

type Title struct {
	Id 			string
	Title 		string
	UserId 		string
	Username	string
	Created 	string
	Hidden 		string
}

func AllTitle(id string) (bool, []Title) {
	sql := ""
	results := []Title{}
	if id == ""{
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u " +
			"WHERE hidden = \"0\" AND t.userId = u.id"
	}else{
		sql = "SELECT t.id, t.title, t.userId, t.created, t.hidden, u.username FROM title t, user u " +
			"WHERE (t.userId = \"" + id + "\" OR hidden = \"0\") AND t.userId = u.id AND hidden != \"2\""
	}
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "查找Title错误"){
		return false, results
	}
	rows, err := stmt.Query()
	if !CheckError(err, "查找Title错误"){
		return false, results
	}
	for rows.Next(){
		result := Title{}
		if !CheckError(rows.Scan(&result.Id, &result.Title, &result.UserId, &result.Created, &result.Hidden,
			&result.Username), "导出Title数据错误"){
			return false, results
		}
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}

func IsTitleIdEqUserId(titleId, userId string) bool {
	sql := "SELECT id FROM title WHERE id = ? AND userId = ?"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "查询titleId与userId匹配错误"){
		return false
	}
	rows, err := stmt.Query(titleId, userId)
	if !CheckError(err, "查询titleId与userId匹配错误"){
		return false
	}
	defer rows.Close()
	defer stmt.Close()
	return rows.Next()
}

func AllTitleWhenAddContent(id string) (bool, []Title) {
	sql := "SELECT t.id, t.title FROM title t, user u WHERE t.userId = ? AND t.userId = u.id"
	stmt, err := DBObject.Prepare(sql)
	results := []Title{}
	if !CheckError(err, "查找Title错误"){
		return false, results
	}
	rows, err := stmt.Query(id)
	if !CheckError(err, "查找Title错误"){
		return false, results
	}
	for rows.Next(){
		result := Title{}
		if !CheckError(rows.Scan(&result.Id, &result.Title), "导出Title数据错误"){
			return false, results
		}
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}

func AddTitle(title, hidden, userId string) bool {
	sql := "INSERT INTO title (title, userId, created, hidden) VALUES (?, ?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	defer stmt.Close()
	if !CheckError(err, "添加标题错误啦"){
		return false
	}
	row, err := stmt.Exec(title, userId, Now(), hidden)
	if !CheckError(err, "添加标题错误啦"){
		return false
	}
	num, err := row.RowsAffected()
	if !CheckError(err, "添加标题错误啦"){
		return false
	}
	if num != int64(1){
		return false
	}
	return true
}

func EditTitle(titleId, title string) bool {
	sql := "UPDATE title SET title = ?, created = ? WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "修改标题错误啦"){
		return false
	}
	row, err := stmt.Exec(title, Now(), titleId)
	if !CheckError(err, "修改标题错误啦"){
		return false
	}
	num, err := row.RowsAffected()
	if !CheckError(err, "修改标题错误啦"){
		return false
	}
	if num != int64(1){
		return false
	}
	defer stmt.Close()
	return true
}

func DelTitle(titleId string) bool {
	sql := "UPDATE title SET hidden = 2 WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "删除标题错误啦"){
		return false
	}
	row, err := stmt.Exec(titleId)
	if !CheckError(err, "删除标题错误啦"){
		return false
	}
	num, err := row.RowsAffected()
	if !CheckError(err, "删除标题错误啦"){
		return false
	}
	if num != int64(1){
		return false
	}
	defer stmt.Close()
	return true
}
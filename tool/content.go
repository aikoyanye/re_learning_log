package tool

type Content struct {
	Id 			string
	Content 	string
	Head 		string
	TitleId 	string
	TitleName	string
	Created 	string
	Hidden 		string
	Username	string
}

func AllContent(titleId, userId string) []Content {
	sql := ""
	if userId != "" && IsTitleIdEqUserId(titleId, userId){
		sql = "SELECT c.id, c.head, c.created FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden != \"2\" AND t.id = ? AND t.id = c.titleId"
	}else{
		sql = "SELECT c.id, c.head, c.created FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden = \"0\" AND t.id = ? AND t.id = c.titleId"
	}
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取content list 错误")
	rows, err := stmt.Query(titleId)
	CheckError(err, "获取content list 错误")
	results := []Content{}
	for rows.Next(){
		result := Content{}
		CheckError(rows.Scan(&result.Id, &result.Head, &result.Created), "获取content list 错误")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return results
}

func GetContentById(id string) Content {
	sql := "SELECT c.content, u.username, c.created, t.title, t.id FROM content c, title t, user u " +
		"WHERE c.id = ? AND u.id = t.userId AND c.titleId = t.id"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取content 错误")
	rows, err := stmt.Query(id)
	CheckError(err, "获取content 错误")
	result := Content{}
	if rows.Next(){
		CheckError(rows.Scan(&result.Content, &result.Username, &result.Created, &result.TitleName, &result.TitleId), "获取content 错误")
	}else{
		result.Content = "文章获取失败，可能是服务器宕机了"
	}
	defer rows.Close()
	defer stmt.Close()
	return result
}

func EditContentById(id, content string) Content {
	sql := "UPDATE content SET content = ?, created = ? WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "更新文章失败")
	result := Content{}
	result.Created = Now()
	_, err = stmt.Exec(content, result.Created, id)
	CheckError(err, "更新文章失败")
	defer stmt.Close()
	return result
}

func DelContentById(id string) bool {
	sql := "UPDATE content SET hidden=2 WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "隐藏文章失败")
	re, err := stmt.Exec(id)
	CheckError(err, "隐藏文章失败")
	defer stmt.Close()
	aff_row, _ := re.RowsAffected()
	if aff_row == int64(1){
		return true
	}
	return false
}
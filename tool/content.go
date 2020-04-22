package tool

type Content struct {
	Id 			string
	Content 	string
	Head 		string
	TitleId 	string
	Title		string
	Created 	string
	Hidden 		string
}

func AllContent(titleId, userId string) []Content {
	sql := ""
	if userId != "" && IsTitleIdEqUserId(titleId, userId){
		sql = "SELECT c.id, c.content, c.head, c.titleId, c.created, c.hidden, t.title FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden != \"2\" AND t.id = ?"
	}else{
		sql = "SELECT c.id, c.content, c.head, c.titleId, c.created, c.hidden, t.title FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden = \"0\" AND t.id = ?"
	}
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取content list 错误")
	rows, err := stmt.Query(titleId)
	CheckError(err, "获取content list 错误")
	results := []Content{}
	for rows.Next(){
		result := Content{}
		CheckError(rows.Scan(&result.Id, &result.Content, &result.Head, &result.TitleId,
			&result.Created, &result.Hidden, &result.Title), "获取content list 错误")
		results = append(results, result)
	}
	return results
}
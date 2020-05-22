package tool

type Comment struct {
	Id 			string
	Email 		string
	Comment 	string
	ContentId 	string
	Created 	string
	Hidden 		string
}

func AllCommentByCId(contentId string) []Comment {
	sql := "SELECT email, comment, created FROM comment WHERE hidden = 0 AND contentId = ? ORDER BY id DESC"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "获取comment list 错误")
	rows, err := stmt.Query(contentId)
	CheckError(err, "获取comment list 错误")
	results := []Comment{}
	for rows.Next(){
		result := Comment{}
		CheckError(rows.Scan(&result.Email, &result.Comment, &result.Created), "获取comment list 错误")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return results
}

func AddComment(email, comment, contentId string) Comment {
	sql := "INSERT INTO comment (email, comment, contentId, created, hidden) VALUES (?, ?, ?, ?, 0)"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "添加comment 错误")
	result := Comment{}
	result.Created = Now()
	result.Email = email
	result.Comment = comment
	_, err = stmt.Exec(email, comment, contentId, Now())
	CheckError(err, "添加comment 错误")
	defer stmt.Close()
	return result
}
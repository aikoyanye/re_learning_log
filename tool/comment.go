package tool

type Comment struct {
	Id 			string
	Email 		string
	Comment 	string
	ContentId 	string
	Created 	string
	Hidden 		string
}

func AllCommentByCId(contentId string) (bool, []Comment) {
	sql := "SELECT email, comment, created FROM comment WHERE hidden = 0 AND contentId = ? ORDER BY id DESC"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "获取comment list 错误"){
		return false, nil
	}
	rows, err := stmt.Query(contentId)
	if !CheckError(err, "获取comment list 错误"){
		return false, nil
	}
	results := []Comment{}
	for rows.Next(){
		result := Comment{}
		CheckError(rows.Scan(&result.Email, &result.Comment, &result.Created), "获取comment list 错误")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}

func AddComment(email, comment, contentId string) (bool, Comment) {
	sql := "INSERT INTO comment (email, comment, contentId, created, hidden) VALUES (?, ?, ?, ?, 0)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加comment 错误"){
		return false, Comment{}
	}
	result := Comment{}
	result.Created = Now()
	result.Email = email
	result.Comment = comment
	re, err := stmt.Exec(email, comment, contentId, Now())
	if !CheckError(err, "添加comment 错误"){
		return false, Comment{}
	}
	row, err := re.RowsAffected()
	if !CheckError(err, "插入的数据不正确") || row != int64(1){
		return false, Comment{}
	}
	defer stmt.Close()
	return true, result
}
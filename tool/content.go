package tool

import (
	"strconv"
)

type Content struct {
	Id 			string
	Content 	string
	Head 		string
	TitleId 	string
	TitleName	string
	Created 	string
	Hidden 		string
	Username	string
	UserId 		string
}

func AllContent(titleId, userId string) (bool, []Content) {
	var sql string
	if userId != "" && IsTitleIdEqUserId(titleId, userId){
		sql = "SELECT c.id, c.head, c.created FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden != \"2\" AND t.id = ? AND t.id = c.titleId"
	}else{
		sql = "SELECT c.id, c.head, c.created FROM content c, title t, user u " +
			"WHERE u.id = t.userId AND c.hidden = \"0\" AND t.id = ? AND t.id = c.titleId"
	}
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "获取content list 错误"){
		return false, nil
	}
	rows, err := stmt.Query(titleId)
	if !CheckError(err, "获取content list 错误"){
		return false, nil
	}
	results := []Content{}
	for rows.Next(){
		result := Content{}
		if !CheckError(rows.Scan(&result.Id, &result.Head, &result.Created), "获取content list 错误"){
			return false, nil
		}
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}

func GetContentById(id string) (bool, Content) {
	sql := "SELECT c.content, u.username, c.created, t.title, t.id, u.id, c.head FROM content c, title t, user u " +
		"WHERE c.id = ? AND u.id = t.userId AND c.titleId = t.id"
	result := Content{}
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "获取content 错误"){
		return false, result
	}
	rows, err := stmt.Query(id)
	if !CheckError(err, "获取content 错误"){
		return false, result
	}
	if rows.Next(){
		if !CheckError(rows.Scan(&result.Content, &result.Username, &result.Created,
			&result.TitleName, &result.TitleId, &result.UserId, &result.Head), "获取content 错误"){
			return false, result
		}
	}else{
		result.Content = "文章获取失败，可能是服务器宕机了"
	}
	defer rows.Close()
	defer stmt.Close()
	return true, result
}

func EditContentById(id, content, head string) (bool, Content) {
	sql := "UPDATE content SET content = ?, created = ?, head = ? WHERE id = ?"
	result := Content{}
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "更新文章失败"){
		return false, result
	}
	result.Created = Now()
	_, err = stmt.Exec(content, result.Created, head, id)
	if !CheckError(err, "更新文章失败"){
		return false, result
	}
	defer stmt.Close()
	return true, result
}

func DelContentById(id string) bool {
	sql := "UPDATE content SET hidden=2 WHERE id = ?"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "隐藏文章失败"){
		return false
	}
	re, err := stmt.Exec(id)
	if !CheckError(err, "隐藏文章失败"){
		return false
	}
	defer stmt.Close()
	aff_row, _ := re.RowsAffected()
	if aff_row == int64(1){
		return true
	}
	return false
}

func AddContent(head, content, titleId, hidden string) string {
	sql := "INSERT INTO content (content, head, titleId, created, hidden) VALUES (?, ?, ?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加文章失败"){
		return "-1"
	}
	re, err := stmt.Exec(content, head, titleId, Now(), hidden)
	if !CheckError(err, "添加文章失败"){
		return "-1"
	}
	defer stmt.Close()
	result, err := re.LastInsertId()
	if err == nil{
		return strconv.FormatInt(result, 10)
	}
	return "-1"
}
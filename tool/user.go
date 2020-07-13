package tool

import "strconv"

type User struct {
	Id 			string
	Email 		string
	Username 	string
	Password 	string
	Type 		string
	Created 	string
}

func SignUp(username, password, email string) bool {
	password = md5V(password)
	sql := "INSERT INTO user (email, username, password, type, created) VALUES (?, ?, ?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	defer stmt.Close()
	if !CheckError(err, "注册用户的SQL语句可能出现错误"){
		return false
	}
	re, err := stmt.Exec(email, username, password, "normal", Now())
	if !CheckError(err, "注册用户的SQL语句可能出现错误"){
		return false
	}
	id, err := re.LastInsertId()
	if !CheckError(err, "注册用户的SQL语句可能出现错误"){
		return false
	}
	CreateDir("./static/Pan/" + strconv.FormatInt(id,10))
	return true
}

func Login(email, password string) (bool, User) {
	password = md5V(password)
	sql := "SELECT id, username, type FROM user WHERE email = ? AND password = ?"
	stmt, err := DBObject.Prepare(sql)
	result := User{}
	if !CheckError(err, "用户登录的SQL语句可能出现错误"){
		return false, result
	}
	rows, err := stmt.Query(email, password)
	CheckError(err, "用户登录的SQL语句可能出现错误")
	if !rows.Next(){
		return false, result
	}
	if !CheckError(rows.Scan(&result.Id, &result.Username, &result.Type), "生成User错误"){
		return false, result
	}
	defer rows.Close()
	defer stmt.Close()
	return true, result
}

func ChangeUserInfo(id, username, oPass, nPass string) (bool, int64) {
	sql := "UPDATE user SET username = ?, password = ? WHERE id = ? AND Password = ?"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "修改用户数据的SQL语句可能出现错误"){
		return false, -1
	}
	result, err := stmt.Exec(username, md5V(nPass), id, md5V(oPass))
	if !CheckError(err, "修改用户数据的SQL语句可能出现错误"){
		return false, -1
	}
	ra, err :=result.RowsAffected()
	if !CheckError(err, "修改用户数据的SQL语句可能出现错误"){
		return false, -1
	}
	defer stmt.Close()
	return true, ra
}
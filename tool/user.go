package tool

type User struct {
	Id 			string
	Email 		string
	Username 	string
	Password 	string
	Type 		string
	Created 	string
}

func SignUp(username, password, email string){
	password = md5V(password)
	sql := "INSERT INTO user (email, username, password, type, created) VALUES (?, ?, ?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "注册用户的SQL语句可能出现错误")
	_, err = stmt.Exec(email, username, password, "normal", Now())
	CheckError(err, "注册用户过程发生错误")
}

func Login(email, password string) User {
	password = md5V(password)
	sql := "SELECT id, username FROM user WHERE email = \"" + email + "\" AND password = \"" + password + "\""
	rows, err := DBObject.Query(sql)
	CheckError(err, "登录出现错误")
	result := User{}
	if !rows.Next(){
		return result
	}
	CheckError(rows.Scan(&result.Id, &result.Username), "生成User错误")
	defer rows.Close()
	return result
}
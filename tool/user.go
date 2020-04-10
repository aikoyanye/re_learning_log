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
	defer stmt.Close()
}

func Login(email, password string) User {
	password = md5V(password)
	sql := "SELECT id, username FROM user WHERE email = ? AND password = ?"
	stmt, err := DBObject.Prepare(sql)
	CheckError(err, "用户登录的SQL语句可能出现错误")
	rows, err := stmt.Query(email, password)
	CheckError(err, "用户登录的SQL语句可能出现错误")
	result := User{}
	if !rows.Next(){
		return result
	}
	CheckError(rows.Scan(&result.Id, &result.Username), "生成User错误")
	defer rows.Close()
	defer stmt.Close()
	return result
}
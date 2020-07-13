package tool

type Notice struct {
	Id 			string
	Content 	string
	Created 	string
}

type UList struct {
	Id 			string
	Version 	string
	Content 	string
	Created 	string
}

func GetReleaseNotice() (bool, Notice) {
	sql := "SELECT content FROM notice ORDER BY id DESC LIMIT 1"
	stmt, err := DBObject.Prepare(sql)
	result := Notice{}
	if !CheckError(err, "获取公告SQL语句错误"){
		return false, result
	}
	rows, err := stmt.Query()
	if !CheckError(err, "获取公告SQL语句错误"){
		return false, result
	}
	if !rows.Next(){
		return false, result
	}
	CheckError(rows.Scan(&result.Content), "生成Notice错误")
	defer rows.Close()
	defer stmt.Close()
	return true, result
}

func GetUpdateList() (bool, []UList) {
	sql := "SELECT version, content, created FROM update_list ORDER BY id DESC"
	results := []UList{}
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "获取更新列表SQL语句错误"){
		return false, results
	}
	rows, err := stmt.Query()
	if !CheckError(err, "获取更新列表SQL语句错误"){
		return false, results
	}
	for rows.Next(){
		result := UList{}
		CheckError(rows.Scan(&result.Version, &result.Content, &result.Created), "生成ULists失败")
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}

func PostNotice(content string) bool {
	sql := "INSERT INTO notice (content, created) VALUES (?, ?)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加noticeSQL语句错误"){
		return false
	}
	_, err = stmt.Exec(content, Now())
	if !CheckError(err, "添加notice错误"){
		return false
	}
	defer stmt.Close()
	return true
}

func PostUList(version, content string) bool {
	sql := "INSERT INTO update_list (version, content, created) VALUES (?, ?, ?)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加update_listSQL语句错误"){
		return false
	}
	_, err = stmt.Exec(version, content, Now())
	if !CheckError(err, "添加update_listSQL语句错误"){
		return false
	}
	defer stmt.Close()
	return true
}

func AddBanIp(ip string) bool {
	sql := "INSERT INTO ban_ip (ip) VALUES (?)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加ban ip语句错误"){
		return false
	}
	_, err = stmt.Exec(ip)
	if !CheckError(err, "添加ban ip语句错误"){
		return false
	}
	BanIps = append(BanIps, ip)
	defer stmt.Close()
	return true
}

func AddReIp(ip string) bool {
	sql := "INSERT INTO re_ip (ip, created) VALUES (?, ?)"
	stmt, err := DBObject.Prepare(sql)
	if !CheckError(err, "添加re ip语句错误"){
		return false
	}
	_, err = stmt.Exec(ip, Now())
	if !CheckError(err, "添加re ip语句错误"){
		return false
	}
	defer stmt.Close()
	return true
}

func AllBanIp() (bool, []string) {
	sql := "SELECT ip FROM ban_ip"
	stmt, err := DBObject.Prepare(sql)
	results := []string{}
	if !CheckError(err, "获取更新列表SQL语句错误"){
		return false, results
	}
	rows, err := stmt.Query()
	if !CheckError(err, "获取更新列表SQL语句错误"){
		return false, results
	}
	for rows.Next(){
		result := ""
		if !CheckError(rows.Scan(&result), "生成re ip失败"){
			return false, results
		}
		results = append(results, result)
	}
	defer rows.Close()
	defer stmt.Close()
	return true, results
}
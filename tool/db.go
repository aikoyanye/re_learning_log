package tool

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DBObject *sql.DB

// 连接sqlite数据库
func init() {
	db, err := sql.Open("sqlite3", "log.db")
	CheckError(err, "数据库链接出错")
	db.SetMaxOpenConns(1)
	DBObject = db
}

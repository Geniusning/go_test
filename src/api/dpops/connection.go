package dpops

import (
	"database/sql"
)

var (
	dbconn *sql.DB
	err    error
)

func init() {
	dbconn, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

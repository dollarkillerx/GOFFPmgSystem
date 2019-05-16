package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "video_com:RiRrTih5aEbbFzyp@tcp(127.0.0.1:3306)/video_com?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

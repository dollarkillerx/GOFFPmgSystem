package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string,pwd string) error {
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name,pwd,salt) VALUES(?,?)")
	if err!=nil {
		return err
	}
	pwd =
	stmtIns.Exec(loginName,pwd)
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error) {
	stmtOut,err := dbConn.Prepare("SELECT ")
}





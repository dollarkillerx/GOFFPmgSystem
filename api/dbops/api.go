package dbops

import (
	"GOFFPmgSystem/api/tootl"
	"database/sql"
	"log"
)


func AddUserCredential(loginName string,pwd string) error {
	stmt, e := dbConn.Prepare("INSERT INTO `users` (`login_name`,`pwd`) VALUE (?,?)")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		return e
	}
	pwd = tootl.Md5String(pwd)
	_, e = stmt.Exec(loginName, pwd)
	return e
}

func GetUserCredential(loginName string) (string,error)  {
	stmt, e := dbConn.Prepare("SELECT `pwd` FROM `users` WHERE `login_name` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		log.Printf("%s",e)
		return "",e
	}
	var pwd string
	e = stmt.QueryRow(loginName).Scan(&pwd)

	// 对空返回 容错
	if e!=nil || e!=sql.ErrNoRows {
		return pwd,nil
	}
	return pwd,e
}

func DeleteUser(loginName string,pwd string) error {
	stmt, e := dbConn.Prepare("DELETE FROM `users` WHERE `login_name` = ? AND `pwd` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil{
		log.Printf("%s",e)
		return e
	}
	pwd = tootl.Md5String(pwd)
	_, e = stmt.Exec(loginName, pwd)
	return e
}

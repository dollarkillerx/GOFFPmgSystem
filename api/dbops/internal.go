package dbops

import (
	"GOFFPmgSystem/api/defs"
	"database/sql"
	"log"
	"sync"
)

func InserSession(sid string,ttl int64,uname string) error {
	stmt, e := dbConn.Prepare("INSERT INTO sessions (session_id,ttl,login_name) VALUE (?,?,?)")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(sid, ttl, uname)
	return e
}

func RetrieveSession(sid string) (*defs.SimpleSession,error) {
	session := &defs.SimpleSession{}
	stmt, e := dbConn.Prepare("SELECT ttl,login_name FROM sessions WHERE session_id = ?")
	defer stmt.Close()
	if e !=nil {
		return nil,e
	}

	var ttl int64
	var uname string

	e = stmt.QueryRow(sid).Scan(&ttl, &uname)
	if e == nil {
		session.TTL = ttl
		session.Username = uname
	}

	return session,e
}

func RetrieveAllSessions() (*sync.Map,error) {
	i := &sync.Map{}
	stmt, e := dbConn.Prepare("SELECT * FROM sessions")
	defer stmt.Close()
	if e != nil {
		return nil,e
	}

	rows, e := stmt.Query()
	if e != nil && e != sql.ErrNoRows {
		return nil,e
	}

	for rows.Next() {
		var id string
		var ttl int64
		var login_name string
		if e := rows.Scan(&id, &ttl, &login_name);e != nil {
			log.Printf("retrive sessions error: %s", e)
			break
		}
		session := &defs.SimpleSession{Username: login_name, TTL: ttl}

		i.Store(id,session)
	}
	return i,nil
}

func DeleteSession(sid string) error {
	stmt, e := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(sid)
	return e
}


package session

import (
	"GOFFPmgSystem/api/dbops"
	"GOFFPmgSystem/api/defs"
	"GOFFPmgSystem/api/utils"
	"strconv"
	"sync"
)

// sync.Map 线程安全的map
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// 重DB更新cache数据
func LoadSessionsFromDB() {
	sessions, e := dbops.RetrieveAllSessions()
	if e != nil {
		return
	}

	// 拿到数据一条条写入
	sessions.Range(func(key, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionMap.Store(key,session)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUIDSimplicity()
	ctime,_ := strconv.ParseInt(utils.GetCurrentTime(),10,64)
	ttl := ctime + 30*60// 过期时间

	session := &defs.SimpleSession{Username: uname, TTL: ttl}
	sessionMap.Store(id,session)
	dbops.InserSession(id, ttl, uname)

	return id
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// session 是否过期
func IsSessionExpired(sid string) (string,bool) {
	value, ok := sessionMap.Load(sid)
	if ok {
		nowTime,_ := strconv.ParseInt(utils.GetCurrentTime(),10,64)
		if value.(*defs.SimpleSession).TTL < nowTime {
			//delete expired session
			deleteExpiredSession(sid)
			return "",true
		}
		return value.(*defs.SimpleSession).Username,false
	}
	return "",true
}


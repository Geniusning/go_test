package session

import (
	"api/defs"
	"api/dpops"
	"api/utils"
	"fmt"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowIntMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dpops.DeleteSession(sid)
}

func loadSessionFromDB() {
	r, err := dpops.RetrieveAllSessions()
	if err != nil {
		fmt.Println("loadSessionFromDB err=", err)
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

//GenerateNewSessionID 生成新sessionid
func GenerateNewSessionID(un string) string {
	id, _ := utils.NewUUID()
	ct := nowIntMilli()
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dpops.InsertSession(id, ttl, un)

	return id
}

//教研IsSessionExpired 是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowIntMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}

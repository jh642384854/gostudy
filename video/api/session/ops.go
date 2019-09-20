package session

import (
	"dev/video/api/dbops"
	"dev/video/api/defs"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}	
}

func LoadSessionsFromDB()  {
	maps,err := dbops.RetrieveAllSession()
	if err != nil{
		return
	}
	maps.Range(func(key, value interface{}) bool {
		sessionMap.Store(key,value)
		return true
	})
}

func GenerateNewSessionID(username string) string {
	uuidv4,_ := uuid.NewV4()
	session_id := uuidv4.String()
	nowtime := nowInMilli()
	ttl := nowtime+30*60*60;//30分钟
	dbops.InsertSession(ttl,session_id,username)
	sessionObj := &defs.Session{
		SessionID:session_id,
		TTL:ttl,
		LoginName:username,
	}
	sessionMap.Store(session_id,sessionObj)
	return session_id
}

func IsSessionExpired(session_id string) (string,bool) {
	if sm,ok := sessionMap.Load(session_id);ok{
		session := sm.(*defs.Session)//类型强制转换
		ttl := session.TTL
		nowtime := nowInMilli()
		if nowtime > ttl{
			//删除session
			deleteExpiredSession(session_id)
			return "",true
		}else{
			return session.LoginName,false
		}
	}else{
		return "",true
	}
}

func nowInMilli() int64  {
	return time.Now().UnixNano()
}

func deleteExpiredSession(session_id string) {
	sessionMap.Delete(session_id)
	dbops.DeleteSession(session_id)
}
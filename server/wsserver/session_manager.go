package wsserver

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
	"sync"
	"time"
)

const (
	sessionCleanTime = 5              //清理连接时间间隔 单位/s
	sessionOverTime  = 5              //连接心跳阈值 单位/s
	SessionDataKey   = "session_data" //session数据key
)

//保存session的map
var sessionMap sync.Map


// 连接的session数据
type SessionData struct {
	SessionId    string `json:"session_id"`     // 连接的唯一id
	UserId       string `json:"user_id"`        // 连接用户id
	LastPingTime int64  `json:"last_ping_time"` // 最后一次心跳时间
	//...
}

// 设置session数据
func SetSessionData(session *melody.Session, data *SessionData) {
	session.Set(SessionDataKey, data)
}

// 获取session数据
func GetSessionData(session *melody.Session) *SessionData {
	get, exists := session.Get(SessionDataKey)
	if !exists {
		logrus.Error("获取session数据失败:session_data 为 nil")
		_ = session.Close()
		return nil
	}
	res := get.(*SessionData)
	return res
}

//添加客户端的连接类记录
func AddSession(session *melody.Session) {
	data := GetSessionData(session)
	if data == nil {
		logrus.Error("添加客户端的连接类记录失败:session_data 为 nil")
		return
	}
	sid := data.SessionId
	sessionMap.Store(sid, session)
}

//删除客户端的连接类记录
func DeleteSession(sid string) {
	sessionMap.Delete(sid)
}

//获取session by session_id
func GetSessionBySessionId(sid string) *melody.Session {
	if v, ok := sessionMap.Load(sid); ok {
		return v.(*melody.Session)
	}
	return nil
}

//获取session by user_id
//func GetSessionByUserId(userId string) *melody.Session {
//	var res *melody.Session
//	sessionMap.Range(func(k, v interface{}) bool {
//		if session := v.(*melody.Session); session != nil {
//			if data := GetSessionData(session); data != nil {
//				if data.UserId == userId {
//					res = session
//					return false
//				}
//			}
//		}
//		return true
//	})
//	return res
//}

//*****************************************************************************
//清理心跳超过阈值的连接
//*****************************************************************************
func CheckConnSession() {
	ticker := time.NewTicker(time.Second * sessionCleanTime)
	for {
		<-ticker.C
		sessionMap.Range(func(key, value interface{}) bool {
			nowTime := time.Now().Unix()
			session := value.(*melody.Session)
			data := GetSessionData(session)
			if nowTime-data.LastPingTime >= sessionOverTime {
				DeleteSession(data.SessionId)
				_ = session.Close()
			}
			return true
		})
	}
}

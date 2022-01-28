package wsserver

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
)

// 收到客户端发送来得消息
func OnMessageBinary(m *melody.Melody, session *melody.Session, bytes []byte) {
	msgId, data, err := CheckBinaryMsg(bytes)
	if err != nil {
		return
	}
	// 根据msgId 分发任务
	switch msgId {
	case MsgIdPing: // ping pong  // 0
		DealPingPong(session, data)
	case MsgIdReqLogin: // 用户登录 // 100
		ReqLogin(m, session, data)
	default:
		logrus.Error("未知消息:msgId:", msgId)
	}
}

//一个新的连接
func OnConnect(m *melody.Melody, session *melody.Session) {
	sessionData := SessionData{
		SessionId: uuid.New().String(),
	}
	//设置session_id
	SetSessionData(session, &sessionData)
	//保存session
	AddSession(session)
}

//服务器发送二进制消息
func OnSentMessageBinary(session *melody.Session, bytes []byte) {
}

//收到文本消息
func OnMessage(m *melody.Melody, session *melody.Session, bytes []byte) {
}

//关闭连接
func OnClose(m *melody.Melody, session *melody.Session, i int, s string) error {
	return nil
}

//断开连接
func OnDisconnect(m *melody.Melody, session *melody.Session) {
	DeleteSession(GetSessionData(session).SessionId)
}

//处理分发
func Handler(m *melody.Melody) {
	// onMessageBinary
	m.HandleMessageBinary(func(session *melody.Session, bytes []byte) {
		OnMessageBinary(m, session, bytes)
	})

	// onMessage
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		OnMessage(m, s, msg)
	})

	// onClose
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		return OnClose(m, session, i, s)
	})

	// onConn
	m.HandleConnect(func(session *melody.Session) {
		OnConnect(m, session)
	})

	// onDisconnect
	m.HandleDisconnect(func(session *melody.Session) {
		OnDisconnect(m, session)
	})

	// onSentMessageBinary
	m.HandleSentMessageBinary(func(session *melody.Session, bytes []byte) {
		OnSentMessageBinary(session, bytes)
	})
}

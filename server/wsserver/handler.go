package wsserver

import (
	"gameserver/server/wsserver/mproto"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"gopkg.in/olahol/melody.v1"
	"time"
)

// 处理心跳消息
func DealPingPong(session *melody.Session, data []byte) {
	var req mproto.PING
	err := proto.Unmarshal(data, &req)
	if err != nil {
		logrus.Debug("DealPingPong", string(data), err.Error())
		SendErrMsg(session, MsgIdPing, "请求数据异常")
		return
	}
	GetSessionData(session).LastPingTime = time.Now().Unix()
}

// 处理登录消息
func ReqLogin(m *melody.Melody, session *melody.Session, data []byte) {

}

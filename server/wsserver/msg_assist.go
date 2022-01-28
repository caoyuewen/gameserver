package wsserver

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"gameserver/server/wsserver/mproto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/wonderivan/logger"
	"google.golang.org/protobuf/proto"
	"gopkg.in/olahol/melody.v1"
)

// 获取二进制消息的消息id
func GetMsgId(mid []byte) uint16 {
	u := binary.BigEndian.Uint16(mid)
	return u
}

// 封装二进制消息
func PkgMsg(msgId uint16, data []byte) []byte {
	bytes := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(bytes[:2], msgId)
	copy(bytes[2:], data)
	return bytes
}

// 打印消息
func PrintMsg(msgInfo string, data interface{}) {
	bytes, _ := json.Marshal(data)
	logger.Debug(msgInfo, string(bytes))
}

func GetAuthKey(userId string) string {
	key := uuid.New().String() + userId
	authKey := MakeAuthKey(key)
	return authKey
}

func MakeAuthKey(key string) string {
	data := []byte(key)
	sum := md5.Sum(data)
	code := fmt.Sprintf("%x", sum)
	return code
}

// 检查客户端发送的二进制消息
func CheckBinaryMsg(bytes []byte) (uint16, []byte, error) {
	dataLen := len(bytes)
	if dataLen <= 2 {
		logrus.WithFields(logrus.Fields{
			"data":bytes,
		}).Error("收到错误的二进制消息")
		return 0, nil, errors.New("data is nil")
	}
	msgId := GetMsgId(bytes[:2])
	data := bytes[2:]
	return msgId, data, nil
}

// 发送错误消息
func SendErrMsg(session *melody.Session, mId uint16, msg string) {
	var errMsgPush mproto.ErrMsg
	errMsgPush.MsgId = int32(mId)
	errMsgPush.Msg = msg
	bytes, _ := proto.Marshal(&errMsgPush)
	_ = session.WriteBinary(PkgMsg(MsgIdErrMsg, bytes))
}



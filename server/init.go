package server

import (
	"encoding/json"
	"gameserver/common"
	"gameserver/server/httpserver"
	"gameserver/server/wsserver"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/olahol/melody.v1"
	"io/ioutil"
	"net/http"
)

var Conf AppConfig

type AppConfig struct {
	Port  string    `json:"port"`
	Mysql MySqlConf `json:"mysql"`
}

type MySqlConf struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

//初始化服务配置
func init() {
	//设置日志输出器
	SetLogger()
	//加载服务器配置
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("加载服务器配置失败!")
	}
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("解析服务器配置失败!")
	}
	logrus.WithFields(logrus.Fields{
		"config": Conf,
	}).Info("加载服务器配置成功")
}

//启动服务
func StartServer(mode, wsPath string) {
	gin.SetMode(mode)
	r := gin.New()
	m := melody.New()
	r.GET(wsPath, func(c *gin.Context) {
		err := m.HandleRequestWithKeys(c.Writer, c.Request, nil)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"func": "StartServer",
				"err":  err.Error(),
			}).Panic()
		}
	})
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	httpserver.Router(r)
	wsserver.Handler(m)
	go wsserver.CheckConnSession()
	err := r.Run(":" + Conf.Port)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"func": "StartServer",
			"err":  err.Error(),
		}).Panic("启动服务失败")
	}
}

//设置日志输出器
func SetLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: common.LogTimeFormat,
	})
}

package httpserver

import (
	"gameserver/server/httpserver/controller"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/v1/api/test", controller.Test)
	r.GET("/v1/api/test2", controller.Test2)
}

package main

import (
	"gameserver/server"
	"github.com/gin-gonic/gin"
)

func main() {
	server.StartServer(gin.ReleaseMode, "/ws")
}




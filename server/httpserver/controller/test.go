package controller

import "github.com/gin-gonic/gin"

func Test(c *gin.Context) {
	// do something
	// ...
	// do something
	c.JSON(HttpCode, NewResp(HttpCodeSucc, "ok"))
}

func Test2(c *gin.Context) {
	// do something
	// ...
	// do something
	c.JSON(HttpCode, NewResp(HttpCodeFail, "err msg"))
}

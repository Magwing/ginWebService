package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Wan-Mi/ginWebService/server"
)

func ping(c *gin.Context) (int, interface{}, error) {
	return server.StatusOK, "ping success", nil
}

func init() {
	server.GET("/ping", ping)
}

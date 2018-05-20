package server

import "github.com/gin-gonic/gin"

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,OPTIONS,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		c.Next()
	}
}

func RegisteMiddleware(c *gin.Engine) {
	c.Use(cors())
	c.Use(gin.Recovery())
}

package ginengine

import (
	"LocalAPIGateway/src/common"
	"log"

	"github.com/gin-gonic/gin"
)

func EngineStarter() {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		log.Println("enter ping return pong")
		clientIP := c.ClientIP()
		log.Println("client ip : ", clientIP)
		c.JSON(200, "pong")
	})
	engine.GET("/pingping", func(c *gin.Context) {
		log.Println("enter pingping return pongpong")
		c.JSON(200, "pongpong")
	})

	engine.Run(":" + common.ServicePort)
}

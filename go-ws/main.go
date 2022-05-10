package main

import (
	"github.com/gin-gonic/gin"
)

//server
func main() {

	gin.SetMode(gin.DebugMode) //线上环境

	go Manager.Start()
	r := gin.Default()
	r.Static("/pub", "public")
	r.GET("/ws", WsHandler)
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9050") // listen and serve on 0.0.0.0:8080
}

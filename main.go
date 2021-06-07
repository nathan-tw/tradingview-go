package main

import (

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", Ping)
	r.POST("/futures_webhook", handleFuturesStrategy)
	r.POST("/webhook", handleStrategy)
	r.POST("/", Test)
	r.Run()
}
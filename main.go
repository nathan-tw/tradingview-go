package main

import (
	"github.com/nathan-tw/tradingview-go/src"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", src.Ping)
	r.POST("/futures_webhook", src.HandleFuturesStrategy)
	r.POST("/webhook", src.HandleStrategy)
	r.POST("/", src.Test)
	r.Run()
}
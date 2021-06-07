package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathan-tw/tradingview-go/src"
)

func main() {
	r := gin.Default()
	r.GET("/", src.Ping)
	r.Use(src.LogResponseBody)
	r.POST("/futures_webhook", src.HandleFuturesStrategy)
	r.POST("/webhook", src.HandleStrategy)
	r.POST("/test_receive_alert", src.TestReceiveAlert)
	r.POST("/test_binance_client", src.TestBinanceClient)
	r.Run()
}

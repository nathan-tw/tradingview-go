package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathan-tw/tradingview-go/src"
)

func main() {
	location, _ := time.LoadLocation(os.Getenv("TZ"))
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%v]: %3d | %15s | %-7s %#v\n%s",
			param.TimeStamp.In(location).Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))
    r.Use(gin.Recovery())
	r.Use(src.LogResponseBody)
    r.GET("/", src.Ping)
	r.POST("/futures_webhook", src.HandleFuturesStrategy)
	r.POST("/webhook", src.HandleStrategy)
	r.POST("/test_receive_alert", src.TestReceiveAlert)
	r.POST("/test_binance_client", src.TestBinanceClient)
	r.Run()
}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathan-tw/tradingview-go/src/middlewares/binance"
	"github.com/nathan-tw/tradingview-go/src/middlewares/general"
	"github.com/nathan-tw/tradingview-go/src/middlewares/logger"
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
	r.Use(logger.LogResponseBody)
	r.GET("/ping", general.Ping)
	r.POST("/binance_futures_webhook", binance.HandleFuturesStrategy)
	r.POST("/binance_futures_webhook_rat", binance.HandleFuturesStrategyForRat)
	r.POST("/binance_webhook", binance.HandleStrategy)
	r.Run()
}

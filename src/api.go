package src

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "tradingview-go",
	})
}

func TestBinanceClient(c *gin.Context) {
	fmt.Println(apiKey, apiSecret)
	client := binance.NewClient(apiKey, apiSecret)
	orders, err := client.NewListOrdersService().Symbol("ETHUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func TestReceiveAlert(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	c.String(200, string(jsonData))
	alert := new(TradingviewAlert)
	err = json.Unmarshal(jsonData, alert)
	if err != nil {
		panic(err)
	}
	fmt.Println(alert.Strategy.OrderAction)
}


func HandleFuturesStrategy(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	alert := new(TradingviewAlert)
	err = json.Unmarshal(jsonData, alert)
	if err != nil {
		panic(err)
	}
	if ok := validatePassPhrase(alert); !ok {
		c.String(http.StatusBadRequest, "wrong passphrase")
		return
	}
	side := strings.ToUpper(alert.Strategy.OrderAction)
	quantity := alert.Strategy.OrderContracts
	symbol := alert.Ticker
	fmt.Printf("trading side: %v, quantity: %v", side, quantity)
	err = createFuturesOrder(symbol, side, strconv.Itoa(quantity))
	if err != nil {
		c.String(http.StatusBadRequest, "create futures order fail %v", err)
		return
	}
	c.String(http.StatusOK, "create futures order success")
}

func HandleStrategy(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	alert := new(TradingviewAlert)
	err = json.Unmarshal(jsonData, alert)
	if err != nil {
		panic(err)
	}
	if ok := validatePassPhrase(alert); !ok {
		c.String(http.StatusBadRequest, "wrong passphrase")
		return
	}
	side := strings.ToUpper(alert.Strategy.OrderAction)
	quantity := alert.Strategy.OrderContracts
	symbol := alert.Ticker
	fmt.Printf("trading side: %v, quantity: %v", side, quantity)
	err = createOrder(symbol, side, strconv.Itoa(quantity))
	if err != nil {
		c.String(http.StatusBadRequest, "create order fail %v", err)
		return
	}
	c.String(http.StatusOK, "create order success")
}


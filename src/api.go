package src

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/gin-gonic/gin"
)

var (
	apiKey    string = os.Getenv("BINANCE_API_KEY")
	apiSecret string = os.Getenv("BINANCE_API_SECRET")
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogResponseBody(c *gin.Context) {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	fmt.Println("Response body: " + w.body.String())
}

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
	quantity := fmt.Sprintf("%f", alert.Strategy.OrderContracts)
	symbol := alert.Ticker
	fmt.Printf("trading side: %v, quantity: %v\n", side, quantity)
	futuresClient := binance.NewFuturesClient(apiKey, apiSecret)
	order, err := futuresClient.NewCreateOrderService().Symbol(symbol).Side(futures.SideType(side)).Type(futures.OrderTypeMarket).Quantity(quantity).Do(context.Background())
	if err != nil {
		c.String(http.StatusBadRequest, "create futures order fail %v", err)
		return
	}
	fmt.Println(order)
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
	quantity := fmt.Sprintf("%f", alert.Strategy.OrderContracts)
	symbol := alert.Ticker
	fmt.Printf("trading side: %v, quantity: %v", side, quantity)
	client := binance.NewClient(apiKey, apiSecret)
	order, err := client.NewCreateOrderService().Symbol(symbol).Side(binance.SideType(side)).Type(binance.OrderTypeMarket).Quantity(quantity).Do(context.Background())
	if err != nil {
		c.String(http.StatusBadRequest, "create order fail %v", err)
		return
	}
	fmt.Println(order)
	c.String(http.StatusOK, "create order success")
}

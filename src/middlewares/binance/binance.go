package binance

import (
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
	"github.com/nathan-tw/tradingview-go/src/webhook"
)

var (
	apiKey    string = os.Getenv("BINANCE_API_KEY")
	apiSecret string = os.Getenv("BINANCE_API_SECRET")
)


func HandleFuturesStrategy(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	alert := new(webhook.TradingviewAlert)
	err = json.Unmarshal(jsonData, alert)
	if err != nil {
		panic(err)
	}
	if ok := webhook.ValidatePassPhrase(alert); !ok {
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
	alert := new(webhook.TradingviewAlert)
	err = json.Unmarshal(jsonData, alert)
	if err != nil {
		panic(err)
	}
	if ok := webhook.ValidatePassPhrase(alert); !ok {
		c.String(http.StatusBadRequest, "wrong passphrase")
		return
	}

	side := strings.ToUpper(alert.Strategy.OrderAction)
	quantity := fmt.Sprintf("%f", alert.Strategy.OrderContracts)
	symbol := alert.Ticker
	fmt.Printf("trading side: %v, quantity: %v\n", side, quantity)
	client := binance.NewClient(apiKey, apiSecret)
	
	order, err := client.NewCreateOrderService().Symbol(symbol).Side(binance.SideType(side)).Type(binance.OrderTypeMarket).Quantity(quantity).Do(context.Background())
	if err != nil {
		c.String(http.StatusBadRequest, "create order fail %v", err)
		return
	}
	fmt.Println(order)
	c.String(http.StatusOK, "create order success")
}




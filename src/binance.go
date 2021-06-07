package src

import(
	"os"
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

var(
	apiKey string = os.Getenv("BINANCE_API_KEY")
	apiSecret string = os.Getenv("BINANCE_API_SECRET")
)

func createFuturesOrder(symbol, side, quantity string) error {
	futuresClient := binance.NewFuturesClient(apiKey, apiSecret)
	order, err := futuresClient.NewCreateOrderService().Symbol(symbol).Side(futures.SideType(side)).Type(futures.OrderTypeMarket).Quantity(quantity).TimeInForce(futures.TimeInForceTypeGTC).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(order)
	return nil
}

func createOrder(symbol, side, quantity string) error {
	client := binance.NewClient(apiKey, apiSecret)
	order, err := client.NewCreateOrderService().Symbol(symbol).Side(binance.SideType(side)).Type(binance.OrderTypeMarket).Quantity(quantity).TimeInForce(binance.TimeInForceTypeGTC).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(order)
	return nil
}

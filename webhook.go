package main

import "time"

type TradingviewAlert struct {
	Passphrase string    `json:"passphrase"`
	Time       time.Time `json:"time"`
	Exchange   string    `json:"exchange"`
	Ticker     string    `json:"ticker"`
	Bar        struct {
		Time   time.Time `json:"time"`
		Open   float64   `json:"open"`
		High   float64   `json:"high"`
		Low    float64   `json:"low"`
		Close  float64   `json:"close"`
		Volume int       `json:"volume"`
	} `json:"bar"`
	Strategy struct {
		PositionSize           int     `json:"position_size"`
		OrderAction            string  `json:"order_action"`
		OrderContracts         string     `json:"order_contracts"`
		OrderPrice             float64 `json:"order_price"`
		OrderID                string  `json:"order_id"`
		MarketPosition         string  `json:"market_position"`
		MarketPositionSize     int     `json:"market_position_size"`
		PrevMarketPosition     string  `json:"prev_market_position"`
		PrevMarketPositionSize int     `json:"prev_market_position_size"`
	} `json:"strategy"`
}
# Tradingview-go

<img src="asset/tradingview-go.png" width="600"/>

Tradingview-go is a crypto-trading bot triggered by tradingview webhook, and invoke api on heroku.
Currently support binance order and futures.

## Usage


#### Local

Add environment variables for binance.

```bash
export BINANCE_API_KEY="YOUR_API_KEY"
export BINANCE_API_SECRET="YOUR_API_SECRET"
```

For security concern, a passphrase is needed when a webhook send to the server, so also add a `PASSPHRASE` value.

```bash
export PASSPHRASE="YOUR_PASSPHRASE"
```

#### Heroku

Add environment variables in `Setting -> Config Variables`.

<img src="asset/heroku_env.png" width="800"/>

#### TradingView Alert

After pasting your webhook url, paste the contents of `tradingview_webhook_payload_format.txt` to the message block.

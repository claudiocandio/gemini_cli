package main

import (
	"encoding/json"
	"fmt"

	"github.com/claudiocandio/gemini-api"
)

func get_account(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	accountDetail, err := api.AccountDetail()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&accountDetail, "", " ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", j), nil
}

func get_tradevolume(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	tradeVolume, err := api.TradeVolume()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&tradeVolume, "", " ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", j), nil
}

func get_depositaddresses(gemini_config_yml string, currency string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	depositAddresses, err := api.DepositAddresses(currency)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&depositAddresses, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func get_symbols(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// get Symbols
	symbols, err := api.Symbols()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&symbols, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func get_balances(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// get Balances
	balances, err := api.Balances()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&balances, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func get_ticker(gemini_config_yml string, ticker string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// get TickerV2
	tickerV2, err := api.TickerV2(ticker)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&tickerV2, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func get_auction(gemini_config_yml string, ticker string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// get Current Auction
	currentAuction, err := api.CurrentAuction(ticker)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&currentAuction, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

// gemini.Args{"timestampms": 0, "limitTrades": 30, "includeBreaks": false}
func get_trades(gemini_config_yml string, ticker string, args gemini.Args) (string, error) {

	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	trades, err := api.Trades(ticker, args)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&trades, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

// Args{"limit_bids": 1, "limit_asks": 1}
func get_orderbook(gemini_config_yml string, ticker string, args gemini.Args) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	orderBook, err := api.OrderBook(ticker, args)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&orderBook, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

func get_orders(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// Get your active orders
	activeOrders, err := api.ActiveOrders()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&activeOrders, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

// Get list of my past trades
// Args{"limit_trades": 0, "timestamp": "2021-12-01T15:04:01"}
func get_past_trades(gemini_config_yml string, ticker string, args gemini.Args) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	pastTrades, err := api.PastTrades(ticker, args)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&pastTrades, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

func get_order(gemini_config_yml string, orderId string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// get order
	order, err := api.OrderStatus(orderId)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&order, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

// gemini.Args{"since": 0, "limit": 100, "includeIndicative": true}
func get_auction_hystory(gemini_config_yml string, ticker string, args gemini.Args) (string, error) {

	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	auctionHistory, err := api.AuctionHistory(ticker, args)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&auctionHistory, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

func cancel_order(gemini_config_yml string, orderId string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// cancel order
	cancelOrder, err := api.CancelOrder(orderId)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&cancelOrder, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func cancel_all_order(gemini_config_yml string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	// This will cancel ALL orders includind those placed through the UI
	cancelAll, err := api.CancelAll()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&cancelAll, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil
}

func new_order(gemini_config_yml, ticker, clientOrderId, side string, amount, price float64) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	newOrder, err := api.NewOrder(ticker, clientOrderId, amount, price, side, []string{"immediate-or-cancel"})
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&newOrder, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

// currency can be bitcoin, ethereum, bitcoincash, litecoin, zcash, or filecoin
func new_deposit_address(gemini_config_yml, currency, label string) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	newDepositAddress, err := api.NewDepositAddress(currency, label)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&newDepositAddress, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

// Withdraw Crypto Funds
// currency can be btc or eth
func withdraw_funds(gemini_config_yml, currency, address string, amount float64) (string, error) {
	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	withdrawFunds, err := api.WithdrawFunds(currency, address, amount)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&withdrawFunds, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

// Args{"timestamp": "2021-12-01T15:04:01", "limit_transfers": 20,"show_completed_deposit_advances": false}
func get_transfers(gemini_config_yml string, args gemini.Args) (string, error) {

	api, err := start_api(gemini_config_yml)
	if err != nil {
		return "", err
	}

	transfers, err := api.Transfers(args)
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(&transfers, "", " ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", j), nil

}

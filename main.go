package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/claudiocandio/gemini-api"
)

var stdlog, errlog *log.Logger

func init() {
	stdlog = log.New(os.Stdout, "", 0)
	errlog = log.New(os.Stderr, "", 0)
}

func main() {

	var gemini_config_yml string

	app := &cli.App{
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,

		Name:    "gemini_cli",
		Usage:   "resti-api cli commands, reference: https://docs.gemini.com/rest-api/",
		Version: "v1.0.2",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage: "--config gemini.yml - Load yml configuration file e.g.\n" +
					"	gemini_api_credentials:\n" +
					"		gemini_api_key: \"mygeminikey\"\n" +
					"		gemini_api_secret: \"mygeminisecret\"\n" +
					"		gemini_api_production: \"false\" (if false it uses sandbox server)\n" +
					"	-\n" +
					"	Instead of a configuration file you can export the following environment variables:\n" +
					"		export GEMINI_API_KEY=\"mygeminikey\"\n" +
					"		export GEMINI_API_SECRET=\"mygeminisecret\"\n" +
					"		export GEMINI_API_PRODUCTION=\"false\"\n" +
					"	Yml configuration file does override the environment variables\n" +
					"	-\n",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Run in debug mode",
			},
			&cli.BoolFlag{
				Name:  "trace",
				Usage: "More debug, this will also show gemini key and secret !",
			},
		},

		Commands: []*cli.Command{
			{
				Name: "get",
				Subcommands: []*cli.Command{
					{
						Name:  "account",
						Usage: "Get account details (Private)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/account
							status, err := get_account(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "balances",
						Usage: "This will show the available balances in the supported currencies (Private)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/balances
							status, err := get_balances(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},

					{
						Name:  "transfers",
						Usage: "shows deposits and withdrawals in the supported currencies (Private)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "timestamp",
								Aliases: []string{"t"},
								Usage:   "e.g. --timestamp 2021-02-05T15:04:01",
							},
							&cli.StringFlag{
								Name:    "limit_transfers",
								Aliases: []string{"l"},
								Usage:   "e.g. --limit_transfers 10",
							},
							&cli.BoolFlag{
								Name:    "show_completed_deposit_advances",
								Aliases: []string{"s"},
								Usage:   "e.g. --show_completed_deposit_advances",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)

							args := gemini.Args{}
							if c.IsSet("limit_transfers") {
								limit_transfers := c.Int("limit_transfers")
								args["limit_transfers"] = limit_transfers
							}
							if c.IsSet("timestamp") || c.IsSet("t") {
								timestamp, err := parseConvertTimestamp(c.String("timestamp"))
								if err != nil {
									return err
								}
								args["timestamp"] = get_timestampms(*timestamp)
							}
							if c.Bool("show_completed_deposit_advances") {
								args["show_completed_deposit_advances"] = "true"
							}
							// /v1/transfers
							status, err := get_transfers(gemini_config_yml, args)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "depositaddresses",
						Usage: "Get deposit addresses (Private)",

						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "currency",
								Aliases:  []string{"c"},
								Usage:    "--currency bitcoin|ethereum|bitcoincash|litecoin|zcash|filecoin",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							validCurrencies := []string{"bitcoin", "ethereum", "bitcoincash", "litecoin", "zcash", "filecoin"}
							currency := c.String("currency")
							for _, value := range validCurrencies {
								if currency == value {
									// /v1/addresses/:network
									status, err := get_depositaddresses(gemini_config_yml, currency)
									if err != nil {
										return err
									}
									stdlog.Print(status)
									return nil
								}
							}
							return fmt.Errorf("Error invalid currency: %s\nValid currencies: %v", currency, strings.Join(validCurrencies, ", "))
						},
					},
					{
						Name:  "new_depositaddresses",
						Usage: "Generate a new deposit addresses (Private)",

						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "currency",
								Aliases:  []string{"c"},
								Usage:    "--currency bitcoin|ethereum|bitcoincash|litecoin|zcash|filecoin",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "label",
								Aliases: []string{"l"},
								Usage:   "Optional label for the deposit address",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							validCurrencies := []string{"bitcoin", "ethereum", "bitcoincash", "litecoin", "zcash", "filecoin"}
							currency := c.String("currency")
							for _, value := range validCurrencies {
								if currency == value {
									// /v1/deposit/:network/newAddress
									status, err := new_deposit_address(gemini_config_yml, currency, c.String("label"))
									if err != nil {
										return err
									}
									stdlog.Print(status)
									return nil
								}
							}
							return fmt.Errorf("Error invalid currency: %s\nValid currencies: %v", currency, strings.Join(validCurrencies, ", "))
						},
					},
					{
						Name:  "symbols",
						Usage: "Retrieves all available symbols for trading (Public)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/symbols
							status, err := get_symbols(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "ticker",
						Usage: "This endpoint retrieves information about recent trading activity for the provided symbol (Public)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "ticker",
								Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v2/ticker/:symbol
							status, err := get_ticker(gemini_config_yml, c.String("ticker"))
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "tradevolume",
						Usage: "Get trade volume, up to 30 days of trade volume for each symbol (Private)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/tradevolume
							status, err := get_tradevolume(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "trades",
						Usage: "Trades that have executed since the specified timestamp (Public)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "ticker",
								//Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "timestamp",
								Aliases: []string{"t"},
								Usage:   "e.g. --timestamp 2021-02-05T15:04:01",
							},
							&cli.StringFlag{
								Name:    "limit_trades",
								Aliases: []string{"l"},
								Usage:   "e.g. --limit_trades 10 (default 50, set to 0 for all)",
							},
							&cli.BoolFlag{
								Name:    "include_breaks",
								Aliases: []string{"b"},
								Usage:   "e.g. --include_breaks",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							args := gemini.Args{}
							if c.IsSet("timestamp") || c.IsSet("t") {
								timestamp, err := parseConvertTimestamp(c.String("timestamp"))
								if err != nil {
									return err
								}
								args["timestamp"] = *timestamp
							}

							if c.IsSet("limit_trades") {
								limit_trades := c.Int("limit_trades")
								args["limit_trades"] = strconv.Itoa(limit_trades)
							}
							if c.Bool("include_breaks") {
								args["include_breaks"] = "true"
							}
							// /v1/trades/:symbol
							status, err := get_trades(gemini_config_yml, c.String("ticker"), args)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "orderbook",
						Usage: "This will return the current order book as two arrays bids/asks (Public)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "ticker",
								Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "limit_bids",
								Aliases: []string{"b"},
								Usage:   "e.g. --limit_bids 10 (default 50, set to 0 for all)",
							},
							&cli.StringFlag{
								Name:    "limit_asks",
								Aliases: []string{"a"},
								Usage:   "e.g. --limit_asks 10 (default 50, set to 0 for all)",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)

							args := gemini.Args{}
							if c.IsSet("limit_bids") {
								limitBids := c.Int("limit_bids")
								args["limit_bids"] = strconv.Itoa(limitBids)
							}
							if c.IsSet("limit_asks") {
								limitAsks := c.Int("limit_asks")
								args["limit_asks"] = strconv.Itoa(limitAsks)
							}
							// /v1/book/:symbol
							status, err := get_orderbook(gemini_config_yml, c.String("ticker"), args)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "auction",
						Usage: "Current auction (Public)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "ticker",
								Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/auction/:symbol
							status, err := get_auction(gemini_config_yml, c.String("ticker"))
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},

					// gemini.Args{"since": 0, "limit": 100, "includeIndicative": true}
					{
						Name:  "auction-hystory",
						Usage: "This will return the auction events, optionally including publications of indicative prices (Public)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "ticker",
								Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "since",
								Aliases: []string{"s"},
								Usage:   "e.g. --since 2021-02-05T15:04:01",
							},
							&cli.StringFlag{
								Name:    "limit",
								Aliases: []string{"l"},
								Usage:   "e.g. --limit 10 (default 50, set to 0 for all)",
							},
							&cli.BoolFlag{
								Name:    "include_indicative",
								Aliases: []string{"i"},
								Usage:   "e.g. --include_indicative",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							args := gemini.Args{}
							if c.IsSet("since") || c.IsSet("s") {
								since, err := parseConvertTimestamp(c.String("since"))
								if err != nil {
									return err
								}
								args["since"] = *since
							}

							if c.IsSet("limit") {
								limit_auction_results := c.Int("limit")
								args["limit_auction_results"] = strconv.Itoa(limit_auction_results)
							}
							if c.Bool("include_indicative") {
								args["include_indicative"] = "true"
							}
							// /v1/auction/:symbol/history
							status, err := get_auction_hystory(gemini_config_yml, c.String("ticker"), args)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
				},
			},
			{
				Name: "order",
				Subcommands: []*cli.Command{
					{
						Name:  "new",
						Usage: "Place a new order (Private)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "client_order_id",
								Aliases: []string{"i"},
								Usage:   "e.g. --client_order_id \"20170208_example\" (Oprional but recommended)",
							},
							&cli.StringFlag{
								Name:     "ticker",
								Aliases:  []string{"t"},
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "side",
								Aliases:  []string{"s"},
								Usage:    "e.g. --side buy (buy or sell)",
								Required: true,
							},
							&cli.Float64Flag{
								Name:     "amount",
								Aliases:  []string{"a"},
								Usage:    "e.g. --amount 0.021 (Decimal amount to purchase)",
								Required: true,
							},
							&cli.Float64Flag{
								Name:     "price",
								Aliases:  []string{"p"},
								Usage:    "e.g. --price 3633.00 (Decimal amount to spend per unit)",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/order/new
							status, err := new_order(gemini_config_yml,
								c.String("ticker"),
								c.String("client_order_id"),
								c.String("side"),
								c.Float64("amount"),
								c.Float64("price"),
							)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "active",
						Usage: "Get active orders (Private)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/orders
							status, err := get_orders(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "past_trades",
						Usage: "Get past trades (Private)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "ticker",
								Usage:    "e.g. --ticker btcusd (ticker is required)",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "limit_trades",
								Aliases: []string{"l"},
								Usage:   "e.g. --limit_trades 10",
							},
							&cli.StringFlag{
								Name:    "timestamp",
								Aliases: []string{"t"},
								Usage:   "e.g. --timestamp 2021-02-05T15:04:01",
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)

							args := gemini.Args{}
							if c.IsSet("limit_trades") {
								limit_trades := c.Int("limit_trades")
								args["limit_trades"] = limit_trades
							}
							if c.IsSet("timestamp") || c.IsSet("t") {
								timestamp, err := parseConvertTimestamp(c.String("timestamp"))
								if err != nil {
									return err
								}
								args["timestamp"] = *timestamp
							}
							// /v1/mytrades
							status, err := get_past_trades(gemini_config_yml, c.String("ticker"), args)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "orderid",
						Usage: "Get order status (Private)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "orderid",
								Aliases:  []string{"o"},
								Usage:    "e.g. --orderid 121212 (orderid is required)",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/order/status
							status, err := get_order(gemini_config_yml, c.String("order"))
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "cancel",
						Usage: "Cancel an order. If the order is already canceled, the message will succeed but have no effect (Private)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "orderid",
								Aliases:  []string{"o"},
								Usage:    "e.g. --orderid 121212 (orderid is required)",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/order/cancel
							status, err := cancel_order(gemini_config_yml, c.String("orderid"))
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
					{
						Name:  "cancel_all",
						Usage: "Cancel ALL orders includind those placed through the UI !!! (Private)",
						Action: func(c *cli.Context) error {
							gemini_config_yml = parse_params(c)
							// /v1/order/cancel/all
							status, err := cancel_all_order(gemini_config_yml)
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						},
					},
				},
			},

			{
				Name:  "withdraw",
				Usage: "Withdraw Crypto Funds - You must have an approved address list for your account (Private)",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "currency",
						Aliases:  []string{"c"},
						Usage:    "--currency btc|eth",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "address",
						Aliases:  []string{"a"},
						Usage:    "Standard string format of cryptocurrency address",
						Required: true,
					},
					&cli.Float64Flag{
						Name:     "amount",
						Usage:    "e.g. --amount 0.021 (Decimal amount to purchase)",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					gemini_config_yml = parse_params(c)
					validCurrencies := []string{"btc", "eth"}
					currency := c.String("currency")
					for _, value := range validCurrencies {
						if currency == value {
							// /v1/withdraw/:currency
							status, err := withdraw_funds(gemini_config_yml, currency, c.String("address"), c.Float64("amount"))
							if err != nil {
								return err
							}
							stdlog.Print(status)
							return nil
						}
					}
					return fmt.Errorf("Error invalid currency: %s\nValid currencies: %v", currency, strings.Join(validCurrencies, ", "))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		errlog.Println(err)
	}

	/*
		TO BE DONE

		// currency can be btc or eth
		func withdraw_funds(gemini_config_yml, currency, address string, amount float64) (string, error) {

			// Withdraw Funds
			withdrawFunds, err := api.WithdrawFunds("btc", "mpXEeWuc7tSB4BVEoYpStmj4MFrmrLiRKL", 0.01)
			if err != nil {
				log.Fatalf("err = %v\n", err)
			}
			j, err = json.MarshalIndent(&withdrawFunds, "", " ")
			if err != nil {
				log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
			}
			fmt.Printf("withdrawFunds = %s\n\n", j)

			//This will cancel all orders opened by this session.
			//This will have the same effect as heartbeat expiration if "Require Heartbeat" is selected for the session.
			cancelSession, err := api.CancelSession()
			if err != nil {
				log.Fatalf("err = %v\n", err)
			}
			j, err = json.MarshalIndent(&cancelSession, "", " ")
			if err != nil {
				log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
			}
			fmt.Printf("cancelSession = %s\n\n", j)

			// Heartbeat
			// This will prevent a session from timing out and canceling orders if the
			// require heartbeat flag has been set. Note that this is only required if
			// no other private API requests have been made. The arrival of any message
			// resets the heartbeat timer.
			heartbeat, err := api.Heartbeat()
			if err != nil {
				log.Fatalf("err = %v\n", err)
			}
			j, err = json.MarshalIndent(&heartbeat, "", " ")
			if err != nil {
				log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
			}
			fmt.Printf("heartbeat = %s\n\n", j)

	*/
}

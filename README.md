[![GitHub license](https://claudiocandio.github.io/img/license_mit.svg)](https://github.com/claudiocandio/gemini-api/blob/master/LICENSE)
[![Language: Go](https://claudiocandio.github.io/img/language-Go.svg)](https://golang.org/)
[![Donate Bitcoin](https://claudiocandio.github.io/img/donate-bitcoin-orange.svg)](https://claudiocandio.github.io/img/donate-bitcoin.html)
[![Donate Ethereum](https://claudiocandio.github.io/img/donate-etherum-green.svg)](https://claudiocandio.github.io/img/donate-ethereum.html)

# CLI command for the Gemini Exchange REST API

Gemini CLI can connect to the production site <https://api.gemini.com> or to the Sandbox site <https://api.sandbox.gemini.com> for testing purposes.

gemini_cli uses the REST API wrapper <https://github.com/claudiocandio/gemini-api> I haven't had time to fully test everything, it is working fine for me but use it at your own risk.

## gemini_cli Build

You need Go installed <https://golang.org>, I'm using go1.15.8

```bash
$ git clone https://github.com/claudiocandio/gemini_cli
$ cd gemini_cli
$ go build .
$ ./gemini_cli --help
```

## gemini_cli API Key

Before using gemini_cli you need to have your Gemini API Key. You can create the API Key from the Gemini site <https://www.gemini.com> check the appropriate Roles: <https://docs.gemini.com/rest-api/#roles>

If you plan to test gemini_cli you need to create the API Key in the Sandbox site <https://exchange.sandbox.gemini.com>

You can use the API Key in gemini_cli either via yml configuration file or using environment variables:

### Yml configuration file

```bash
$ cat gemini_config.yml
gemini_api_credentials:
  gemini_api_key: "mygeminikey"
  gemini_api_secret: "mygeminisecret"
  gemini_api_production: "false"
```

### Use environment variables

```bash
$ export GEMINI_API_KEY="mygeminikey"
$ export GEMINI_API_SECRET="mygeminisecret"
$ export GEMINI_API_PRODUCTION="false"
```

If the yml gemini_api_production or env GEMINI_API_PRODUCTION is false then gemini_cli will point to the Gemini Sandbox site <https://api.sandbox.gemini.com>

If the yml gemini_api_production or env GEMINI_API_PRODUCTION is true then gemini_cli will point to the Gemini Production site <https://api.gemini.com> (Real Money !!)

To use the yml configuration file:

```bash
$ gemini_cli --config gemini_config.yml ...
```

Otherwise with environment variables just:

```bash
$ gemini_cli ...
```

Using a yml configuration file will override the environment variables.

### gemini_cli bash autocompletition

It is possible to enable gemini_cli bash auto-completion for the current shell session, it needs to use bash_autocomplete script included in this repo and provided by <https://github.com/urfave/cli>

To use bash_autocomplete set an environment variable named PROG to gemini_cli as follow:

```bash
$ PROG=gemini_cli source path/to/cli/autocomplete/bash_autocomplete
```

### gemini_cli Help

```bash
$ gemini_cli --help
NAME:
   gemini_cli - resti-api cli commands, reference: https://docs.gemini.com/rest-api/

USAGE:
   gemini_cli [global options] command [command options] [arguments...]

VERSION:
   v1.0.2

COMMANDS:
   get       
   order     
   withdraw  Withdraw Crypto Funds (Private)
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  --config gemini.yml - Load yml configuration file e.g.
                             gemini_api_credentials:
                               gemini_api_key: "mygeminikey"
                               gemini_api_secret: "mygeminisecret"
                               gemini_api_production: "false" (if false it uses sandbox server)
                             -
                             Instead of a configuration file you can export the following environment variables:
                               export GEMINI_API_KEY="mygeminikey"
                               export GEMINI_API_SECRET="mygeminisecret"
                               export GEMINI_API_PRODUCTION="false"
                             Yml configuration file does override the environment variables
                             -
   --debug, -d               Run in debug mode (default: false)
   --trace                   More debug, this will also show gemini key and secret ! (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)

```

### gemini_cli get Help

```bash
$ gemini_cli get --help
NAME:
   gemini_cli get - A new cli application

USAGE:
   gemini_cli get command [command options] [arguments...]

COMMANDS:
   account               Get account details (Private)
   balances              This will show the available balances in the supported currencies (Private)
   transfers             shows deposits and withdrawals in the supported currencies (Private)
   depositaddresses      Get deposit addresses (Private)
   new_depositaddresses  Generate a new deposit addresses (Private)
   symbols               Retrieves all available symbols for trading (Public)
   ticker                This endpoint retrieves information about recent trading activity for the provided symbol (Public)
   tradevolume           Get trade volume, up to 30 days of trade volume for each symbol (Private)
   trades                Trades that have executed since the specified timestamp (Public)
   orderbook             This will return the current order book as two arrays bids/asks (Public)
   auction               Current auction (Public)
   auction-hystory       This will return the auction events, optionally including publications of indicative prices (Public)
   help, h               Shows a list of commands or help for one command
```

### gemini_cli order Help

```bash
$ gemini_cli order --help
NAME:
   gemini_cli order - A new cli application

USAGE:
   gemini_cli order command [command options] [arguments...]

COMMANDS:
   new          Place a new order (Private)
   active       Get active orders (Private)
   past_trades  Get past trades (Private)
   orderid      Get order status (Private)
   cancel       Cancel an order. If the order is already canceled, the message will succeed but have no effect (Private)
   cancel_all   Cancel ALL orders includind those placed through the UI !!! (Private)
   help, h      Shows a list of commands or help for one command
```

## Usage Examples

To get account information:

```bash
$ gemini_cli get account
{
 "account": {
  "accountname": "Primary",
  "shortname": "primary",
  "type": "exchange",
  "created": "1612381252885",
  "createdt": "2021-02-01T22:55:42.685+01:00"
 },
 "users": [
  {
   "name": "Claudio Candio sandbox",
   "lastsignin": "2021-02-12T21:01:54.26Z",
   "status": "Active",
   "countrycode": "IT",
   "isverified": true
  }
 ],
 "memo_reference_code": "FAKECODE"
}
```

To get balance information:

```bash
$ gemini_cli get balances
[
 {
  "currency": "BTC",
  "amount": "0.001",
  "available": "0.001",
  "availableForWithdrawal": "0.001",
  "type": "exchange"
 },
 {
  "currency": "ETH",
  "amount": "0.01",
  "available": "0.01",
  "availableForWithdrawal": "0.01",
  "type": "exchange"
 },
 ...
 ```

To get ticker btcusd information with debug information:

```bash
$ gemini_cli --debug -c /home/joker/gemini_sanbox.yml get ticker -t btcusd
DEBUG  [2021-02-16T22:10:42+01:00] Debug enabled                                
DEBUG  [2021-02-16T22:10:42+01:00] Connecting to Gemini Sandbox site.           
DEBUG  [2021-02-16T22:10:42+01:00] func TickerV2                                 url:="https://api.sandbox.gemini.com/v2/ticker/btcusd"
DEBUG  [2021-02-16T22:10:42+01:00] func request: http.NewRequest                 params:="map[]" url:="https://api.sandbox.gemini.com/v2/ticker/btcusd" verb:=GET
DEBUG  [2021-02-16T22:10:43+01:00] func request: Http Client response            resp:="&{200 OK 200 HTTP/2.0 2 0 map[Content-Length:[398] Content-Type:[application/json] Date:[Tue, 16 Feb 2021 21:10:43 GMT] Server:[nginx] Vary:[Origin]] {0xc0005422c0} 398 [] false false map[] 0xc00016e300 0xc000116420}"
DEBUG  [2021-02-16T22:10:43+01:00] func request: Http Client body                body:="{\"symbol\":\"BTCUSD\",\"open\":\"46912.73\",\"high\":\"46912.73\",\"low\":\"45700\",\"close\":\"45810.54\",\"changes\":[\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"45810.54\",\"46912.73\",\"45810.54\"],\"bid\":\"20900.00\",\"ask\":\"25000.00\"}"
DEBUG  [2021-02-16T22:10:43+01:00] func TickerV2: unmarshal                      tickerV2:="{BTCUSD 46912.73 46912.73 45700 45810.54 [45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 45810.54 46912.73 45810.54] 20900 25000}"
{
 "symbol": "BTCUSD",
 "open": "46912.73",
 "high": "46912.73",
 "low": "45700",
 "close": "45810.54",
 "changes": [
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "45810.54",
  "46912.73",
  "45810.54"
 ],
 "bid": "20900",
 "ask": "25000"
}
```

Keep going using the gemini_cli --help as reference

Have fun !

## Disclaimer

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND
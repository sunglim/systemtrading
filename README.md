# systemtrading

My first system trading application

## Currently supporting brokerage

- [Korea investment(한국투자증권)](https://apiportal.koreainvestment.com/about) is the only brokerage supporting REST API in Korea.


# How to run

``` sh
> go run main.go -koreainvestment_url=<your korea investment server URL> -koreainvestment_appkey=<your korea invesment app key> \
     -koreainvestment_appsecret=<your korea invesment app secret> -koreainvestment_account=<your account> -telegram_chat_id=<telegram chat id> -telegram_token=<telegram token>
```

## Korea Investment API

`/pkg/koreainvestment` is a package to call Korea investment APIs. Some old codes still live in `/order/koreainvestment`, but new code should reside in pkg directory.

Ideally, I have a plan to make this package a fully generated code.

## Strategry

A strategy matches to an trading algorithm.

### BuyOneStockEveryDay stategry

As the name explains, this strategry buy aone stock every day.


More strategry will be introduced..

## Logging system

By default, the logging system writes messages to standard output, and also to telegram as long as telegram configuration is set.

### Sending log messages to telegram

The application requires a telegram bot token and chat ID. See [BotFather](https://core.telegram.org/bots/features#botfather) to get a telegram token.

See `go run ./src/main -h` explains how to pass the token and chat id.

# Development plan

[Development plan](./docs/development_plan.md)

# Research

* Existing system trading applications

#


# The system trading application should be

* A moudle to receive current stock price
  -  It should be pluggable, provide an interface so that any stockbrokerage can be integrated

* A moudle to buy/sell

* Strategry
  - When to sell, when to buy, when to hold

* Monitoring
  - Provide informations like profits, loss.
 
* Easy to simulate

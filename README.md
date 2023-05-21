# systemtrading
My first system trading application

## Currently supporting brokage

- Korea investment(한국투자증권)

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

# Run demo

``` sh
> go run main.go <domain> <appkey> <appsecret>
```

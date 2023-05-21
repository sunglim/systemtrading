# Development Plan


## Demo

The demo program should be able to order one single stock periodically.

## Setting environment variables

This application does not use [os.GetEnv()](https://pkg.go.dev/os#Getenv). Instead, receive all env variables
 as arguments. e.g.

 ```
 go run main.go -koreainvestment_url=http://xxx.www -koreainvestment_appsecret=abced
 ```

## Logging system

For system trading applications, logging system is very critically important feature as it handles *money*.
Sending log message to stdout, stderr is not enough. The message should be sent to the customer who is not
in front of laptop.

### Telegram integration

Using the [telegram bot](https://core.telegram.org/bots/features#botfather), customers can receive logging
messages with their phone or desktop. 

#### Useful link
- https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id

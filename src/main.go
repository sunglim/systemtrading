package main

import (
	"flag"
	gologger "log"

	log "sunglim.github.com/sunglim/log"
)

var exit = make(chan bool)

func main() {
	var telegramChatId int64
	flag.Int64Var(&telegramChatId, "telegram_chat_id", 1234, "telegram chat ID")
	var telegramToken string
	flag.StringVar(&telegramToken, "telegram_token", "default", "telegram token")
	var koreaInvestmentUrl string
	flag.StringVar(&koreaInvestmentUrl, "koreainvestment_url", "default", "kroeainvestment url")
	var koreaAppKey string
	flag.StringVar(&koreaAppKey, "koreainvestment_appkey", "default", "kroeainvestment appkey")
	var koreaAppSecret string
	flag.StringVar(&koreaAppSecret, "koreainvestment_appsecret", "default", "kroeainvestment appkey")
	var koreaAccount string
	flag.StringVar(&koreaAccount, "koreainvestment_account", "default", "kroeainvestment account")

	flag.Parse()

	log.SetTelegramLogger(gologger.New(log.CreateTelegramWriter(telegramToken, telegramChatId), "", gologger.Ldate|gologger.Ltime))
	log.Println("Starting after telegram: ", telegramToken)

	//go order.StrategryBuyEveryDay()

	// Infinite.
	<-exit
}

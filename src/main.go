package main

import (
	"flag"
	gologger "log"

	log "sunglim.github.com/sunglim/log"
	"sunglim.github.com/sunglim/order"
	"sunglim.github.com/sunglim/order/koreainvestment"
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

	if telegramToken != "default" {
		log.SetTelegramLogger(gologger.New(log.CreateTelegramWriter(telegramToken, telegramChatId), "", gologger.Ldate|gologger.Ltime))
	}
	log.Println("Starting after telegram: ", telegramToken)

	koreainvestment.Initialize(koreaInvestmentUrl, koreaAppKey, koreaAppSecret, koreainvestment.KoreaInvestmentAccount{
		CANO:         koreaAccount,
		ACNT_PRDT_CD: "01",
	})

	// Buy Samsung eletronics at 10 am.
	//go order.StrategryBuyEveryDay(koreaexchange.Code삼성전자, "10:00")
	order.Demo()

	// Infinite.
	<-exit
}

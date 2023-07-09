package main

import (
	"flag"
	gologger "log"

	krxcode "github.com/sunglim/go-korea-stock-code/code"
	log "sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
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
		telegramWriter := log.CreateTelegramWriter(telegramToken, telegramChatId)
		// nil when offline.
		if telegramWriter != nil {
			log.SetTelegramLogger(gologger.New(telegramWriter, "", gologger.Ldate|gologger.Ltime))
		}
	}
	log.Println("Starting after telegram: ", telegramToken)

	koreainvestment.Initialize(koreaInvestmentUrl, koreaAppKey, koreaAppSecret, ki.KoreaInvestmentAccount{
		CANO:         koreaAccount,
		ACNT_PRDT_CD: "01",
	})

	// Buy Samsung eletronics at 10 am.
	//go order.StrategryBuyEveryDay(koreaexchange.Code삼성전자, "10:00")

	go order.StrategryBuyEveryDayIfBelowAverage("03:07")

	go order.StrategryBuyEveryDayIfLowerThan("03:00", []order.CodePrice{
		{
			Code:  krxcode.Code농심홀딩스,
			Price: 67000,
		},
		{
			Code:  krxcode.CodeBNK금융지주,
			Price: 6800,
		},
		{
			Code:  krxcode.Code기업은행,
			Price: 9600,
		},
		{
			Code:  krxcode.CodeDGB금융지주,
			Price: 7000,
		},
	})
	//order.Demo()

	// Infinite.
	<-exit
}

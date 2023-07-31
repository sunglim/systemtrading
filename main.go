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
	flag.StringVar(&telegramToken, "telegram_token", "", "telegram token")
	var koreaInvestmentUrl string
	flag.StringVar(&koreaInvestmentUrl, "koreainvestment_url", "default", "kroeainvestment url")
	var koreaAppKey string
	flag.StringVar(&koreaAppKey, "koreainvestment_appkey", "default", "kroeainvestment appkey")
	var koreaAppSecret string
	flag.StringVar(&koreaAppSecret, "koreainvestment_appsecret", "default", "kroeainvestment appkey")
	var koreaAccount string
	flag.StringVar(&koreaAccount, "koreainvestment_account", "default", "kroeainvestment account")

	flag.Parse()

	if telegramToken != "" {
		telegramWriter := log.CreateTelegramWriter(telegramToken, telegramChatId)
		// nil when offline.
		if telegramWriter != nil {
			log.SetTelegramLogger(gologger.New(telegramWriter, "", gologger.Ldate|gologger.Ltime))
		}
	}

	koreainvestment.Initialize(koreaInvestmentUrl, koreaAppKey, koreaAppSecret, ki.KoreaInvestmentAccount{
		CANO:         koreaAccount,
		ACNT_PRDT_CD: "01",
	})

	// Buy Samsung eletronics at 10 am.
	//go order.StrategryBuyEveryDay(koreaexchange.Code삼성전자, "10:00")

	go order.StrategryBuyEveryDayIfBelowAverage("12:07", []order.StrategryBuyEveryDayIfBelowOrder{
		{
			Code:     krxcode.Code농심홀딩스,
			Quantity: 1,
		},
		{
			Code:     krxcode.CodeBNK금융지주,
			Quantity: 3,
		},
		{
			Code:     krxcode.CodeDGB금융지주,
			Quantity: 3,
		},
		{
			Code:     "102110", // tiger 200
			Quantity: 10,
		},
	})

	go order.StrategryBuyEveryDayIfLowerThan("12:00", []order.StrategryOrder{
		{
			Code:     krxcode.Code부국증권,
			Price:    17500,
			Quantity: 2,
		},
		{
			Code:     krxcode.CodeKB금융,
			Price:    48000,
			Quantity: 1,
		},
		{
			Code:     krxcode.Code삼성카드,
			Price:    28400,
			Quantity: 1,
		},
		{
			Code:     krxcode.Code삼성전자,
			Price:    60000,
			Quantity: 1,
		},
		{
			Code:     krxcode.Code신한지주,
			Price:    30000,
			Quantity: 1,
		},
		{
			Code:     krxcode.Code하나금융지주,
			Price:    33000,
			Quantity: 1,
		},
		{
			Code:     krxcode.CodeBNK금융지주,
			Price:    6500,
			Quantity: 3,
		},
		{
			Code:     krxcode.Code기업은행,
			Price:    9600,
			Quantity: 2,
		},
		{
			Code:     krxcode.CodeDGB금융지주,
			Price:    7000,
			Quantity: 3,
		},
		{
			Code:     krxcode.Code우리금융지주,
			Price:    11400,
			Quantity: 2,
		},
		{
			Code:     krxcode.Code신한지주,
			Price:    30000,
			Quantity: 1,
		},
		{
			Code:     krxcode.Code케이티앤지,
			Price:    80000,
			Quantity: 1,
		},
		{
			Code:     "102110", // tiger 200
			Price:    29000,
			Quantity: 10,
		},
		{
			Code:     "148020", // kbstar 200
			Price:    29000,
			Quantity: 10,
		},
	})

	sellStrategry := order.NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage("12:00", []order.StrategryBuyEveryDayIfBelowOrder{{}})
	go sellStrategry.Start()

	/*
		go order.StrategrySellEveryDayIfHigherThan("15:00", []order.StrategryOrder{{
			Code:     krxcode.Code농심홀딩스,
			Price:    67500,
			Quantity: 1,
		}})
	*/
	//order.Demo()

	// Infinite.
	<-exit
}

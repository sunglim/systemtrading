package order

import (
	"os"

	sunglimlog "sunglim.github.com/sunglim/log"
	"sunglim.github.com/sunglim/order/koreainvestment"
)

type Order struct {
	logger *sunglimlog.Logger
}

func (order Order) DailyOrder() {
	url := os.Args[1]
	appKey := os.Args[2]
	appSecret := os.Args[3]
	account := os.Args[4]
	koreainvestment.Initialize(url, appKey, appSecret, koreainvestment.KoreaInvestmentAccount{
		CANO:         account,
		ACNT_PRDT_CD: "01",
	})
	koreainvestment.DemoCallFunction()
}

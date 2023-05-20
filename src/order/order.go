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
	koreainvestment.Initialize(url, appKey, appSecret)
	koreainvestment.DemoCallFunction()
}

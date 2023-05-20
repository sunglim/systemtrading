package order

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/koreaexchange"
	"sunglim.github.com/sunglim/order/koreainvestment"
)

// Buy single stock every day at 10 am.

func order(apiOrderCash *koreainvestment.ApiOrderCash) {
	response := apiOrderCash.Call()

	handleResponse(response)
}

func StrategryBuyEveryDay() {
	err := initializeKoreaInvestment()
	if err != nil {
		fmt.Printf("initialization failed %s", err.Error())
	}

	apiOrderCash := koreainvestment.CreateApiOrderCash(koreaexchange.Code삼성전자)

	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At("10:00")
	s.Do(order, apiOrderCash)
	s.StartAsync()
}

func handleResponse(response *koreainvestment.ApiOrdeCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

func isSuccess(rtcd string) bool {
	return rtcd == "1"
}

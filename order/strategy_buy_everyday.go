package order

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
)

// Buy single stock every day at 10 am.

func orderOrderCash(apiOrderCash *koreainvestment.ApiOrderCash) {
	response := apiOrderCash.Call()

	handleOrderOrderCashResponse(response)
}

func StrategryBuyEveryDay(code, buytime string) {
	err := initializeKoreaInvestment()
	if err != nil {
		fmt.Printf("initialization failed %s", err.Error())
	}

	apiOrderCash := koreainvestment.CreateApiOrderCash(code)
	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(orderOrderCash, apiOrderCash)
	s.StartAsync()
}

func handleOrderOrderCashResponse(response *koreainvestment.ApiOrderCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

func isSuccess(rtcd string) bool {
	return rtcd == "1"
}

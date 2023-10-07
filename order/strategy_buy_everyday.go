package order

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sunglim/systemtrading/order/koreainvestment"
	ki "github.com/sunglim/systemtrading/pkg/koreainvestment"
)

// Buy single stock every day at 10 am.

func orderOrderCash(apiOrderCash *ki.ApiOrderCash) {
	response := apiOrderCash.Call()

	handleOrderOrderCashResponse(response)
}

func StrategryBuyEveryDay(code, buytime string) *gocron.Scheduler {
	apiOrderCash := ki.CreateApiOrderCash(code,
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())

	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(orderOrderCash, apiOrderCash)
	s.StartAsync()

	return s
}

func handleOrderOrderCashResponse(response *ki.ApiOrderCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

func isSuccess(rtcd string) bool {
	return rtcd == "1"
}

package order

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	krxcode "github.com/sunglim/go-korea-stock-code/code"
	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

// Buy single stock every day at 10 am.

func order(logger *log.Logger) {
	logger.Printf("Triggered")
	balanceResponse := koreainvestment.ApiInqueryBalance{}.Call()
	if !balanceResponse.IsSucess() {
		logger.Printf("Getting blance failed")
		return
	}

	for _, output := range balanceResponse.Output1 {
		orderCash(output, logger)
	}

}

func orderCash(balanceResponseOutput koreainvestment.ApiInqueryBalanceResponseOutput, logger *log.Logger) {
	// Core logic starts.
	plus_minus, _ := strconv.Atoi(balanceResponseOutput.EvluPflsAmt)
	if plus_minus > 0 {
		logger.Println("Didn't buy a stock;", "name", balanceResponseOutput.PrdtName,
			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
			"plus-minus", plus_minus)
		return
	}

	code := balanceResponseOutput.PdNo

	var api = ki.CreateApiOrderCash(code,
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	// Hack :(
	if code == krxcode.CodeDGB금융지주 {
		api = ki.NewApiOrderCash(code, 3,
			koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
			koreainvestment.GetDefaultAccount(),
			koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	}
	if code == krxcode.CodeBNK금융지주 {
		api = ki.NewApiOrderCash(code, 3,
			koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
			koreainvestment.GetDefaultAccount(),
			koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	}
	response := api.Call()
	handleResponse(response)
	if !response.IsSuccess() {
		logger.Printf("Getting Api order cash failed from the strategry")
		logger.Printf("Error[%s]", response.Msg1)
		return
	}

	logger.Println("An order is successfully sent", "name", balanceResponseOutput.PrdtName, "response", response.Msg1)
}

func StrategryBuyEveryDayIfBelowAverage(buytime string) {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is below] ")

	logger.Println("start new stragegy")

	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(order, logger)
	s.StartAsync()
}

func handleResponse(response *ki.ApiOrderCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

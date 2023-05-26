package order

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/log"
	"sunglim.github.com/sunglim/order/koreainvestment"
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
		logger.Printf("Didn't buy a stock;", "name", balanceResponseOutput.PrdtName,
			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
			"plus-minus", plus_minus)
		return
	}

	code := balanceResponseOutput.PdNo

	response := koreainvestment.CreateApiOrderCash(code).Call()
	handleResponse(response)
	if !response.IsSuccess() {
		logger.Printf("Getting Api order cash failed from the strategry")
		logger.Printf("Error[%s]", response.Msg1)
		return
	}

	logger.Printf("An order is successfully sent [%v]", response)
}

func StrategryBuyEveryDayIfBelowAverage(code, buytime string) {
	logger := log.Default()
	logger.SetPrefix("Buy stock if average is below")

	logger.Println("start new stragegy")

	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(order, logger)
	s.StartAsync()
}

func handleResponse(response *koreainvestment.ApiOrderCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

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
	logger.Printf("Buy stock below is triggered")
	apiInqueryBalance := koreainvestment.ApiInqueryBalance{}
	balanceResponse := apiInqueryBalance.Call()
	if balanceResponse.RtCd != "0" {
		logger.Printf("Getting blance failed from the strategry[buy every day if below average]")
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
		logger.Printf("Didn't buy a stock; plus minus is [%d]", plus_minus)
		return
	}

	code := balanceResponseOutput.PdNo

	apiOrderCash := koreainvestment.CreateApiOrderCash(code)
	response := apiOrderCash.Call()
	handleResponse(response)
	if response.RtCd != "0" {
		logger.Printf("Getting Api order cash failed from the strategry[buy every day if below average]")
		logger.Printf("Error message [%s]", response.Msg1)
		return
	}

	logger.Printf("Order successfully made[%v]", response)
}

func StrategryBuyEveryDayIfBelowAverage(code, buytime string) {
	log.Println("start new stragegy")
	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(order, log.Default())
	s.StartAsync()
}

func handleResponse(response *koreainvestment.ApiOrdeCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
	}
}

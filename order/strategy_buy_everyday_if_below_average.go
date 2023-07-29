package order

import (
	"fmt"
	"strconv"

	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

// Buy single stock every day at 10 am.

func order(codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	logger.Printf("Triggered")
	balanceResponse := koreainvestment.ApiInqueryBalance{}.Call()
	if !balanceResponse.IsSucess() {
		logger.Printf("Getting blance failed")
		return
	}

	for _, output := range balanceResponse.Output1 {
		orderCash(output, codeQuantity, logger)
	}

}

func getQuantityByCode(code string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) int {
	for _, single := range codeQuantity {
		if code == single.Code {
			return single.Quantity
		}
	}
	return 1
}

func orderCash(balanceResponseOutput koreainvestment.ApiInqueryBalanceResponseOutput, codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	// Core logic starts.
	plus_minus, _ := strconv.Atoi(balanceResponseOutput.EvluPflsAmt)
	if plus_minus > 0 {
		logger.Println("Didn't buy a stock;", "name", balanceResponseOutput.PrdtName,
			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
			"plus-minus", plus_minus)
		return
	}

	code := balanceResponseOutput.PdNo

	api := ki.NewApiOrderCash(code, getQuantityByCode(code, codeQuantity),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	response := api.Call()
	handleResponse(response)
	if !response.IsSuccess() {
		logger.Printf("Getting Api order cash failed from the strategry")
		logger.Printf("Error[%s]", response.Msg1)
		return
	}

	logger.Println("An order is successfully sent", "name", balanceResponseOutput.PrdtName, "response", response.Msg1)
}

type StrategryBuyEveryDayIfBelowOrder struct {
	Code     string
	Quantity int
}

func StrategryBuyEveryDayIfBelowAverage(buytime string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is below] ")

	logger.Println("start new stragegy")

	s := NewSeoulScheduler().Every(1).Day().At(buytime)
	s.Do(order, codeQuantity, logger)
	s.StartAsync()
}

func handleResponse(response *ki.ApiOrderCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
		return
	}
	fmt.Printf("Call fail. error code[%s], msg[%s], responseTime[%v]\n", response.RtCd, response.Msg1, response.ResponseTime)
}

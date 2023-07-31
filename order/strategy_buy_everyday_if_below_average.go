package order

import (
	"strconv"

	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

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

package order

import (
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
	logger.Println("quantity:" + balanceResponseOutput.HldgQty)
	quantity, _ := strconv.Atoi(balanceResponseOutput.HldgQty)
	if quantity <= 0 {
		return
	}

	plus_minus, _ := strconv.Atoi(balanceResponseOutput.EvluPflsAmt)
	if plus_minus > 0 {
		logger.Println("Didn't buy a stock;", "name", balanceResponseOutput.PrdtName,
			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
			"plus-minus", plus_minus)
		return
	}

	code := balanceResponseOutput.PdNo
	orderAmount := getQuantityByCode(code, codeQuantity)

	api := ki.NewApiOrderCash(code, orderAmount,
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
	Code     string `yaml:"code"`
	Quantity int    `yaml:"quantity"`
}

func StrategryBuyEveryDayIfBelowAverage(buytime string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is below] ")

	logger.Println("start new stragegy")

	s := NewSeoulScheduler().Every(1).Day().At(buytime)
	s.Do(order, codeQuantity, logger)
	s.StartAsync()
}

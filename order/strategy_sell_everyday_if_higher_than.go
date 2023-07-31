package order

import (
	"strconv"

	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
)

type StrategrySellEveryDayIfBelowOrder struct {
	Code     string
	Quantity int
}

func NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage(buytime string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) *StrategySellEveryDayIfAverageIsHigherThanAveragePercentage {
	return &StrategySellEveryDayIfAverageIsHigherThanAveragePercentage{buytime, codeQuantity}
}

type StrategySellEveryDayIfAverageIsHigherThanAveragePercentage struct {
	buyTime      string
	codeQuantity []StrategryBuyEveryDayIfBelowOrder
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) Start() {
	logger := log.Default()
	logger.SetPrefix("[Sell if average is higher] ")

	logger.Println("start new stragegy")

	s := NewSeoulScheduler().Every(1).Day().At(f.buyTime)
	s.Do(f.order, f.codeQuantity, logger)
	s.StartAsync()
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) order(codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	logger.Printf("Triggered")
	balanceResponse := koreainvestment.ApiInqueryBalance{}.Call()
	if !balanceResponse.IsSucess() {
		logger.Printf("Getting blance failed")
		return
	}

	for _, output := range balanceResponse.Output1 {
		f.orderCash(output, codeQuantity, logger)
	}

}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) orderCash(balanceResponseOutput koreainvestment.ApiInqueryBalanceResponseOutput, codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	// Core logic starts.
	// atoi X
	gain_percentage, _ := strconv.ParseFloat(balanceResponseOutput.EvluPflsRt, 32)
	if gain_percentage < 3.0 {
		logger.Println("Didn't sell a stock;", "name", balanceResponseOutput.PrdtName,
			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
			"gain_percentage", gain_percentage)
		return
	}

	logger.Println("Gain is above 3%", "name", balanceResponseOutput.PrdtName, "gain", gain_percentage)

	/* WIP
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
	*/
}

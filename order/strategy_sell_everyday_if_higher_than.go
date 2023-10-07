package order

import (
	"strconv"

	"github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

type StrategrySellEveryDayIfBelowOrder = CodeAndQuantity

func NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage(buytime string, codeQuantity []StrategrySellEveryDayIfBelowOrder) *StrategySellEveryDayIfAverageIsHigherThanAveragePercentage {
	return &StrategySellEveryDayIfAverageIsHigherThanAveragePercentage{buytime, codeQuantity}
}

type StrategySellEveryDayIfAverageIsHigherThanAveragePercentage struct {
	buyTime      string
	codeQuantity []StrategrySellEveryDayIfBelowOrder
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) Start() *gocron.Scheduler {
	s := NewSeoulScheduler().Every(1).Day().At(f.buyTime)
	s.Do(f.order, f.codeQuantity, log.Default())
	s.StartAsync()

	return s
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) order(codeQuantity []StrategrySellEveryDayIfBelowOrder, logger *log.Logger) {
	api := ki.NewApiInquireBalance(koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())

	balanceResponse, err := api.Call()
	if !balanceResponse.IsSucess() || err != nil {
		logger.Printf("Getting blance failed" + err.Error())
		return
	}

	for _, output := range balanceResponse.Output1 {
		f.orderCash(output, codeQuantity, logger)
	}

}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) orderCash(balanceResponseOutput ki.ApiInquireBalanceResponseOutput, codeQuantity []StrategrySellEveryDayIfBelowOrder, logger *log.Logger) {
	// Core logic starts.
	// atoi X
	gain_percentage, _ := strconv.ParseFloat(balanceResponseOutput.EvluPflsRt, 32)
	if gain_percentage < 3.0 {
		//#logger.Println("Didn't sell a stock;", "name", balanceResponseOutput.PrdtName,
		//			"current price", balanceResponseOutput.Prpr, "average", balanceResponseOutput.PchsAvgPric,
		//			"gain_percentage", gain_percentage)
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

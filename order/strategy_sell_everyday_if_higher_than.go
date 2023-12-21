package order

import (
	"strconv"

	"github.com/go-co-op/gocron"
	"github.com/sunglim/systemtrading/log"
	"github.com/sunglim/systemtrading/order/koreainvestment"
	ki "github.com/sunglim/systemtrading/pkg/koreainvestment"
)

type StrategrySellEveryDayIfBelowOrder = CodeAndQuantityAndPrice

func NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage(buytime string, codeQuantityPrice []StrategrySellEveryDayIfBelowOrder) *StrategySellEveryDayIfAverageIsHigherThanAveragePercentage {
	return &StrategySellEveryDayIfAverageIsHigherThanAveragePercentage{buytime, codeQuantityPrice}
}

type StrategySellEveryDayIfAverageIsHigherThanAveragePercentage struct {
	buyTime           string
	codeQuantityPrice []StrategrySellEveryDayIfBelowOrder
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) Start() *gocron.Scheduler {
	s := NewSeoulScheduler().Every(1).Day().At(f.buyTime)
	logger := log.Default().With("name", "StrategySellEveryDayIfAverageIsHigherThanAveragePercentage")
	s.Do(f.order, f.codeQuantityPrice, logger)
	s.StartAsync()

	return s
}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) order(codeQuantityPrice []StrategrySellEveryDayIfBelowOrder, logger *log.Logger) {
	logger.Info("triggered")

	api := ki.NewApiInquireBalance(koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())

	balanceResponse, err := api.Call()
	if (balanceResponse != nil && !balanceResponse.IsSucess()) || err != nil {
		logger.Error("Getting blance failed" + err.Error())
		return
	}

	for _, output := range balanceResponse.Output1 {
		// filter with codequantity
		f.orderCash(output, codeQuantityPrice, logger)
	}

}

func (f StrategySellEveryDayIfAverageIsHigherThanAveragePercentage) orderCash(balanceResponseOutput ki.ApiInquireBalanceResponseOutput, codeQuantity []StrategrySellEveryDayIfBelowOrder, logger *log.Logger) {
	gain_percentage, _ := strconv.ParseFloat(balanceResponseOutput.EvluPflsRt, 32)
	// DO NOT SELL IF GAIN IS LESS THAN 3%!
	if gain_percentage < 3.0 {
		return
	}

	logger.Info("Gain is above 3%", "name", balanceResponseOutput.PrdtName, "gain", gain_percentage)

	//code := balanceResponseOutput.PdNo
	//numbersToSell := getQuantityByCode(code, codeQuantity)

	/* WIP


	api := ki.NewApiOrderCash(code, getQuantityByCode(code, codeQuantity),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	response := api.Call()
	if !response.IsSuccess() {
		logger.Printf("Getting Api sell cash failed from the strategry")
		logger.Printf("Error[%s]", response.Msg1)
		return
	}

	logger.Info("Sell order is successfully sent", "name", balanceResponseOutput.PrdtName, "response", response.Msg1)
	*/
}

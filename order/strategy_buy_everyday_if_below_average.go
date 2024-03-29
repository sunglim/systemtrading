package order

import (
	"strconv"

	"github.com/go-co-op/gocron"
	"github.com/sunglim/systemtrading/log"
	"github.com/sunglim/systemtrading/order/koreainvestment"
	ki "github.com/sunglim/systemtrading/pkg/koreainvestment"
	"k8s.io/apimachinery/pkg/util/sets"
)

func order(codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	logger.Info("StrategryBuyEveryDayIfBelowOrder is triggered")

	api := ki.NewApiInquireBalance(koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())

	balanceResponse, err := api.Call()
	if (balanceResponse != nil && !balanceResponse.IsSucess()) || err != nil {
		logger.Error("Getting blance failed" + err.Error())
		return
	}

	for _, output := range balanceResponse.Output1 {
		orderCash(output, codeQuantity, logger)
	}

}

func orderCash(balanceResponseOutput ki.ApiInquireBalanceResponseOutput, codeQuantity []StrategryBuyEveryDayIfBelowOrder, logger *log.Logger) {
	// Core logic starts.
	quantity, _ := strconv.Atoi(balanceResponseOutput.HldgQty)
	if quantity <= 0 {
		return
	}

	plus_minus, _ := strconv.Atoi(balanceResponseOutput.EvluPflsAmt)
	if plus_minus > 0 {
		return
	}

	code := balanceResponseOutput.PdNo
	orderQuantity := getQuantityByCode(code, codeQuantity)

	api := ki.NewApiOrderCash(code, orderQuantity,
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken())
	response := api.Call()
	if !response.IsSuccess() && !IsAllowedError(response.MsgCd) {
		logger.Error("api order cash failed", "code", code, "quantity", orderQuantity, "msg", response.Msg1, "msgcode", response.MsgCd)
		return
	}

	logger.Info("BUY", "name", balanceResponseOutput.PrdtName)
}

type StrategryBuyEveryDayIfBelowOrder = CodeAndQuantity

func StrategryBuyEveryDayIfBelowAverage(buytime string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) *gocron.Scheduler {
	logger := log.Default().With("name", "StrategryBuyEveryDayIfBelowAverage")

	s := NewSeoulScheduler().Every(1).Day().At(buytime)
	s.Do(order, codeQuantity, logger)
	s.StartAsync()

	return s
}

func IsAllowedError(errorCode ki.ErrorCode) bool {
	return sets.New(ki.OrderQuantityErrorCode, ki.OrderClosingTimeErrorCode).Has(errorCode)
}

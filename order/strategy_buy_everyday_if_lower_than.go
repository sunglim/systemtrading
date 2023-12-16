package order

import (
	"strconv"

	"github.com/go-co-op/gocron"
	krxcode "github.com/sunglim/go-korea-stock-code/code"
	"github.com/sunglim/systemtrading/log"
	"github.com/sunglim/systemtrading/order/koreainvestment"
	ki "github.com/sunglim/systemtrading/pkg/koreainvestment"
)

// Buy a stock if the price is lower than ...

func buyLowerOrder(codePrices []StrategryOrder, logger *log.Logger) {
	for _, codePrice := range codePrices {
		currentPrice := koreainvestment.ApiInqueryPrice{}.Call(codePrice.Code)
		currentPriceInt, _ := strconv.Atoi(currentPrice)
		if currentPriceInt > codePrice.Price {
			continue
		}

		logger.Info("buy lower", "name", krxcode.CodeToName(codePrice.Code), "orderPrice",
			codePrice.Price, "currentPrice", currentPriceInt)
		BuyLowerOrderCash(codePrice, logger)
	}

}

func BuyLowerOrderCash(code StrategryOrder, logger *log.Logger) {
	response := ki.CreateApiOrderCash(code.Code,
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetCredential(),
		koreainvestment.GetDefaultAccount(),
		koreainvestment.GetDefaultKoreaInvestmentInstance().GetBearerAccessToken()).Call()
	if response != nil && !response.IsSuccess() {
		logger.Printf("orde failed with error[%s]", response.Msg1)
		return
	}

	logger.Info("An order is successfully sent", "response", response)
}

type StrategryOrder = CodeAndQuantityAndPrice

func StrategryBuyEveryDayIfLowerThan(buytime string, codePrices []StrategryOrder) *gocron.Scheduler {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is lower than] ")

	s := NewSeoulScheduler().Every(1).Day().At(buytime)
	s.Do(buyLowerOrder, codePrices, logger)
	s.StartAsync()

	return s
}

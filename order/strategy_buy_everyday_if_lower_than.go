package order

import (
	"strconv"

	krxcode "github.com/sunglim/go-korea-stock-code/code"
	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
)

// Buy a stock if the price is lower than ...

func buyLowerOrder(codePrices []StrategryOrder, logger *log.Logger) {
	for _, codePrice := range codePrices {
		currentPrice := koreainvestment.ApiInqueryPrice{}.Call(codePrice.Code)
		currentPriceInt, _ := strconv.Atoi(currentPrice)
		if currentPriceInt > codePrice.Price {
			continue
		}

		logger.Println("name", krxcode.CodeToName(codePrice.Code), "orderPrice",
			codePrice.Price, "currentPrice", currentPriceInt)
		BuyLowerOrderCash(codePrice, logger)
	}

}

func BuyLowerOrderCash(code StrategryOrder, logger *log.Logger) {
	response := koreainvestment.CreateApiOrderCash(code.Code).Call()
	handleResponse(response)
	if !response.IsSuccess() {
		logger.Printf("orde failed with error[%s]", response.Msg1)
		return
	}

	logger.Printf("An order is successfully sent [%v]", response)
}

type StrategryOrder struct {
	Code     string
	Price    int
	Quantity int
}

func StrategryBuyEveryDayIfLowerThan(buytime string, codePrices []StrategryOrder) {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is lower than] ")

	logger.Println("start new stragegy")

	s := NewSeoulScheduler().Every(1).Day().At(buytime)
	s.Do(buyLowerOrder, codePrices, logger)
	s.StartAsync()
}

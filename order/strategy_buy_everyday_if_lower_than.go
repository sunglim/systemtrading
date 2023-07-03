package order

import (
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
)

// Buy a stock if the price is lower than ...

func buyLowerOrder(codePrices []CodePrice, logger *log.Logger) {
	for _, codePrice := range codePrices {
		currentPrice := koreainvestment.ApiInqueryPrice{}.Call(codePrice.Code)
		priceInt, _ := strconv.Atoi(currentPrice)
		if priceInt > codePrice.Price {
			continue
		}
		logger.Println("buy lower than is triggerd", "code", codePrice.Code, "orderprice", priceInt)
		BuyLowerOrderCash(codePrice.Code, logger)
	}

}

func BuyLowerOrderCash(code string, logger *log.Logger) {
	response := koreainvestment.CreateApiOrderCash(code).Call()
	handleResponse(response)
	if !response.IsSuccess() {
		logger.Printf("Getting Api order cash failed from the strategry")
		logger.Printf("Error[%s]", response.Msg1)
		return
	}

	logger.Printf("An order is successfully sent [%v]", response)
}

type CodePrice struct {
	Code  string
	Price int
}

func StrategryBuyEveryDayIfLowerThan(buytime string, codePrices []CodePrice) {
	logger := log.Default()
	logger.SetPrefix("[Buy if average is lower than] ")

	logger.Println("start new stragegy")

	s := gocron.NewScheduler(time.Now().Location()).Every(1).Day().At(buytime)
	s.Do(buyLowerOrder, codePrices, logger)
	s.StartAsync()
}

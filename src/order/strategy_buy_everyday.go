package order

import (
	"fmt"

	"sunglim.github.com/sunglim/koreaexchange"
	"sunglim.github.com/sunglim/order/koreainvestment"
)

// Buy single stock every day at 10 am.

func StrategryBuyEveryDay() error {
	err := initializeKoreaInvestment()
	if err != nil {
		return fmt.Errorf("initialization failed %s", err.Error())
	}

	apiOrderCash := koreainvestment.CreateApiOrderCash(koreaexchange.Code삼성전자)

	response := apiOrderCash.Call()

	handleResponse(response)

	return nil
}

func handleResponse(response *koreainvestment.ApiOrdeCashResponse) {
	if isSuccess(response.RtCd) {
		fmt.Printf("Call success\n")
	} else {
		fmt.Printf("Call fail. error code[%s]\n", response.RtCd)
	}
}

func isSuccess(rtcd string) bool {
	return rtcd == "1"
}

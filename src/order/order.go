package order

import (
	"sunglim.github.com/sunglim/history"
)

type Order struct {
}

func Demo() {
	/*
		balance := koreainvestment.ApiInqueryDailyItemChartPrice{}
		response := balance.Call(koreaexchange.Code삼성전자)
		fmt.Printf("%v", response)
	*/

	history.GetHistoricalData()
}

package order

import (
	"fmt"

	"sunglim.github.com/sunglim/order/koreainvestment"
)

type Order struct {
}

func Demo() {
	balance := koreainvestment.ApiInqueryBalance{}
	response := balance.Call()
	fmt.Printf(response.Msg1)
}

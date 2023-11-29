package order

func getQuantityByCode(code string, codeQuantity []StrategryBuyEveryDayIfBelowOrder) int {
	for _, single := range codeQuantity {
		if code == single.Code {
			return single.Quantity
		}
	}

	return 1
}

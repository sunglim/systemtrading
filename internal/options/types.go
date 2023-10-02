package options

import "sunglim.github.com/sunglim/systemtrading/order"

type MetricSet map[string]struct{}

type BuyEveryDayIfBelowAverage struct {
	BuyEveryDayIfBelowAverage BuyEveryDayIfBelowAverageConfig `yaml:"BuyEveryDayIfBelowAverage"`
}

type BuyEveryDayIfBelowAverageConfig struct {
	// "12:00"
	ExecutionTime   string                                   `yaml:"ExecutionTime"`
	CodeAndQuantity []order.StrategryBuyEveryDayIfBelowOrder `yaml:"CodeAndQuantity"`
}

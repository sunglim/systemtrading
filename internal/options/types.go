package options

import "sunglim.github.com/sunglim/systemtrading/order"

type MetricSet map[string]struct{}

type BuyEveryDayIfBelowAverageConfig struct {
	BelowAverage []order.StrategryBuyEveryDayIfBelowOrder `yaml:"BuyEveryDayIfBelowAverage"`
}

package options

import "github.com/sunglim/systemtrading/order"

type MetricSet map[string]struct{}

type BuyEveryDayIfBelowAverage struct {
	BuyEveryDayIfBelowAverage BuyEveryDayIfBelowAverageConfig `yaml:"BuyEveryDayIfBelowAverage"`
}

type BuyEveryDayIfBelowAverageConfig struct {
	// "12:00"
	ExecutionTime   string                  `yaml:"ExecutionTime"`
	CodeAndQuantity []order.CodeAndQuantity `yaml:"CodeAndQuantity"`
}

type BuyEveryDayIfLowerThan struct {
	BuyEveryDayIfLowerThan BuyEveryDayIfLowerThanConfig `yaml:"BuyEveryDayIfLowerThan"`
}

type BuyEveryDayIfLowerThanConfig struct {
	ExecutionTime           string                 `yaml:"ExecutionTime"`
	CodeAndQuantityAndPrice []order.StrategryOrder `yaml:"CodeAndQuantityAndPrice"`
}

type SellEveryDayIfHigherThan struct {
	SellEveryDayIfLowerThan SellEveryDayIfHigherThanConfig `yaml:"SellEveryDayIfHigherThan"`
}

type SellEveryDayIfHigherThanConfig struct {
	ExecutionTime           string                 `yaml:"ExecutionTime"`
	CodeAndQuantityAndPrice []order.StrategryOrder `yaml:"CodeAndQuantityAndPrice"`
}

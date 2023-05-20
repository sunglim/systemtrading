package main

import (
	"sunglim.github.com/sunglim/order"
)

var exit = make(chan bool)

func main() {
	go order.StrategryBuyEveryDay()

	// Infinite.
	<-exit
}

package order

import (
	"os"

	gocron "github.com/go-co-op/gocron"
	sunglimlog "sunglim.github.com/sunglim/log"
	"sunglim.github.com/sunglim/order/koreainvestment"
)

type Order struct {
	logger *sunglimlog.Logger
}

func getFrequency(s *gocron.Scheduler) *gocron.Scheduler {
	if len(os.Args) > 2 && os.Args[1] == "Production" {
		return s.Every(1).Day().At("10:30")
	}

	return s.Every(10).Seconds()
}

func (order Order) DailyOrder() {
	url := os.Args[1]
	appKey := os.Args[2]
	appSecret := os.Args[3]
	koreainvestment.Initialize(url, appKey, appSecret)
	koreainvestment.DemoCallFunction()

	//	s := gocron.NewScheduler(time.UTC)
	//getFrequency(s).Do(KiwoomOrder)
	//s.StartBlocking()
}

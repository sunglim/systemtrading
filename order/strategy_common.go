package order

import (
	"time"

	"github.com/go-co-op/gocron"
)

func NewSeoulScheduler() *gocron.Scheduler {
	location, _ := time.LoadLocation("Asia/Seoul")
	return gocron.NewScheduler(location)
}

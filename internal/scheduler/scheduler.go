package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func GetScheduler() *gocron.Scheduler {
	s := gocron.NewScheduler(time.Now().Location())
	return s
}

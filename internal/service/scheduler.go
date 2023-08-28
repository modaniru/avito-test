package service

import (
	"fmt"
	"time"

	log "log/slog"

	"github.com/modaniru/avito/internal/storage"
)

type Scheduler struct {
	followStorage storage.Follow
}

func NewScheduler(followStorage storage.Follow) *Scheduler {
	return &Scheduler{followStorage: followStorage}
}

func (s *Scheduler) RunScheduler() chan int {
	channel := make(chan int)
	go func() {
		for true {
			count, err := s.followStorage.DeleteExpiredFollows()
			log.Info(fmt.Sprintf("scheduler delete %d rows", count))
			if err != nil {
				log.Error("scheduler error", log.String("error", err.Error()))
				channel <- -1
				return
			}
			time.Sleep(time.Minute)
		}
	}()
	return channel
}

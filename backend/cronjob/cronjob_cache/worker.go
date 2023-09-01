package cronjob_cache

import (
	"context"
	"github.com/jasonlvhit/gocron"
	"thichlab-backend-docs/cronjob/utils"
	"thichlab-backend-docs/infrastructure/cache"
	"time"
)

type Cache struct {
	CacheClient cache.Client
}

func Run(env *utils.CronjobEnvironment) error {
	running := false
	for range time.NewTicker(time.Minute).C {
		_, cancel := context.WithTimeout(context.Background(), 15*time.Second)

		if !running {
			running = true

			// Start all the pending jobs
			<-gocron.Start()

		}
		cancel()

	}
	return nil

}

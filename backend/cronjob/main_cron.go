package cronjob

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"thichlab-backend-docs/cronjob/utils"
)

func initEnvironment() (*utils.CronjobEnvironment, error) {
	env := &utils.CronjobEnvironment{
		CacheRedis: redis.NewClient(&redis.Options{
			Addr:         viper.GetString("Redis.Host"),
			Password:     viper.GetString("Redis.Password"),
			PoolSize:     viper.GetInt("Redis.PoolSize"),
			MinIdleConns: viper.GetInt("Redis.MinIdleConns"),
		}),
	}
	return env, nil
}

// Run Start Cron jobs
func Run() {
	env, err := initEnvironment()
	if env == nil || err != nil {
		err := fmt.Errorf("[Cron] init environment for cronjob failed")
		panic(err)
	}

	//defer func() {
	//}()

	process := &utils.Processes{}
	//process.Add("user-one-day", getuser1day.Run, env)
	//process.Add("user-five-min", getuser5min.Run, env)
	process.Run()
}

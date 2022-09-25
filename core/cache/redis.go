package cache

import (
	"enterprise-api/app/config"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	conf := config.GetConfig().RedisConfig
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         conf.Host + conf.Port,
		Password:     conf.Password,
		DB:           conf.Db,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pong, err := RedisClient.Ping().Result()
	fmt.Println(pong, err)
}

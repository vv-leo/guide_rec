package model

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var redisClient *redis.Client

func InitRedis() {
	redisAddress := viper.GetString("redis.address")
	redisPassword := viper.GetString("redis.password")
	redisDb := viper.GetInt("redis.db")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       redisDb,
	})
}

func SetRedisValue(key string, value interface{}, expiration int) {
	redisClient.Set(key, value, time.Duration(expiration))
}

func GetRedisValue(key string) string {
	value, err := redisClient.Get(key).Result()
	if err != nil {
		return ""
	}
	return value
}

func DelRedisKey(key string) {
	redisClient.Del(key)
}

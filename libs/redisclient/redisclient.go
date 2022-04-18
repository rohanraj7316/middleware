package redisclient

import (
	"github.com/go-redis/redis/v8"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

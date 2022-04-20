package redisclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func New(config Config) (*redis.Client, error) {
	if config.HostName == "" || config.Port == "" {
		return nil, errors.New("inavlid hostname and password")
	}
	address_url := fmt.Sprintf("%s%s", config.HostName, config.Password)

	newClient := redis.NewClient(&redis.Options{
		Addr:     address_url,
		Password: config.Password,
		DB:       config.DB,
	})

	if _, redis_err := newClient.Ping(context.Background()).Result(); redis_err != nil {
		return nil, errors.New("unable to connect to redis")
	}
	return newClient, nil
}

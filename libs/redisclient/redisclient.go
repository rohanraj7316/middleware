package redisclient

import (
	"github.com/go-redis/redis/v8"
	"github.com/rohanraj7316/logger"
)

// New responsible for creating new redis client
func New(config ...Config) (Config, error) {
	err := logger.Configure()
	if err != nil {
		return Config{}, err
	}

	cfg, err := configDefault(config...)
	if err != nil {
		return cfg, err
	}

	cfg.Client = redis.NewClient(cfg.Redis)

	return cfg, nil
}

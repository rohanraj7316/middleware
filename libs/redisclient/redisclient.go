package redisclient

import (
	"github.com/go-redis/redis/v8"
)

// New responsible for creating new redis client
func New(config ...Config) (Config, error) {
	cfg, err := configDefault(config...)
	if err != nil {
		return cfg, err
	}

	cfg.Client = redis.NewClient(cfg.Redis)

	return cfg, nil
}

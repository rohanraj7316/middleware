package redisclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rohanraj7316/logger"
)

// GetValue retrieves the value of environment variable named by the key.
// It returns the value, if not then it will return the defult value passed
// as the second parameter.
func GetValue(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	logger.Error(fmt.Sprintf("[REDISError] empty key: %s", k))
	return d
}

type Config struct {
	Redis  *redis.Options
	Client *redis.Client
}

var ConfigDefault = Config{
	Redis: &redis.Options{
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			return nil
		},
		TLSConfig: &tls.Config{},
	},
}

func configDefault(config ...Config) (Config, error) {
	ConfigDefault.Redis.Password = GetValue("REDIS_AUTH", "")
	ConfigDefault.Redis.Addr = fmt.Sprintf("%s:%s", GetValue("REDIS_HOST", ""), GetValue("REDIS_PORT", ""))

	mRetries, err := strconv.Atoi(GetValue("REDIS_MAX_RETRIES", "5"))
	if err != nil {
		return Config{}, err
	}
	ConfigDefault.Redis.MaxRetries = mRetries

	pSize, err := strconv.Atoi(GetValue("REDIS_POOL_SIZE", "10"))
	if err != nil {
		return Config{}, err
	}
	ConfigDefault.Redis.PoolSize = pSize

	if len(config) < 1 {
		return ConfigDefault, nil
	}

	cfg := config[0]

	return cfg, nil
}

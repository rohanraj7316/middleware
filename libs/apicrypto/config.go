package apicrypto

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rohanraj7316/middleware/libs/redisclient"
	"github.com/rohanraj7316/middleware/libs/utils"
)

const (
	ENV_API_CRYPTO_REDIS_HOST                   = "API_CRYPTO_REDIS_HOST"
	ENV_API_CRYPTO_REDIS_PORT                   = "API_CRYPTO_REDIS_PORT"
	ENV_API_CRYPTO_REDIS_AUTH                   = "API_CRYPTO_REDIS_AUTH"
	ENV_API_CRYPTO_REDIS_PREFIX                 = "API_CRYPTO_REDIS_PREFIX"
	ENV_API_CRYPTO_REDIS_POOL_SIZE              = "API_CRYPTO_REDIS_POOL_SIZE"
	ENV_API_CRYPTO_REDIS_MAX_RETRIES            = "API_CRYPTO_REDIS_MAX_RETRIES"
	ENV_API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED = "API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED"
)

type ValidClient func(ctx context.Context, clientId string) bool

type Config struct {
	isCryptoEnable bool
	redisConfig    redisclient.Config

	ValidClientFunc ValidClient
}

var ConfigDefault = Config{}

func configDefault(config ...Config) Config {
	// extracting configs from env
	rConfigs := []string{
		ENV_API_CRYPTO_REDIS_HOST,
		ENV_API_CRYPTO_REDIS_PORT,
		ENV_API_CRYPTO_REDIS_AUTH,
		ENV_API_CRYPTO_REDIS_PREFIX,
	}
	uConfigs := []string{
		ENV_API_CRYPTO_REDIS_POOL_SIZE,
		ENV_API_CRYPTO_REDIS_MAX_RETRIES,
		ENV_API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED,
	}
	eConfig := utils.EnvData(rConfigs, uConfigs)

	if eConfig[ENV_API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED] != "" {
		isCryptoEnable, err := strconv.ParseBool(eConfig[ENV_API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED])
		if err != nil {
			panic(fmt.Sprintf("failed to parse %s with err: %s", ENV_API_CRYPTO_PAYLOAD_CRYPTOGRAPHY_ENABLED, err))
		}
		ConfigDefault.isCryptoEnable = isCryptoEnable
	} else {
		ConfigDefault.isCryptoEnable = false
	}

	ConfigDefault.redisConfig = redisclient.Config{
		Redis: &redis.Options{
			Addr:     fmt.Sprintf("%s:%s", eConfig[ENV_API_CRYPTO_REDIS_HOST], eConfig[ENV_API_CRYPTO_REDIS_PORT]),
			Password: eConfig[ENV_API_CRYPTO_REDIS_AUTH],
		},
		Prefix: eConfig[ENV_API_CRYPTO_REDIS_PREFIX],
	}

	if len(config) < 1 {
		panic("`ValidClientFunc` is missing")
	}

	cfg := config[0]

	ConfigDefault.ValidClientFunc = cfg.ValidClientFunc

	return ConfigDefault
}

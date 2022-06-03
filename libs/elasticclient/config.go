package elasticclient

import (
	"fmt"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rohanraj7316/logger"
)

func GetValue(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	logger.Error(fmt.Sprintf("[Elasticerror] empty key: %s", k))
	return d
}

type Config struct {
	ElasticSearchConfig *elasticsearch.Config
	Client              *elasticsearch.Client
}

var ConfigDefault = Config{
	ElasticSearchConfig: &elasticsearch.Config{},
}

func configDefault(config ...Config) (Config, error) {
	mRetries, err := strconv.Atoi(GetValue("ELASTIC_MAX_RETRIES", ""))
	if err != nil {
		return Config{}, err
	}
	ConfigDefault.ElasticSearchConfig.MaxRetries = mRetries

	//TODO connection pool logic

	if len(config) < 1 {
		return ConfigDefault, fmt.Errorf("elastic config length is less than one")
	}

	cfg := config[0]

	return cfg, nil
}

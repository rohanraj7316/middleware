package elasticclient

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rohanraj7316/logger"
)

func New(config ...Config) (Config, error) {
	err := logger.Configure()
	if err != nil {
		return Config{}, err
	}

	cfg, err := configDefault(config...)
	if err != nil {
		return cfg, nil
	}

	cfg.Client, err = elasticsearch.NewClient(*cfg.ElasticSearchConfig)

	return cfg, nil
}

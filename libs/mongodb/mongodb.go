package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(ctx context.Context, config ...Config) (cfg Config, err error) {
	cfg, err = configDefault()
	if err != nil {
		return cfg, err
	}

	client, err = mongo.Connect(ctx, cfg.cOptions)

	if err != nil {
		return cfg, err
	}

	return Config{
		Client: client,
		Db:     client.Database(cfg.DbName, cfg.dbOptions),
	}, nil
}

package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func New(ctx context.Context, config ...Config) (handler Handler, err error) {
	handler = Handler{}

	cfg, err := configDefault()
	if err != nil {
		return handler, err
	}

	client, err := mongo.Connect(ctx, cfg.cOptions)
	if err != nil {
		return handler, err
	}
	handler.Client = client

	if cfg.DbName != "" {
		handler.DB = client.Database(cfg.DbName, cfg.dOptions)
	}

	return handler, nil
}

// func (h Handler) Find(ctx context.Context, filter interface{}) {}

// func (h Handler) FindById() {}

// func (h Handler) Create() {}

// func (h Handler) DeleteById() {}

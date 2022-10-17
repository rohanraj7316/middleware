package mongodb

import (
	"context"
	"testing"

	env "github.com/joho/godotenv"
)

const (
	envFile = ".env"
)

var loadEnv = env.Load

func TestMongoDbConnection(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	_, err = New(ctx)
	if err != nil {
		t.Errorf("failed to connect to mongodb: %s", err)
	}
}

func TestMultipleMongoDbConnection(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	_, err = New(ctx)
	if err != nil {
		t.Errorf("failed to connect to mongodb: %s", err)
	}

	_, err = New(ctx)
	if err != nil {
		t.Errorf("failed to connect to mongodb: %s", err)
	}
}

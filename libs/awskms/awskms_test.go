package awskms_test

import (
	"context"
	"testing"

	env "github.com/joho/godotenv"
	"github.com/rohanraj7316/middleware/libs/awskms"
)

const (
	envFile = ".env"
)

var loadEnv = env.Load

func TestGenerateKMSKey(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	client, err := awskms.New()
	if err != nil {
		t.Errorf("failed to initalize kms: %s", err)
	}

	_, err = client.GenerateKMSKey(ctx, 32)
	if err != nil {
		t.Errorf("failed to generate rsa key: %s", err)
	}
}

// TODO: need to test this logic first
func TestGenerateRsaKeyPair(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	client, err := awskms.New()
	if err != nil {
		t.Errorf("failed to initalize kms: %s", err)
	}

	_, err = client.GenerateRsaKeyPair(ctx)
	if err != nil {
		t.Errorf("failed to generate rsa key: %s", err)
	}
}

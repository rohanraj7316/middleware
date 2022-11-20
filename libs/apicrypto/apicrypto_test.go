package apicrypto_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	env "github.com/joho/godotenv"
	"github.com/rohanraj7316/middleware/libs/apicrypto"
)

const (
	envFile = ".env"
)

var loadEnv = env.Load

func getPayload() (string, interface{}) {
	key := "ewd$#128djdyAgbjau&YAnmcbagryt5x"
	payload := map[string]interface{}{
		"name": "test",
		"time": time.Now(),
	}

	return key, payload
}

func TestHandshake(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	u, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %s", err)
	}

	cfg := apicrypto.Config{
		ValidClientFunc: func(ctx context.Context, clientId string) bool {
			return true
		},
	}
	handler, err := apicrypto.New(cfg)
	if err != nil {
		t.Errorf("failed to initalize handler: %s", err)
	}

	res := handler.Handshake(ctx, u.String())
	if res.StatusCode != 200 {
		t.Errorf("failed handshake response: %s", res.Err)
	}
}

func TestEncryptPayload(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	cfg := apicrypto.Config{
		ValidClientFunc: func(ctx context.Context, clientId string) bool {
			return true
		},
	}
	handler, err := apicrypto.New(cfg)
	if err != nil {
		t.Errorf("failed to initalize handler: %s", err)
	}

	plainSecretKey, payload := getPayload()
	payloadByte, _ := json.Marshal(payload)

	res := handler.EncryptPayload(ctx, plainSecretKey, string(payloadByte))
	if res.StatusCode != 200 {
		t.Errorf("failed to encrypt payload with status code: %d", res.StatusCode)
		t.Logf("%+v", res)
	}
}

func TestDecryptPayload(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	cfg := apicrypto.Config{
		ValidClientFunc: func(ctx context.Context, clientId string) bool {
			return true
		},
	}
	handler, err := apicrypto.New(cfg)
	if err != nil {
		t.Errorf("failed to initalize handler: %s", err)
	}

	plainSecretKey, payload := getPayload()
	payloadByte, _ := json.Marshal(payload)

	res := handler.EncryptPayload(ctx, plainSecretKey, string(payloadByte))
	if res.StatusCode != 200 {
		t.Errorf("failed to encrypt payload with status code: %d", res.StatusCode)
		t.Logf("%+v", res)
	}

	encryptedPayload, _ := res.Data.(map[string]string)
	res = handler.DecryptPayload(ctx, plainSecretKey, encryptedPayload["payload"])
	if res.StatusCode != 200 {
		t.Errorf("failed to decrypt payload with status code: %d", res.StatusCode)
		t.Logf("%+v", res)
	}
}

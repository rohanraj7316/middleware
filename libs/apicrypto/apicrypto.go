package apicrypto

import (
	"context"
	"encoding/base64"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/middleware/constants"
	"github.com/rohanraj7316/middleware/libs/apicrypto/helpers"
	"github.com/rohanraj7316/middleware/libs/awskms"
	"github.com/rohanraj7316/middleware/libs/redisclient"
	"github.com/rohanraj7316/middleware/libs/response"
)

type Handler struct {
	validClient ValidClient

	kmsHandler   awskms.Handler
	cacheHandler redisclient.Handler
}

var internalHandler Handler

// initializing as a lib
func New(config ...Config) (handler Handler, err error) {
	handler = Handler{}
	cfg := configDefault(config...)

	redisHandler, err := redisclient.New(cfg.redisConfig)
	if err != nil {
		return handler, err
	}
	handler.cacheHandler = redisHandler

	kmsHandler, err := awskms.New()
	if err != nil {
		return handler, err
	}
	handler.kmsHandler = kmsHandler

	handler.validClient = cfg.ValidClientFunc

	return handler, nil
}

// initialize as new middleware
func NewMiddleware(config ...Config) fiber.Handler {
	var (
		once sync.Once
	)

	once.Do(func() {
		cfg := configDefault(config...)
		redisHandler, err := redisclient.New(cfg.redisConfig)
		if err != nil {
		}
		internalHandler.cacheHandler = redisHandler

		kmsHandler, err := awskms.New()
		if err != nil {
		}
		internalHandler.kmsHandler = kmsHandler

		internalHandler.validClient = cfg.ValidClientFunc
	})

	return func(c *fiber.Ctx) error {
		clientId := c.Get(constants.REQUEST_CRYPTO_CLIENT_ID_HEADER_KEY)
		encryptionSecretKey := c.Get(constants.REQUEST_CRYPTO_ENCRYPTION_KEY_HEADER_KEY)

		res := internalHandler.DecryptKey(c.UserContext(), clientId, encryptionSecretKey)
		if res.StatusCode != 200 {
			// send back the response
			return response.NewBody(c, res.StatusCode, res.Message, res.Data, res.Err)
		}

		out := map[string]string{}
		err := c.BodyParser(&out)
		if err != nil {

		}

		encryptedPayload := out["payload"]
		res = internalHandler.DecryptPayload(c.UserContext(), encryptionSecretKey, encryptedPayload)
		if res.StatusCode != 200 {

		}

		err = c.Next()
		if err != nil {

		}

		// internalHandler.EncryptPayload(c.UserContext(), encryptionSecretKey, string(rByte))
		return nil
	}
}

func (h Handler) DecryptKey(ctx context.Context, clientId, encryptedSecretKey string) response.BodyStruct {
	isValidClient := h.validClient(ctx, clientId)
	if !isValidClient {
		return DecryptKeyErr("InvalidClient", nil)
	}

	privateKeyCacheId := helpers.GetPrivateKeyId(clientId)
	privateKeyCache, err := h.cacheHandler.Get(ctx, privateKeyCacheId).Result()
	if err != nil {
		if err == redis.Nil {
			out, err := h.kmsHandler.GenerateRsaKeyPair(ctx)
			if err != nil {
				return DecryptKeyErr("FailedToGenerateKeyPairs", err)
			}

			privateKeyCache = base64.StdEncoding.EncodeToString(out.PrivateKeyPlaintext)
			if err := h.cacheHandler.Set(ctx, privateKeyCacheId, privateKeyCache, 0).Err(); err != nil {
				return DecryptKeyErr("FailedToUpdatePrivateKeyCache", err)
			}

			publicKeyCacheId := helpers.GetPublicKeyId(clientId)
			publicKeyCache := base64.StdEncoding.EncodeToString(out.PublicKey)
			if err := h.cacheHandler.Set(ctx, publicKeyCacheId, publicKeyCache, 0).Err(); err != nil {
				return DecryptKeyErr("FailedToUpdatePublicKeyCache", err)
			}

			return DecryptKeySuccess("SuccessfullyCreatedPublicKey", clientId, "", publicKeyCache)
		} else {
			return DecryptKeyErr("FailedToFetchClientCache", err)
		}
	}

	secretKey, err := helpers.DecryptRSA(ctx, encryptedSecretKey, privateKeyCache)
	if err != nil {
		return DecryptKeyErr("FailedToDecryptKey", err)
	}

	return DecryptKeySuccess("Success", clientId, secretKey, "")
}

func (h Handler) DecryptPayload(ctx context.Context, secretKey,
	encryptedPayload string) response.BodyStruct {
	payload, err := helpers.DecryptAES(ctx, encryptedPayload, secretKey)
	if err != nil {
		return DecryptPayloadErr("FailedToDecryptPayload", err)
	}

	return DecryptPayloadSuccess(payload)
}

func (h Handler) EncryptPayload(ctx context.Context, secretKey,
	payload string) response.BodyStruct {
	encryptedPayload, err := helpers.EncryptAES(ctx, payload, secretKey)
	if err != nil {
		return EncryptPayloadErr("FailedToEncryptPayload", err)
	}

	return EncryptPayloadSuccess(encryptedPayload)
}

func (h Handler) Handshake(ctx context.Context, clientId string) response.BodyStruct {
	isValidClient := h.validClient(ctx, clientId)
	if !isValidClient {
		return HandshakeErr("InvalidClient", nil)
	}

	publicKeyCacheId := helpers.GetPublicKeyId(clientId)
	publicKeyCache, err := h.cacheHandler.Get(ctx, publicKeyCacheId).Result()
	if err != nil {
		if err == redis.Nil {
			out, err := h.kmsHandler.GenerateRsaKeyPair(ctx)
			if err != nil {
				return HandshakeErr("FailedToGenerateKeyPairs", err)
			}

			privateKeyCacheId := helpers.GetPrivateKeyId(clientId)
			privateKeyCache := base64.StdEncoding.EncodeToString(out.PrivateKeyPlaintext)
			if err := h.cacheHandler.Set(ctx, privateKeyCacheId, privateKeyCache, 0).Err(); err != nil {
				return HandshakeErr("FailedToUpdatePrivateKeyCache", err)
			}

			publicKeyCache = base64.StdEncoding.EncodeToString(out.PublicKey)
			if err := h.cacheHandler.Set(ctx, publicKeyCacheId, publicKeyCache, 0).Err(); err != nil {
				return HandshakeErr("FailedToUpdatePublicKeyCache", err)
			}
		} else {
			return HandshakeErr("FailedToFetchClientCache", err)
		}
	}

	return HandshakeSuccess(publicKeyCache)
}

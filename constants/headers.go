package constants

import "time"

const (
	REQUEST_ID_HEADER_KEY = "X-Request-ID"
	REQUEST_ID_PROP       = "requestId"

	REQUEST_AUTH_TOKEN_HEADER_KEY = "X-AuthToken"
	REQUEST_AUTH_TOKEN_PROP       = "authToken"

	REQUEST_CRYPTO_CLIENT_ID_HEADER_KEY = "X-API-Client-Id"
	REQUEST_CRYPTO_CLIENT_ID_PROP       = "apiClientId"

	REQUEST_CRYPTO_ENCRYPTION_KEY_HEADER_KEY = "X-API-Encryption-Key"
	REQUEST_CRYPTO_ENCRYPTION_KEY_PROP       = "apiEncryptionKey"

	REQUEST_APP_ID_HEADER_KEY = "X-AppId"
	REQUEST_APP_ID_PROP       = "appId"

	REQUEST_APP_SECRET_HEADER_KEY = "X-AppSecretKey"
	REQUEST_APP_SECRET_PROP       = "appSecretKey"

	REQUEST_MSG_SEQUENCE_HEADER_KEY = "X-MsgSequence"
	REQUEST_MSG_SEQUENCE_PROP       = "msgSequence"

	REQUEST_MSG_GROUP_HEADER_KEY = "X-MsgGroup"
	REQUEST_MSG_GROUP_PROP       = "msgGroup"

	REQUEST_SOURCE_HEADER_KEY = "X-Source"
	REQUEST_SOURCE_PROP       = "source"

	REQUEST_SOURCE_CHANNEL_HEADER_KEY = "X-SourceChannel"
	REQUEST_SOURCE_CHANNEL_PROP       = "sourceChannel"

	TOKEN_CACHE_KEY_PATTERN = "%s_auth"
)

var REQUEST_AUTH_TOKEN_EXPIRE_AT = func() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 00, n.Location())
}

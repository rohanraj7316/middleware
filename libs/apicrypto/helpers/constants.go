package helpers

var RSA_CONSTANTS = map[string]interface{}{}

var AES_256_GCM_CONSTANTS = struct {
	KEY_LENGTH     int
	NONCE_LENGTH   int
	DATA_SEPARATOR string
}{
	KEY_LENGTH:     32,
	NONCE_LENGTH:   16,
	DATA_SEPARATOR: ".",
}

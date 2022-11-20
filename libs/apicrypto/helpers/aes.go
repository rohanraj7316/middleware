package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func EncryptAES(ctx context.Context, payload, secretKey string) (string, error) {
	if len(secretKey) < AES_256_GCM_CONSTANTS.KEY_LENGTH {
		return "", fmt.Errorf("key len should be %d byte", AES_256_GCM_CONSTANTS.KEY_LENGTH)
	}

	payloadBytes := []byte(payload)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, AES_256_GCM_CONSTANTS.NONCE_LENGTH)
	if err != nil {
		return "", err
	}

	// Create a nonce. Nonce should be from GCM
	nonce := make([]byte, AES_256_GCM_CONSTANTS.NONCE_LENGTH)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedPayloadBytes := aesGCM.Seal(nil, nonce, payloadBytes, nil)

	nonceStr := base64.StdEncoding.EncodeToString(nonce)

	authTagStart := len(encryptedPayloadBytes) - 16
	authTagStr := base64.StdEncoding.EncodeToString(encryptedPayloadBytes[authTagStart:])

	encryptedPayload := base64.StdEncoding.EncodeToString(encryptedPayloadBytes[:authTagStart])

	fencryptedPayload := strings.Join([]string{nonceStr, authTagStr, encryptedPayload}, AES_256_GCM_CONSTANTS.DATA_SEPARATOR)

	return fencryptedPayload, nil
}

func DecryptAES(ctx context.Context, encryptedPayload, secretKey string) (string, error) {
	if len(secretKey) < AES_256_GCM_CONSTANTS.KEY_LENGTH {
		return "", fmt.Errorf("key len should be %d byte", AES_256_GCM_CONSTANTS.KEY_LENGTH)
	}
	fmt.Println(encryptedPayload)
	encryptedPayloadArr := strings.Split(encryptedPayload, AES_256_GCM_CONSTANTS.DATA_SEPARATOR)
	if len(encryptedPayloadArr) != 3 {
		return "", fmt.Errorf("after split array have %d want 3 length", len(encryptedPayloadArr))
	}

	nonceBytes, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[0])
	if err != nil {
		return "", err
	}

	authTagByte, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[1])
	if err != nil {
		return "", err
	}

	encryptedPayloadBytes, err := base64.StdEncoding.DecodeString(encryptedPayloadArr[2])
	if err != nil {
		return "", err
	}

	fEncryptedPayloadBytes := append(encryptedPayloadBytes, authTagByte...)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, AES_256_GCM_CONSTANTS.NONCE_LENGTH)
	if err != nil {
		return "", err
	}

	// Decrypt the data
	payloadBytes, err := aesGCM.Open(nil, nonceBytes, fEncryptedPayloadBytes, nil)
	if err != nil {
		return "", err
	}

	return string(payloadBytes), nil
}

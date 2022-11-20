package helpers

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func derToPemString(key, keyType string) string {
	return fmt.Sprintf("-----BEGIN %s-----\n%s\n-----END %s-----", keyType, key, keyType)
}

func LoadPublicKey(publicKey string) (*rsa.PublicKey, error) {
	var parsedKey interface{}
	dPem, _ := pem.Decode([]byte(publicKey))

	if dPem.Type == "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid key type: %s", dPem.Type)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(dPem.Bytes)
	if err != nil {
		return nil, err
	}

	if pubKey, ok := parsedKey.(*rsa.PublicKey); ok {
		return pubKey, nil
	}

	return nil, fmt.Errorf("failed to parse rsa public key")
}

func LoadPrivateKey(privateKey, privateKeyPassword string) (priKey *rsa.PrivateKey, err error) {
	privPem, _ := pem.Decode([]byte(privateKey))

	if privPem.Type == "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid key type: %s", privPem.Type)
	}

	var privPemBytes []byte
	if privateKeyPassword != "" {
		privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(privateKeyPassword))
		if err != nil {
			return nil, err
		}
	} else {
		privPemBytes = privPem.Bytes
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil {
			return nil, err
		}
	}

	if priKey, ok := parsedKey.(*rsa.PrivateKey); ok {
		return priKey, nil
	}

	return nil, fmt.Errorf("failed to parse rsa private key")
}

func EncryptRSA(ctx context.Context, payload string, publicKeyStr string) (string, error) {

	publicKey, err := LoadPublicKey(derToPemString(publicKeyStr, "PUBLIC KEY"))
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(payload), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptRSA(ctx context.Context, encryptedPayload string, privateKeyStr string) (string, error) {
	privateKey, err := LoadPrivateKey(derToPemString(privateKeyStr, "PRIVATE KEY"), "")
	if err != nil {
		return "", err
	}

	ct, err := base64.StdEncoding.DecodeString(encryptedPayload)
	if err != nil {
		return "", err
	}

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ct, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

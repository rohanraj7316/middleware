package awskms

import (
	"context"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type Handler struct {
	kmsKeyId *string

	Client *kms.KMS
}

func New(config ...Config) (Handler, error) {
	// fetching config details
	cfg := configDefault(config...)

	// seting up to run below code once
	sess, err := session.NewSession(cfg.AwsConfig)
	if err != nil {
		return Handler{}, err
	}
	svc := kms.New(sess)

	return Handler{
		kmsKeyId: cfg.KmsKeyId,
		Client:   svc,
	}, nil
}

func (h Handler) GenerateKMSKey(ctx context.Context, numberOfBytes int64) (encryptedKMSKey string, err error) {
	input := kms.GenerateDataKeyInput{
		KeyId:         h.kmsKeyId,
		NumberOfBytes: aws.Int64(numberOfBytes),
	}

	generateDataKeyOtp, err := h.Client.GenerateDataKeyWithContext(ctx, &input)
	if err != nil {
		return "", err
	}

	encryptedKMSKey = base64.RawStdEncoding.EncodeToString(generateDataKeyOtp.Plaintext)

	return encryptedKMSKey, nil
}

func (h Handler) GenerateRsaKeyPair(ctx context.Context) (*kms.GenerateDataKeyPairOutput, error) {
	in := &kms.GenerateDataKeyPairInput{
		KeyId:       h.kmsKeyId,
		KeyPairSpec: aws.String(kms.DataKeyPairSpecRsa2048),
	}

	out, err := h.Client.GenerateDataKeyPairWithContext(ctx, in)
	if err != nil {
		return out, err
	}

	return out, nil
}

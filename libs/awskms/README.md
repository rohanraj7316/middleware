# KMSClient


## Signature

```
func New(config ...*aws.Config) (*Config, error)
```

```
func GenerateEncryptedKMSKey(ctx context.Context, keyId string, numberOfBytes int64) (string, error)
func GenerateRsaKeyPair(ctx context.Context, keyId string) (*kms.GenerateDataKeyPairOutput, error)
```


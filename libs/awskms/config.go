package awskms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/rohanraj7316/middleware/libs/utils"
)

type Config struct {
	AwsConfig *aws.Config
	KmsKeyId  *string
}

var ConfigDefault = Config{
	AwsConfig: &aws.Config{},
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		rConfig := []string{
			ENV_AWS_DEFAULT_REGION,
		}
		uConfig := []string{
			ENV_AWS_KMS_KEY_ID,
		}
		eConfig := utils.EnvData(rConfig, uConfig)

		ConfigDefault.AwsConfig.Region = aws.String(eConfig[ENV_AWS_DEFAULT_REGION])
		ConfigDefault.KmsKeyId = aws.String(eConfig[ENV_AWS_KMS_KEY_ID])

		return ConfigDefault
	}

	cfg := config[0]

	return cfg
}

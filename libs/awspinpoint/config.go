package awspinpoint

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/rohanraj7316/logger"
)

type Config struct {
	AwsConfig   *aws.Config
	EmailClient *pinpointemail.PinpointEmail
}

var ConfigDefault = Config{
	AwsConfig: &aws.Config{},
}

func log(i ...interface{}) {
	dbLog := logger.Field{
		Key:   "aws-pin-point-logger",
		Value: i,
	}
	logger.Info("[AWS-PIN-POINT]", dbLog)
}

func configDefault(config ...Config) Config {
	ConfigDefault.AwsConfig.Region = aws.String(GetValue(AWS_DEFAULT_REGION, ""))
	ConfigDefault.AwsConfig.Logger = aws.LoggerFunc(log)
	// ConfigDefault.AwsConfig.Endpoint = aws.String(GetValue(AWS_PINPOINT_URL, ""))

	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	return cfg
}

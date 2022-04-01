package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/rohanraj7316/logger"
)

var AWSConfigDefault = &aws.Config{}

func log(i ...interface{}) {
	dbLog := logger.Field{
		Key:   "dynamodb-logger",
		Value: i,
	}
	logger.Info("[DYNAMO-DB ORM]", dbLog)
}

func configDefault(config ...*aws.Config) *aws.Config {

	AWSConfigDefault.Region = aws.String(GetValue(AWS_DEFAULT_REGION, ""))
	AWSConfigDefault.Logger = aws.LoggerFunc(log)

	if len(config) < 1 {
		return AWSConfigDefault
	}

	cfg := config[0]

	return cfg
}

package awspinpoint

import (
	"github.com/aws/aws-sdk-go/service/pinpointemail"
)

func New(config ...Config) Config {
	cfg := configDefault(config...)
	return cfg
}

func (c Config) Mail(input pinpointemail.SendEmailInput) (pinpointemail.SendEmailOutput, error) {
	return pinpointemail.SendEmailOutput{}, nil
}

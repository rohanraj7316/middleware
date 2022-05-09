package awspinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
)

func New(config ...Config) (Config, error) {
	cfg := configDefault(config...)
	se, err := session.NewSession(cfg.AwsConfig)
	if err != nil {
		return cfg, err
	}
	mSession := session.Must(se, nil)

	cfg.EmailClient = pinpointemail.New(mSession, cfg.AwsConfig)

	return cfg, nil
}

func (c Config) SendEmail(ctx context.Context, in *pinpointemail.SendEmailInput) (out *pinpointemail.SendEmailOutput,
	err error) {
	out, err = c.EmailClient.SendEmailWithContext(ctx, in)
	if err != nil {
		return out, err
	}
	return out, nil
}

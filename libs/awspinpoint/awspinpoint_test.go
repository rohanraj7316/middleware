package awspinpoint

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	env "github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

const (
	envFile    = ".env"
	EMAIL_TO   = "EMAIL_TO"
	EMAIL_FROM = "EMAIL_FROM"
)

var loadEnv = env.Load

func TestSimpleEmails(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	_, exists := os.LookupEnv(AWS_ACCESS_KEY_ID)
	if !exists {
		t.Errorf("missing %s", AWS_ACCESS_KEY_ID)
	}

	_, exists = os.LookupEnv(AWS_SECRET_ACCESS_KEY)
	if !exists {
		t.Errorf("missing %s", AWS_SECRET_ACCESS_KEY)
	}

	_, exists = os.LookupEnv(AWS_DEFAULT_REGION)
	if !exists {
		t.Errorf("missing %s", AWS_DEFAULT_REGION)
	}

	toEmail, exists := os.LookupEnv(EMAIL_TO)
	if !exists {
		t.Errorf("missing %s", EMAIL_TO)
	}

	fromEmail, exists := os.LookupEnv(EMAIL_FROM)
	if !exists {
		t.Errorf("missing %s", EMAIL_FROM)
	}

	// send mail
	eClient, err := New()
	if err != nil {
		t.Errorf("failed to create client %s", err.Error())
	}

	eBody := &pinpointemail.SendEmailInput{
		Content: &pinpointemail.EmailContent{
			Simple: &pinpointemail.Message{
				Body: &pinpointemail.Body{
					Text: &pinpointemail.Content{
						Data: aws.String("this is test email"),
					},
				},
				Subject: &pinpointemail.Content{
					Data: aws.String("TEST SUBJECT"),
				},
			},
		},
		Destination: &pinpointemail.Destination{
			ToAddresses: aws.StringSlice([]string{toEmail}),
		},
		FromEmailAddress: aws.String(fromEmail),
	}

	out, err := eClient.SendEmail(ctx, eBody)
	if err != nil {
		t.Errorf("error in sending email: %s", err.Error())
	}

	t.Log(out)
}

func TestAttachmentEmails(t *testing.T) {
	ctx := context.Background()

	// load env config
	err := loadEnv(envFile)
	if err != nil {
		t.Error(err)
	}

	_, exists := os.LookupEnv(AWS_ACCESS_KEY_ID)
	if !exists {
		t.Errorf("missing %s", AWS_ACCESS_KEY_ID)
	}

	_, exists = os.LookupEnv(AWS_SECRET_ACCESS_KEY)
	if !exists {
		t.Errorf("missing %s", AWS_SECRET_ACCESS_KEY)
	}

	_, exists = os.LookupEnv(AWS_DEFAULT_REGION)
	if !exists {
		t.Errorf("missing %s", AWS_DEFAULT_REGION)
	}

	toEmail, exists := os.LookupEnv(EMAIL_TO)
	if !exists {
		t.Errorf("missing %s", EMAIL_TO)
	}

	fromEmail, exists := os.LookupEnv(EMAIL_FROM)
	if !exists {
		t.Errorf("missing %s", EMAIL_FROM)
	}

	// send mail
	eClient, err := New()
	if err != nil {
		t.Errorf("failed to create client %s", err.Error())
	}

	msg := gomail.NewMessage()
	msg.SetHeader("Subject", "TESTING ATTACHMENT EMAIL")
	msg.SetBody("text/plain", "this is just a sample test email. pls, don't reply")
	msg.Attach("./attachments/sample_txt.txt")

	var emailRaw bytes.Buffer
	_, err = msg.WriteTo(&emailRaw)
	if err != nil {
		t.Errorf("failed to dump msg to bytes buffer")
	}

	eBody := &pinpointemail.SendEmailInput{
		Content: &pinpointemail.EmailContent{
			Raw: &pinpointemail.RawMessage{Data: emailRaw.Bytes()},
		},
		Destination: &pinpointemail.Destination{
			ToAddresses: aws.StringSlice([]string{toEmail}),
		},
		FromEmailAddress: aws.String(fromEmail),
	}

	out, err := eClient.SendEmail(ctx, eBody)
	if err != nil {
		t.Errorf("error in sending email: %s", err.Error())
	}

	t.Log(out)
}

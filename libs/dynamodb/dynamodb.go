package dynamodb

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rohanraj7316/logger"
)

var db *dynamo.DB

// NewDBSession create new dynamo session using aws sdk config
func NewDBSession(config ...*aws.Config) (*dynamo.DB, error) {
	var once sync.Once

	// seting up to run below code once
	once.Do(func() {
		cfg := configDefault(config...)

		sess, err := session.NewSession(cfg)
		if err != nil {
			logger.Error(err.Error())
		}

		db = dynamo.New(sess, cfg)
	})

	return db, nil
}

// Err format lib's error msg.
func Err(err error) error {
	return err
}

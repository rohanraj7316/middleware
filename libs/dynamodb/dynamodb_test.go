package dynamodb

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

const tableName = "testtable"

var (
	testDB *dynamo.DB
)

type TestTable struct {
}

func init() {
	if region := os.Getenv("DYNAMO_TEST_REGION"); region != "" {
		cfg := &aws.Config{
			Region: aws.String(region),
		}
		db, err := NewDBSession(cfg)
		if err != nil {

		}

		testDB = db
	}
}

func TestTableList(t *testing.T) {
	if testDB == nil {
		t.Skip("aws config not set")
	}

	table := testDB.Table(tableName)

	// insert into table

	// list table data
	var out interface{}
	err := Err(table.Scan().All(&out))
	if err != nil {
		t.Error(err)
	}

}

package response

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBodyStruct(t *testing.T) {
	tests := map[string]BodyStruct{
		"ErrTointerface": {
			StatusCode: http.StatusOK,
			Err:        fmt.Errorf("just a error msg"),
		},
		"byteToInterface": {
			StatusCode: http.StatusOK,
			Err:        []byte("this is me"),
		},
	}

	for _, test := range tests {
		v, err := test.ToJson()
		if err != nil {
			t.Errorf(err.Error())
		}

		t.Errorf("%s", v)
	}
}

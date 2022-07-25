package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rohanraj7316/httpclient"
)

type HttpClient struct {
	client *httpclient.HttpClient
	// passOnHeader    map[string]bool
	// passOnHeader passed on in the future requests.
	// Optional. Default: [empty map]
	passOnHeader map[string]bool
}

// New configure httpclient with the Config passed.
// Recommended - use it in singleton partern.
// avoid creating new client for each request call.
func New(config ...httpclient.Config) (*HttpClient, error) {
	client, err := httpclient.New(config...)
	if err != nil {
		return nil, err
	}

	return &HttpClient{
		client:       client,
		passOnHeader: map[string]bool{},
	}, nil
}

// responsible for updating 'hClient' flag.
// Default: "".
// send '#' seperated string.
// ie: "X-Auth-Key#X-Auth"
func (hClient *HttpClient) SetPassOnHeaders(headers string) {
	hArr := strings.Split(headers, "#")

	// trim the keys which are passed
	for i := 0; i < len(hArr); i++ {
		val := strings.TrimSpace(hArr[i])

		if _, ok := hClient.passOnHeader[val]; !ok {
			hClient.passOnHeader[val] = true
		}
	}
}

// Request should be used for sending http request with in
// resource model. helps to use fiber.Ctx insted of generating
// onw of it's own.
// Note: never create your own context.
func (hClient *HttpClient) Request(c Option) (*http.Response, error) {
	rCtx := c.Ctx.UserContext()
	return hClient.client.Request(rCtx, c.Method, c.Url, c.Header, c.RequestBody)
}

// RequestSDK should be used to send http request.
// trigger it from SDK. has an added layer for checking
// requestId.
// Note: never create your own context.
func (hClient *HttpClient) RequestSDK(c OptionSDK) (resBody interface{}, err error) {
	rID := c.Ctx.Value("requestId")
	if rID == "" {
		return nil, fmt.Errorf("'requestId' is missing in the context")
	}

	if val, ok := c.Header["Content-Type"]; !ok {
		c.Header["Content-Type"] = "application/json"
	} else if val != "application/json" {
		return nil, fmt.Errorf("invalid `Content-Type`")
	}

	// adding pass on headers.
	for key := range hClient.passOnHeader {
		c.Header[key] = c.Ctx.Value(key).(string)
	}

	reqBody, err := json.Marshal(c.RequestBody)
	if err != nil {
		return nil, err
	}

	r, err := hClient.client.Request(c.Ctx, c.Method,
		c.Url, c.Header, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(r.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

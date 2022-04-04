package httpclient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/httpclient"
)

type HttpClient struct {
	client *httpclient.HttpClient
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
		client: client,
	}, nil
}

// Request should be used for sending http request with in
// resource model. helps to use fiber.Ctx insted of generating
// onw of it's own.
// Note: never create your own context.
func (hClient *HttpClient) Request(c *fiber.Ctx, method, url string, headers map[string]string,
	request io.Reader) (*http.Response, error) {
	rCtx := c.UserContext()
	return hClient.client.Request(rCtx, method, url, headers, request)
}

// RequestSDK should be used to send http request.
// trigger it from SDK. has an added layer for checking
// requestId.
// Note: never create your own context.
func (hClient *HttpClient) RequestSDK(c context.Context, method, url string, headers map[string]string,
	request io.Reader) (*http.Response, error) {
	rID := c.Value("requestId")
	if rID == "" {
		return nil, fmt.Errorf("'requestId' is missing in the context")
	}

	return hClient.client.Request(c, method, url, headers, request)
}

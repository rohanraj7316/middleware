package httpclient

import (
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

// Request responsible for sending http request
// by using the Config set at the time of initialization.
func (hClient *HttpClient) Request(c *fiber.Ctx, method, url string, headers map[string]string,
	request io.Reader) (*http.Response, error) {
	rCtx := c.UserContext()
	return hClient.client.Request(rCtx, method, url, headers, request)
}

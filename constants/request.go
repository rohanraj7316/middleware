package constants

import "time"

type RequestIDType string

var REQUEST_TOKEN_TOKEN_TIMEOUT = func() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 00, n.Location())
}

const (
	REQUEST_ID_HEADER_KEY = "X-Request-ID"
	REQUEST_ID_PROP       = "requestId"

	REQUEST_ACCESS_TOKEN_KEY  = "X-Access-Token"
	REQUEST_ACCESS_TOKEN_PROP = "accessToken"

	REQUEST_TIMEOUT = "20s"
)

# Httpclient

Wrapper build on [httpclient](https://github.com/rohanraj7316/httpclient)

## Signature

```
func New(config ...httpclient.Config) (*HttpClient, error)
```

### Config

```
// https://github.com/rohanraj7316/httpclient config
type Config struct {
	// Timeout gives you timeout for request
	// Default: 30s
	Timeout time.Duration

	// bool flag which help us in configuring proxy
	// Default: false
	UseProxy bool

	// url need to do the proxy
	// Default: nil
	ProxyURL string

	// LogReqResEnable helps in logging request & responses.
	// Default true
	LogReqResEnable bool

	// LogReqResBodyEnable helps in logging request and responses body
	// Default true
	LogReqResBodyEnable bool
}
```

### Request

```
// Request should be used for sending http request with in
// resource model. helps to use fiber.Ctx insted of generating
// onw of it's own.
// Note: never create your own context.
```

### RequestSDK

```
// RequestSDK should be used to send http request.
// trigger it from SDK. has an added layer for checking
// requestId.
// Note: never create your own context.
```
# Middleware

it's a cluster of some pre-required middleware
## File Structure

### [config](config/)

below are the list of default configs:

- ServerDefault - it's **fiber.Config**
### [response](libs/response)

used to structure all the responses in **BodyStruct**

```
type BodyStruct struct {
	StatusCode int          `json:"statusCode,omitempty"`
	Status     string       `json:"status,omitempty"`
	Message    interface{}  `json:"message,omitempty"`
	Err        *ErrorStruct `json:"error,omitempty"`
	Data       interface{}  `json:"data,omitempty"`
}
```

### [httpclient](libs/httpclient)

wrapper build using [httpclient](https://github.com/rohanraj7316/httpclient). can be used to send http request.

### [validate](libs/validate)

helps in validating request body
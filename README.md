# Middleware

it's a cluster of some pre-required middleware
## File Structure

### [config](configs/)

below are the list of default configs:

- ServerDefault - it's **fiber.Config**

### [constant](constants/)

- TIME_FORMAT - **RFC3339** using this as default time format.
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

### [redisclient](libs/redisclient)

helps in creating connection with redis client

## [elasticclient] (libs/elasticclient)

helps in creating connection with elasticsearch
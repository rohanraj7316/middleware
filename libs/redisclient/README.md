# Redisclient

Redisclient built to establish connection with redis.

## Signature

```
func New(config Config) (*redis.Client, error)
```

## Config

type Config struct {
	HostName string
	Port     string
	Password string
	DB       int
}

Populate this Config struct from environment and send it to middleware redisclient New() function as a parameter.
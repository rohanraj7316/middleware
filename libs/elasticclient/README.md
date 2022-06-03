# Elasticclient

**elasticclient** built to establish connection with elasticsearch.

## Signature

```
func New(config ...Config) (Config, error)
```

## Config

```
type Config struct {
	ElasticSearchConfig *elasticsearch.Config
	Client              *elasticsearch.Client
}
```
## Env-variable

```
export ELASTIC_MAX_RETRIES
```
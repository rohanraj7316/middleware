package mongodb

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/rohanraj7316/middleware/libs/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ENV_MONGO_DB_CONNECTION_STRING = "MONGO_DB_CONNECTION_STRING"
	ENV_MONGO_DB_DATABASE_NAME     = "MONGO_DB_DATABASE_NAME"
	ENV_MONGO_DB_MAX_IDLE_TIME     = "MONGO_DB_MAX_IDLE_TIME"
	ENV_MONGO_DB_MAX_POOL_SIZE     = "MONGO_DB_MAX_POOL_SIZE"
	ENV_MONGO_DB_MIN_POOL_SIZE     = "MONGO_DB_MIN_POOL_SIZE"
	ENV_MONGO_DB_IS_TLS            = "MONGO_DB_IS_TLS"
	ENV_MONGO_DB_TLS_PEM_FILE_PATH = "MONGO_DB_TLS_PEM_FILE_PATH"
)

type Config struct {
	cOptions *options.ClientOptions
	dOptions *options.DatabaseOptions

	DbName string
	Client *mongo.Client
}

var ConfigDefault = Config{
	cOptions: &options.ClientOptions{},
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)
	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, fmt.Errorf("failed parsing pem file")
	}

	return tlsConfig, nil
}

func configDefault(config ...Config) (Config, error) {
	if len(config) < 1 {
		rEnvCfg := []string{
			ENV_MONGO_DB_CONNECTION_STRING,
			ENV_MONGO_DB_DATABASE_NAME,
		}
		uEnvCfg := []string{
			ENV_MONGO_DB_MAX_IDLE_TIME,
			ENV_MONGO_DB_MIN_POOL_SIZE,
			ENV_MONGO_DB_MAX_POOL_SIZE,
			ENV_MONGO_DB_IS_TLS,
			ENV_MONGO_DB_TLS_PEM_FILE_PATH,
		}
		envCfg := utils.EnvData(rEnvCfg, uEnvCfg)

		ConfigDefault.cOptions.ApplyURI(envCfg[ENV_MONGO_DB_CONNECTION_STRING])

		if _, ok := envCfg[ENV_MONGO_DB_MAX_IDLE_TIME]; ok {
			maxConnIdleTime, err := time.ParseDuration(fmt.Sprintf("%ss", envCfg[ENV_MONGO_DB_MAX_IDLE_TIME]))
			if err != nil {
				return ConfigDefault, err
			}
			ConfigDefault.cOptions = ConfigDefault.cOptions.SetMaxConnIdleTime(maxConnIdleTime)
		}

		if _, ok := envCfg[ENV_MONGO_DB_MAX_POOL_SIZE]; ok {
			maxPoolSize, err := strconv.Atoi(envCfg[ENV_MONGO_DB_MAX_POOL_SIZE])
			if err != nil {
				return ConfigDefault, err
			}
			ConfigDefault.cOptions = ConfigDefault.cOptions.SetMaxPoolSize(uint64(maxPoolSize))
		}

		if _, ok := envCfg[ENV_MONGO_DB_MIN_POOL_SIZE]; ok {
			minPoolSize, err := strconv.Atoi(envCfg[ENV_MONGO_DB_MIN_POOL_SIZE])
			if err != nil {
				return ConfigDefault, err
			}
			ConfigDefault.cOptions = ConfigDefault.cOptions.SetMinPoolSize(uint64(minPoolSize))
		}

		if isTls, ok := envCfg[ENV_MONGO_DB_IS_TLS]; ok {
			if isTls == "true" {
				if filepath, ok := envCfg[ENV_MONGO_DB_TLS_PEM_FILE_PATH]; ok {
					tlsCfg, err := getCustomTLSConfig(filepath)
					if err != nil {
						return ConfigDefault, err
					}

					ConfigDefault.cOptions = ConfigDefault.cOptions.SetTLSConfig(tlsCfg)
				} else {
					return ConfigDefault, fmt.Errorf("empty tls pathname")
				}
			}
		}

		if _, ok := envCfg[ENV_MONGO_DB_DATABASE_NAME]; ok {
			ConfigDefault.DbName = envCfg[ENV_MONGO_DB_DATABASE_NAME]
		}
	}

	return ConfigDefault, nil
}

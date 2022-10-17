package mongodb

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rohanraj7316/middleware/libs/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ENV_MONGO_DB_CONNECTION_STRING = "MONGO_DB_CONNECTION_STRING"
	ENV_MONGO_DB_HOSTS             = "MONGO_DB_HOSTS"
	ENV_MONGO_DB_MAX_IDLE_TIME     = "MONGO_DB_MAX_IDLE_TIME"
	ENV_MONGO_DB_MAX_POOL_SIZE     = "MONGO_DB_MAX_POOL_SIZE"
	ENV_MONGO_DB_MIN_POOL_SIZE     = "MONGO_DB_MIN_POOL_SIZE"
	ENV_MONGO_DB_USERNAME          = "MONGO_DB_USERNAME"
	ENV_MONGO_DB_PASSWORD          = "MONGO_DB_PASSWORD"
)

type Config struct {
	cOptions  *options.ClientOptions
	dbOptions *options.DatabaseOptions

	DbName string
	Client *mongo.Client
	Db     *mongo.Database
}

var ConfigDefault = Config{}

func configDefault(config ...Config) (Config, error) {
	ConfigDefault.cOptions = &options.ClientOptions{}

	rEnvCfg := []string{
		ENV_MONGO_DB_HOSTS,
		// ENV_MONGO_DB_USERNAME,
		// ENV_MONGO_DB_PASSWORD,
		ENV_MONGO_DB_MAX_IDLE_TIME,
		ENV_MONGO_DB_MIN_POOL_SIZE,
		ENV_MONGO_DB_MAX_POOL_SIZE,
	}
	rCfg := utils.RequiredFields(rEnvCfg)

	hostsStr := rCfg[ENV_MONGO_DB_HOSTS]
	hostArr := strings.Split(hostsStr, ";")
	ConfigDefault.cOptions = ConfigDefault.cOptions.SetHosts(hostArr)

	ConfigDefault.cOptions.ApplyURI()

	iDur, err := time.ParseDuration(fmt.Sprintf("%ss", rCfg[ENV_MONGO_DB_MAX_IDLE_TIME]))
	if err != nil {
		return ConfigDefault, err
	}
	ConfigDefault.cOptions = ConfigDefault.cOptions.SetMaxConnIdleTime(iDur)

	pMaxSize, err := strconv.Atoi(rCfg[ENV_MONGO_DB_MAX_POOL_SIZE])
	if err != nil {
		return ConfigDefault, err
	}
	ConfigDefault.cOptions = ConfigDefault.cOptions.SetMaxPoolSize(uint64(pMaxSize))

	pMinSize, err := strconv.Atoi(rCfg[ENV_MONGO_DB_MIN_POOL_SIZE])
	if err != nil {
		return ConfigDefault, err
	}
	ConfigDefault.cOptions = ConfigDefault.cOptions.SetMinPoolSize(uint64(pMinSize))

	ConfigDefault.cOptions = ConfigDefault.cOptions.SetAuth(options.Credential{
		Username: rCfg[ENV_MONGO_DB_USERNAME],
		Password: rCfg[ENV_MONGO_DB_PASSWORD],
	})

	return ConfigDefault, nil
}

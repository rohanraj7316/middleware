package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rohanraj7316/logger"
	"github.com/rohanraj7316/middleware/constants"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// loggerReqResLogEnable flag set for logging all request and response
	// Optional. Default: true
	loggerReqResLogEnable bool
	// TODO: need to check this why are we using this
	// loggerCodeFlowLogEnable    bool
	// if loggerReqResLogBodyEnabled flag disabled. it wouldn't log
	// reqHeaders, respHeader, reqBody & resBody.
	// Optional. Default: true
	loggerReqResLogBodyEnabled bool
	requestIdConfig            requestid.Config
	reqResLogger               flogger.Config
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool
	// requestTimeout is the maximum amount of time to wait for the
	// request to send back the response.
	// Default: 20sec
	requestTimeout string
}

var ConfigDefault = Config{
	requestIdConfig: requestid.Config{
		Header:     constants.REQUEST_ID_HEADER_KEY,
		ContextKey: constants.REQUEST_ID_PROP,
		Next: func(c *fiber.Ctx) bool {
			rId := c.Get(constants.REQUEST_ID_HEADER_KEY)
			if rId != "" {
				// Set new id to response header
				c.Set(constants.REQUEST_ID_HEADER_KEY, rId)

				// Add the request ID to locals
				c.Locals(constants.REQUEST_ID_PROP, rId)
				return true
			}
			return false
		},
	},
	loggerReqResLogEnable:      true,
	loggerReqResLogBodyEnabled: true,
	reqResLogger: flogger.Config{
		Format: constants.REQ_RES_RECV_MSG_FORMAT,
	},
	requestTimeout: "20s",
	Next:           nil,
}

func configDefault(config ...Config) Config {
	ConfigDefault.requestTimeout = "20s"
	ConfigDefault.reqResLogger.Output = ConfigDefault
	ConfigDefault.reqResLogger.Next = func(c *fiber.Ctx) bool {
		return !ConfigDefault.loggerReqResLogEnable
	}

	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	return cfg
}

// responsible for updating 'loggerReqResLogEnable' flag.
// if set true all the http request hitting the servers
// gonna be get logged.
// Default: true.
func (c *Config) SetReqResLog(reqResLog bool) {
	c.loggerReqResLogEnable = reqResLog
}

// responsible for updating 'loggerReqResLogBodyEnabled' flag.
// if set true all the http request body, response body
// gonna be get logged.
// Default: true.
func (c *Config) SetReqResBodyLog(reqResBodyLog bool) {
	c.loggerReqResLogBodyEnabled = reqResBodyLog
}

// responsible for updating 'requestTimeout' flag.
// Default: 20s.
func (c *Config) SetRequestTimeout(timeoutStr string) {
	c.requestTimeout = timeoutStr
}

func (c Config) Write(p []byte) (int, error) {
	go func() {
		lBody := []zapcore.Field{}
		lMessage := []interface{}{}
		tagString := string(p)
		tagArr := strings.Split(tagString, ",")
		for i := 0; i < len(tagArr); i++ {
			tag := strings.Split(tagArr[i], "=")
			key := tag[0]
			value := strings.Join(tag[1:], "=")
			// avoid request and response body logging if 'loggerReqResLogBodyEnabled' is disabled
			if c.loggerReqResLogBodyEnabled ||
				(key != "reqHeaders" && key != "respHeader" && key != "reqBody" && key != "resBody") {
				lBody = append(lBody, zapcore.Field{
					Key:    key,
					String: value,
					Type:   zapcore.StringType,
				})
			}

			if key == "status" || key == "method" ||
				key == "path" || key == "latency" {
				lMessage = append(lMessage, tag[1])
			}
		}
		// TODO: do we need to log status != 200 as error?
		logger.Info(fmt.Sprintf(constants.REQ_RES_LOG_MSG, lMessage...), lBody...) // [REQ-RES-LOG] 200 POST /health 2.3sec
	}()
	return 0, nil
}

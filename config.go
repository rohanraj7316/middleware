package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rohanraj7316/middleware/constants"
)

type Config struct {
	// loggerReqResLogEnable flag set for logging all request and response
	// Optional. Default: true
	loggerReqResLogEnable bool
	// TODO: need to check this why are we using this
	// loggerCodeFlowLogEnable    bool
	// loggerReqResLogBodyEnabled flag set for logging all request and response's body
	// Optional. Default: true
	loggerReqResLogBodyEnabled bool
	requestIdConfig            requestid.Config
	reqResLogger               logger.Config
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
	reqResLogger: logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	},
	requestTimeout: "20s",
	Next:           nil,
}

func configDefault(config ...Config) Config {
	// ConfigDefault.reqResLogger.Output = logger.Writer
	ConfigDefault.reqResLogger.Next = func(c *fiber.Ctx) bool {
		return !ConfigDefault.loggerReqResLogEnable
	}

	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	ConfigDefault.requestTimeout = "20s"

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

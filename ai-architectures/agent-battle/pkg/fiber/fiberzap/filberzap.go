package fiberzap

import (
	"os"
	"sync"
	"time"

	"agent-battle/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Config defines the config for middleware
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool
}

// New creates a new middleware handler
func New(config Config) fiber.Handler {
	var (
		errPadding = 15
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) error {
		if config.Next != nil && config.Next(c) {
			return c.Next()
		}

		once.Do(func() {
			errHandler = c.App().Config().ErrorHandler
			stack := c.App().Stack()
			for m := range stack {
				for r := range stack[m] {
					if len(stack[m][r].Path) > errPadding {
						errPadding = len(stack[m][r].Path)
					}
				}
			}
		})

		start := time.Now()

		chainErr := c.Next()

		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		req := Req(c)
		// resp := Resp(c.Response())

		fields := []zap.Field{
			zap.Int("pid", os.Getpid()),
			zap.Duration("latency", time.Since(start)),
			zap.Object("http", req),
			// zap.Object("http.response", resp),
			zap.Int("http.status_code", c.Response().StatusCode()),
			zap.String("network.client.ip", c.IP()),
			zap.String("user", req.user),
		}

		if chainErr != nil {
			fields = append(fields, zap.Error(chainErr))
			logger.GetLoggerInstanceFromContext(c.UserContext()).With(fields...).Error(req.method + " " + req.path)
			return nil
		}

		logger.GetLoggerInstanceFromContext(c.UserContext()).With(fields...).Info(req.method + " " + req.path)

		return nil
	}
}

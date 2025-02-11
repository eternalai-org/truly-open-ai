package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"agent-battle/pkg/fiber/fiberzap"
	"agent-battle/pkg/logger"
	"agent-battle/pkg/utils"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	SwaggerPath     = "/swagger"
	HealthCheckPath = "/api/health"
	HalfOfMegaBytes = 1 << 19

	LocalAuthUserAddress = "user_address"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	logger.GetLoggerInstanceFromContext(c.UserContext()).Error("DefaultErrorHandler",
		zap.Error(err))
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	errDetail := utils.FromErr(err)
	return c.Status(errDetail.HttpCode).JSON(errDetail)
}

var jsonEncoder = JsonEncoder()

func JsonEncoder() fiberUtils.JSONMarshal {
	ops := protojson.MarshalOptions{
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}
	return func(v interface{}) ([]byte, error) {
		if protoMessage, ok := v.(proto.Message); ok {
			return ops.Marshal(protoMessage)
		}
		return json.Marshal(v)
	}
}

func Setup() *fiber.App {
	defaultConfig()
	disableSwagger := viper.GetBool("DISABLE_SWAGGER_HANDLER")

	app := fiber.New(fiber.Config{
		ErrorHandler:   DefaultErrorHandler,
		JSONEncoder:    jsonEncoder,
		BodyLimit:      viper.GetInt("BODY_LIMIT_IN_BYTES"),
		ReadBufferSize: HalfOfMegaBytes,
	})

	app.Use(corsHandler())
	app.Use(recoverHandler())
	app.Use(compressHandler())
	app.Use(logMiddleware())

	// Health check
	app.Get(HealthCheckPath, func(c *fiber.Ctx) error {
		return c.SendString("UP")
	})

	if !disableSwagger {
		app.Get(fmt.Sprintf("%s/*", SwaggerPath), swagger.New())
	}

	return app
}

func defaultConfig() {
	viper.SetDefault("BODY_LIMIT_IN_BYTES", 30*1024*1024) // 30MB
}

func recoverHandler() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			req := fiberzap.Req(c)
			if err, ok := e.(error); ok {
				logger.GetLoggerInstanceFromContext(c.UserContext()).Error("panic",
					zap.Object("request", req),
					zap.Error(err))
			} else {
				logger.GetLoggerInstanceFromContext(c.UserContext()).Error("panic",
					zap.Object("request", req),
					zap.Any("error", e))
			}
		},
	})
}

func logMiddleware() fiber.Handler {
	return fiberzap.New(fiberzap.Config{
		Next: func(c *fiber.Ctx) bool {
			return publicPathFilter(c)
		},
	})
}

var publicPathFilter = func(ctx *fiber.Ctx) bool {
	method := strings.ToUpper(ctx.Method())
	if http.MethodOptions == method {
		return true
	}

	path := ctx.Path()
	if HealthCheckPath == path {
		return true
	}
	if strings.Contains(path, SwaggerPath) {
		return true
	}
	return false
}

func corsHandler() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     viper.GetString("CORS_ALLOW_ORIGINS"),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		ExposeHeaders:    "content-disposition,content-length",
	})
}

func compressHandler() fiber.Handler {
	return compress.New()
}

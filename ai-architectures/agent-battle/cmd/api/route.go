package api

import (
	"agent-battle/cmd/setting"
	"agent-battle/docs"
	"agent-battle/internal/adapters/handler/game"
	"agent-battle/internal/core/middleware"
	"agent-battle/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type APIServer struct {
	App *fiber.App
}

func Init(s *setting.Setting) *APIServer {
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	app := middleware.Setup()

	// Middleware for /api/v1
	apiV1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("version", "v1")
		return c.Next()
	})

	if utils.IsWorker() {
		return &APIServer{App: app}
	}

	game.NewGameHandler(apiV1, s.GameUsecase)

	return &APIServer{App: app}
}

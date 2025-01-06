package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) startRouter() *gin.Engine {
	app := gin.New()
	app.Use(
		gin.LoggerWithWriter(gin.DefaultWriter,
			"/api/health"),
		gin.Recovery(),
	)

	app.GET("/api/health", s.health)

	return app
}

func (s *Server) health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

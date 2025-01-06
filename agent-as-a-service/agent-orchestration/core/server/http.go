package server

import (
	"context"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

// HTTPServer ...
type HTTPServer struct {
	engine *gin.Engine
	logger *zap.Logger
}

// NewHTTPServer ...
func NewHTTPServer(logger *zap.Logger) *HTTPServer {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// Middleware
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	server := &HTTPServer{
		engine: r,
		logger: logger,
	}

	return server
}

// Engine ...
func (s *HTTPServer) Engine() *gin.Engine {
	return s.engine
}

// Logger ...
func (s *HTTPServer) Logger() *zap.Logger {
	return s.logger
}

// GetPagingFromContext ...
func (s *HTTPServer) GetPagingFromContext(c *gin.Context) (uint, uint) {
	var (
		pageS  = c.DefaultQuery("page", "1")
		limitS = c.DefaultQuery("limit", "10")
		page   int
		limit  int
		err    error
	)

	page, err = strconv.Atoi(pageS)
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(limitS)
	if err != nil {
		limit = 10
	}

	return uint(page), uint(limit)
}

// GetContext ...
func (s *HTTPServer) GetContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	return ctx
}

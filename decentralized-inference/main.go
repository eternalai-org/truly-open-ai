package main

import (
	"decentralized-inference/cmd/server"
	"decentralized-inference/internal/logger"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	app := &cli.App{
		Name:  "decentralized-inference",
		Usage: "EAI decentralized-inference application",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "start the HTTP server to serve the decentralized-inference API",
				Action: func(c *cli.Context) error {
					logger.GetLoggerInstanceFromContext(c.Context).Info("Starting server")
					svr, err := server.NewServer()
					if err != nil {
						logger.GetLoggerInstanceFromContext(c.Context).Error("Failed to start server", zap.Error(err))
						return err
					}
					svr.Start()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

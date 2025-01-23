package main

import (
	"decentralized-inference/cmd/other/chat"
	"decentralized-inference/cmd/server"
	"decentralized-inference/internal/logger"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	app := &cli.App{
		Name:  "eai",
		Usage: "EAI application",
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
			{
				Name:        "chat",
				Usage:       "EAI chat with agent, use 'chat <agent_id>' to start chat",
				Description: "EAI chat",
				Category:    "chat",
				Args:        true,
				Action: func(c *cli.Context) error {
					agentID := c.Args().Get(0)
					if agentID == "" {
						return fmt.Errorf("Agent ID is required")
					}
					_, err := strconv.Atoi(agentID)
					if err != nil {
						return fmt.Errorf("Agent ID is invalid, it must be a number :%v", agentID)
					}
					//logger.GetLoggerInstanceFromContext(c.Context).Info("Starting chat with agent terminal")
					return chat.AgentTerminalChat(c.Context, agentID)
				},
				Subcommands: []*cli.Command{
					{
						Name: "config-all",
						Action: func(c *cli.Context) error {
							return chat.AgentTerminalChatConfig(c.Context)
						},
					},
				},
			},
			{
				Name:  "sys-prompt",
				Usage: "EAI get agent system prompt, use 'chat <agent_id>' to get it",
				Action: func(c *cli.Context) error {
					agentID := c.Args().Get(0)
					if agentID == "" {
						return fmt.Errorf("agent ID is required")
					}
					agentIDInt, err := strconv.Atoi(agentID)
					if err != nil {
						return fmt.Errorf("agent ID is invalid, it must be a number :%v", agentID)
					}
					return chat.GetSystemPromptFromContract(c.Context, agentIDInt)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

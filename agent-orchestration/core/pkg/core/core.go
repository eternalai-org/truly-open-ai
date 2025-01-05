package core

import (
	"github.com/eternalai-org/eternal-ai/agent-orchestration/core/pkg/core/wallet"
)

// Core ...
type Core struct {
	Endpoint string
	Wallet   *wallet.Wallet
}

// Init ...
func Init(e string) *Core {
	return &Core{
		Endpoint: e,
		Wallet:   wallet.Init(e),
	}
}

package factory

import (
	"solo/internal/port"
	"solo/internal/usecase"
)

func NewLocalChain(localChainVersion string) port.ICMDLocalChain {
	switch localChainVersion {
	case "v1":
		return usecase.NewCMDLocalChainV1()
	}
	return usecase.NewCMDLocalChainV2()
}

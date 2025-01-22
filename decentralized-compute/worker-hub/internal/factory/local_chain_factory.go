package factory

import (
	"solo/internal/port"
	"solo/internal/usecase"
)

func NewLocalChain(localChainVersion string) port.ICMDLocalChain {
	var localChainCMD port.ICMDLocalChain
	switch localChainVersion {
	case "v1":
		localChainCMD, _ = usecase.NewCMDLocalChainV1()
	default:
		localChainCMD, _ = usecase.NewCMDLocalChainV2()
	}

	return localChainCMD

}

package ethapi

import (
	"errors"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/gpumanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (c *Client) GPUManagerGetModelID(contactAddress string) (uint32, error) {
	client, err := c.getClient()
	if err != nil {
		return 0, err
	}
	instance, err := gpumanager.NewGPUManager(helpers.HexToAddress(contactAddress), client)
	if err != nil {
		return 0, err
	}
	modelIds, err := instance.GetModelIds(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	if len(modelIds) == 0 {
		return 0, errors.New("modelId not found")
	}
	return uint32(modelIds[0].Int64()), nil
}

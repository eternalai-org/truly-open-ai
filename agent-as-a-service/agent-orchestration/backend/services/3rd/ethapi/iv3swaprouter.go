package ethapi

import (
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/iv3swaprouter"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func (c *Client) IV3SwapRouterExactInputSingleData(params *iv3swaprouter.IV3SwapRouterExactInputSingleParams) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(iv3swaprouter.IV3SwapRouterABI))
	if err != nil {
		return nil, err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"exactInputSingle",
			params,
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"refundETH",
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		multicallBytes,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

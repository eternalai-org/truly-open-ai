package ethapi

import (
	"errors"
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/iworkerhub"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) WorkerHubAssignment(workerhubAddress string, assignmentId uint) (*big.Int, []byte, error) {
	if !common.IsHexAddress(workerhubAddress) {
		return nil, nil, errors.New("workerhubAddress is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return nil, nil, err
	}
	instance, err := iworkerhub.NewIWorkerHub(helpers.HexToAddress(workerhubAddress), client)
	if err != nil {
		return nil, nil, err
	}
	assignments, err := instance.GetAllAssignments(&bind.CallOpts{}, big.NewInt(int64(assignmentId)), big.NewInt(1))
	if err != nil {
		return nil, nil, err
	}
	return assignments[0].InferenceId, assignments[0].Output, nil
}

func (c *Client) WorkerHubInference(workerhubAddress string, inferenceId uint) (*iworkerhub.IWorkerHubInference, error) {
	if !common.IsHexAddress(workerhubAddress) {
		return nil, errors.New("workerhubAddress is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := iworkerhub.NewIWorkerHub(helpers.HexToAddress(workerhubAddress), client)
	if err != nil {
		return nil, err
	}
	inference, err := instance.GetInferenceInfo(&bind.CallOpts{}, big.NewInt(int64(inferenceId)))
	if err != nil {
		return nil, err
	}
	return &inference, nil
}

func (c *Client) GetAssignmentByInferenceId(workerhubAddress string, inferenceId uint) ([]iworkerhub.IWorkerHubAssignment, error) {
	if !common.IsHexAddress(workerhubAddress) {
		return nil, errors.New("workerhubAddress is invalid")
	}
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	instance, err := iworkerhub.NewIWorkerHub(helpers.HexToAddress(workerhubAddress), client)
	if err != nil {
		return nil, err
	}
	assignments, err := instance.GetAssignmentByInferenceId(&bind.CallOpts{}, big.NewInt(int64(inferenceId)))
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

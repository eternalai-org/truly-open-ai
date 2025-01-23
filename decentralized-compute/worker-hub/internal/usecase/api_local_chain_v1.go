package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"solo/internal/contracts/v1/hybrid_model"
	"solo/internal/contracts/v1/worker_hub"

	"solo/internal/model"
	"solo/pkg"
	"solo/pkg/eth"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type API_Local_Chain_V1 struct {
	gasPrice *big.Int
	gasLimit uint64

	rpc         string
	chainID     string
	prvKey      string
	cnf         *model.LocalChain
	workerHub   *worker_hub.WorkerHub
	hybridModel *hybrid_model.HybridModel
	client      *ethclient.Client
}

func NewAPILocalChainV1() (*API_Local_Chain_V1, error) {

	gasPrice := big.NewInt(pkg.LOCAL_CHAIN_GAS_PRICE)
	c := &API_Local_Chain_V1{
		gasPrice: gasPrice,
		gasLimit: pkg.LOCAL_CHAIN_GAS_LIMIT,
	}

	localCnf := c.ReadLocalChainCnf()
	if localCnf != nil {
		c.rpc = localCnf.Rpc
		c.chainID = localCnf.ChainID
		c.prvKey = localCnf.PrivateKey
	}

	cnf := c.ReadLocalChainCnf()
	cnf.Rpc = "http://hardhat:8545"
	c.cnf = cnf
	client, err := eth.NewEthClient(cnf.Rpc)
	if err != nil {
		return nil, err
	}

	hbContract, err := hybrid_model.NewHybridModel(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1]), client)
	if err != nil {
		return nil, err
	}

	wkHub, err := worker_hub.NewWorkerHub(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER_V1]), client)
	if err != nil {
		return nil, err
	}

	c.workerHub = wkHub
	c.hybridModel = hbContract
	c.client = client
	return c, nil
}

func (c *API_Local_Chain_V1) ReadLocalChainCnf() *model.LocalChain {
	resp := new(model.LocalChain)
	resp.Contracts = make(map[string]string)
	resp.Miners = make(map[string]model.Miners)
	path := fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir())
	_b, err := os.ReadFile(path)
	if err != nil {
		err1 := os.Mkdir(fmt.Sprintf(pkg.ENV_FOLDER, pkg.CurrentDir()), os.ModePerm)
		if err1 == nil {
			err2 := pkg.CreateFile(path, []byte{})
			if err2 != nil {
				return nil
			}
		}
		return resp
	}

	err = json.Unmarshal(_b, resp)
	if err != nil {
		return resp
	}

	return resp
}

func (c *API_Local_Chain_V1) CreateInfer(ctx context.Context, request model.LLMInferRequest) (*types.Transaction, *uint64, *model.LLMInferResponse, error) {
	cnf := c.cnf
	privKey := cnf.PrivateKey
	client := c.client

	auth, err := eth.CreateBindTransactionOpts(ctx, c.client, privKey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		return nil, nil, nil, err
	}

	p := c.hybridModel
	_, pubkey, err := eth.GetAccountInfo(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	_b, err := json.Marshal(request)
	tx, err := p.Infer(auth, _b, true)
	if err != nil {
		fmt.Println("wkHubAddress:", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1])
		fmt.Println("err:", err)
		return nil, nil, nil, err
	}

	txReceipt, err := eth.WaitForTxReceipt(client, tx.Hash())
	if err != nil {
		return nil, nil, nil, errors.Join(err, errors.New("Error while waiting for tx"))
	}

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, nil, nil, err
	}

	_ = txReceipt

	wkHub := c.workerHub
	logs := receipt.Logs
	inferIdBig := big.NewInt(0)
	for _, item := range logs {
		inferData, err := wkHub.ParseNewInference(*item)
		if err == nil {
			inferIdBig = inferData.InferenceId
			break
		}
	}

	var inferId = inferIdBig.Uint64()
	//wait for result
	chatCompletion := &model.LLMInferResponse{}
	index := 0

break_here:
	for index < 150 {
		time.Sleep(2 * time.Second)
		assignmentIDs, err := wkHub.GetAssignmentsByInference(nil, inferIdBig)
		if err != nil {
			return tx, &inferId, nil, err
		}

		for _, assismentID := range assignmentIDs {

			aInfo, err := wkHub.GetAssignmentInfo(nil, assismentID)
			if err != nil {
				return tx, &inferId, nil, err
			}

			outByte := aInfo.Output
			out := string(outByte)
			if out != "" {

				response := &model.Response{}
				err := json.Unmarshal([]byte(out), response)
				if err != nil {
					return tx, &inferId, nil, err
				}

				encodedString := response.Data
				decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
				if err != nil {
					return tx, &inferId, nil, err
				}

				if err = json.Unmarshal(decodedBytes, chatCompletion); err != nil {
					return tx, &inferId, nil, err
				}
				break break_here
			}
		}

		index += 1
	}

	if chatCompletion == nil {
		return tx, &inferId, nil, errors.New("error while parse response")
	}

	if len(chatCompletion.Choices) == 0 {
		return tx, &inferId, nil, errors.New("error get data")
	}

	_ = txReceipt
	_ = receipt
	_ = pubkey

	return tx, &inferId, chatCompletion, nil
}

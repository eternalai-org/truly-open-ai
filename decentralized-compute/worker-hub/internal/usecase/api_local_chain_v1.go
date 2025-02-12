package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"math/big"
	"os"
	"time"

	"solo/internal/contracts/v1/hybrid_model"
	"solo/internal/contracts/v1/worker_hub"

	"github.com/ethereum/go-ethereum/ethclient"

	"solo/internal/model"
	"solo/pkg"
	"solo/pkg/eth"

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
	rdb         *redis.Client
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

	cnf.Rpc = fmt.Sprintf("http://%s:8545", pkg.MINER_SERVICE_HARDHAT)
	cnf.PubSubURL = fmt.Sprintf("%s:6379", pkg.REDIS_PUBSUB)

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

	// Connect to the Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.PubSubURL,
		Password: "", // no password set
		DB:       0,  // default DB
	})

	c.workerHub = wkHub
	c.hybridModel = hbContract
	c.client = client
	c.rdb = rdb
	return c, nil
}

func (c *API_Local_Chain_V1) ReadLocalChainCnf() *model.LocalChain {
	resp := new(model.LocalChain)
	resp.Contracts = make(map[string]string)
	resp.Miners = make(map[string]model.Miners)
	path := fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir())
	_b, err := os.ReadFile(path)
	if err != nil {
		if err := os.Mkdir(fmt.Sprintf(pkg.ENV_FOLDER, pkg.CurrentDir()), os.ModePerm); err == nil {
			if err := pkg.CreateFile(path, []byte{}); err != nil {
				return nil
			}
		}
		return resp
	}

	_ = json.Unmarshal(_b, resp)
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
	if err != nil {
		return nil, nil, nil, err
	}

	tx, err := p.Infer(auth, _b, true)
	if err != nil {
		fmt.Println("wkHubAddress:", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1])
		fmt.Println("err:", err)
		return nil, nil, nil, err
	}

	txReceipt, err := eth.WaitForTxReceipt(client, tx.Hash())
	if err != nil {
		return nil, nil, nil, errors.Join(err, errors.New("error while waiting for tx"))
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

	inferId := inferIdBig.Uint64()
	// wait for result
	chatCompletion := &model.LLMInferResponse{}
	index := 0

	pBFTCommittee := []string{}
	_pBFTCommittee := make(map[string]string)
	proposer := ""
break_here:
	for index < 15000 {
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
			_pBFTCommittee[aInfo.Worker.Hex()] = aInfo.Worker.Hex()

			outByte := aInfo.Output
			out := string(outByte)
			if out != "" {
				if aInfo.Role == 2 {
					proposer = aInfo.Worker.Hex()
				}

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

	// if chatCompletion == nil {
	// 	return tx, &inferId, nil, errors.New("error while parse response")
	// }

	if len(chatCompletion.Choices) == 0 {
		return tx, &inferId, nil, errors.New("error get data")
	}

	_ = txReceipt
	_ = receipt
	_ = pubkey

	for _, v := range _pBFTCommittee {
		pBFTCommittee = append(pBFTCommittee, v)
	}

	chatCompletion.OnchainData.InferTx = tx.Hash().Hex()
	chatCompletion.OnchainData.InferId = inferId
	chatCompletion.OnchainData.PbftCommittee = pBFTCommittee
	chatCompletion.OnchainData.Proposer = proposer
	chatCompletion.OnchainData.ProposeTx = ""
	return tx, &inferId, chatCompletion, nil
}

func (c *API_Local_Chain_V1) HealthCheck(ctx context.Context) (bool, error) {
	return true, nil
}

func (c *API_Local_Chain_V1) CreateInferWithStream(ctx context.Context, request model.LLMInferRequest, out chan model.StreamDataChannel) (*types.Transaction, *uint64, *model.LLMInferResponse, error) {
	cnf := c.cnf
	privKey := cnf.PrivateKey
	client := c.client
	var err error
	listeningOnChannel := pkg.STREAM_DATA_CHANNEL
	pubsub := c.rdb.Subscribe(listeningOnChannel)
	chatCompletion := &model.LLMInferResponse{}
	defer pubsub.Close()
	inferIdP := new(string)

	defer func() {
		fullMSG := ""
		existed := make(map[string]bool)
		format := "%s_%d"

		//TEST
		fmt.Println(pkg.PrintText("listening channel", listeningOnChannel))
		for {
			currentTime := time.Now().UTC()
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				continue
			}

			fmt.Printf("[%s] - Received message: %s\n", currentTime, msg.Payload)

			_b := []byte(msg.Payload)
			_dt := model.StreamingData{}

			err1 := json.Unmarshal(_b, &_dt)
			if err1 != nil {
				continue
			}

			_, ok := existed[fmt.Sprintf(format, _dt.InferenceID, _dt.StreamID)]
			if ok {
				continue
			}

			if inferIdP == nil {
				continue
			}

			if *inferIdP != _dt.InferenceID {
				continue
			}

			chatCompletion.Choices = make([]model.LLMInferChoice, len(_dt.Data.Choices))
			chatCompletion.Id = _dt.Data.Id
			chatCompletion.Object = _dt.Data.Object
			chatCompletion.Model = _dt.Data.Model
			chatCompletion.Created = _dt.Data.Created
			chatCompletion.IsStop = _dt.Stop
			for k, choice := range _dt.Data.Choices {
				chatCompletion.Choices[k].Message.Role = choice.Delta.Role
				chatCompletion.Choices[k].Message.Content = choice.Delta.Content
				fullMSG += choice.Delta.Content
			}

			_out := model.StreamDataChannel{
				Err:  err,
				Data: chatCompletion,
			}

			if chatCompletion != nil {
				_out.InferID = chatCompletion.OnchainData.InferId
			}

			out <- _out
			existed[fmt.Sprintf(format, _dt.InferenceID, _dt.StreamID)] = true
			if _dt.Stop {
				break
			}
			//time.Sleep(500 * time.Microsecond)
		}

		close(out)
	}()

	auth, err := eth.CreateBindTransactionOpts(ctx, c.client, privKey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		return nil, nil, nil, err
	}

	p := c.hybridModel
	_, _, err = eth.GetAccountInfo(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	_b, err := json.Marshal(request)
	if err != nil {
		return nil, nil, nil, err
	}

	tx, err := p.Infer(auth, _b, true)
	if err != nil {
		return nil, nil, nil, err
	}

	txReceipt, err := eth.WaitForTxReceipt(client, tx.Hash())
	if err != nil {
		err = errors.Join(err, errors.New("error while waiting for tx"))
		return nil, nil, nil, err
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

	inferId := inferIdBig.Uint64()
	inferIDStr := inferIdBig.String()
	inferIdP = &inferIDStr
	// wait for result

	pBFTCommittee := []string{}
	_pBFTCommittee := make(map[string]string)
	proposer := ""

	assignmentIDs, err := wkHub.GetAssignmentsByInference(nil, inferIdBig)
	if err != nil {
		return tx, &inferId, nil, err
	}

	for _, assismentID := range assignmentIDs {
		aInfo, err := wkHub.GetAssignmentInfo(nil, assismentID)
		if err != nil {
			return tx, &inferId, nil, err
		}
		_pBFTCommittee[aInfo.Worker.Hex()] = aInfo.Worker.Hex()

		if aInfo.Role == 2 {
			proposer = aInfo.Worker.Hex()
		}

	}

	for _, v := range _pBFTCommittee {
		pBFTCommittee = append(pBFTCommittee, v)
	}

	chatCompletion.OnchainData.InferTx = tx.Hash().Hex()
	chatCompletion.OnchainData.InferId = inferId
	chatCompletion.OnchainData.PbftCommittee = pBFTCommittee
	chatCompletion.OnchainData.Proposer = proposer
	chatCompletion.OnchainData.ProposeTx = ""
	return tx, &inferId, chatCompletion, nil
}

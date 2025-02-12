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
	"runtime"
	"strconv"
	"strings"
	"time"

	"solo/internal/contracts/erc20"
	"solo/internal/contracts/v1/hybrid_model"
	"solo/internal/contracts/v1/worker_hub"
	"solo/internal/contracts/w_eai"
	"solo/internal/model"
	"solo/internal/port"
	"solo/pkg"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type cmdLocalChainV1 struct {
	port.ICMDLocalChain
	gasPrice *big.Int
	gasLimit uint64

	rpc     string
	chainID string
	prvKey  string
	rdb     *redis.Client
}

func NewCMDLocalChainV1() port.ICMDLocalChainV1 {
	c := &cmdLocalChainV1{
		ICMDLocalChain: NewCMDLocalChainV2(),
	}

	c.gasPrice = big.NewInt(pkg.LOCAL_CHAIN_GAS_PRICE)
	c.gasLimit = pkg.LOCAL_CHAIN_GAS_LIMIT

	localCnf := c.ReadLocalChainCnf()
	if localCnf != nil {
		c.rpc = localCnf.Rpc
		c.chainID = localCnf.ChainID
		c.prvKey = localCnf.PrivateKey

		// Connect to the Redis server
		rdb := redis.NewClient(&redis.Options{
			Addr:     localCnf.PubSubURL,
			Password: "", // no password set
			DB:       0,  // default DB
		})

		c.rdb = rdb
	}
	return c
}

func (c *cmdLocalChainV1) DeployContracts(rpc, chainID, prvkey string) (*model.LocalChain, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Deploying contracts v1")

	cnf := c.ReadLocalChainCnf()

	env := "#HARDHAT\n"
	for i := 1; i <= 3; i++ {
		row := fmt.Sprintf("HARDHAT_PRIVATE_KEY_WORKER_%d=%s\n", i, cnf.PrivateKey)
		env += row
	}

	env += "HARDHAT_L2_OWNER_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65\n" // account 4 of Hardhat
	env += "HARDHAT_TREASURY_ADDRESS=\n"
	env += "HARDHAT_COLLECTION_ADDRESS=\n"
	env += "HARDHAT_GPU_MANAGER_ADDRESS=\n"
	env += "HARDHAT_PROMPT_SCHEDULER_ADDRESS=\n"
	env += "HARDHAT_DAGENT_721_ADDRESS=\n"
	env += "HARDHAT_MODEL_LOAD_BALANCER_ADDRESS=\n"
	env += "HARDHAT_WEAI=\n"

	if err := pkg.CreateFile(fmt.Sprintf(pkg.ENV_CONTRACT_V1_ENV, pkg.CurrentDir()), []byte(env)); err != nil {
		return nil, err
	}

	if err := c.ContractDeployment(); err != nil {
		return nil, err
	}

	_b, err := os.ReadFile(fmt.Sprintf(pkg.ENV_CONTRACT_V1_DEPLOYED_ADDRESS, pkg.CurrentDir()))
	if err != nil {
		return nil, err
	}

	resp := make(map[string]map[string]string)
	if err := json.Unmarshal(_b, &resp); err != nil {
		return nil, err
	}

	data, ok := resp["LOCALHOST"]
	if ok {
		cnf.Contracts = data
	}

	// save
	_b1, err := json.Marshal(cnf)
	if err != nil {
		return nil, err
	}

	if err = pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b1); err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *cmdLocalChainV1) ContractDeployment() error {
	_, ok := c.RpcHealthCheck()
	if !ok {
		err := errors.New("rpc is not started")
		return err
	}

	solV2Folder := fmt.Sprintf(pkg.ENV_CONTRACT_V1_PATH, pkg.CurrentDir())
	solV2AutoDeploy := fmt.Sprintf(pkg.ENV_CONTRACT_V1_AUTO_DEPLOY, pkg.CurrentDir())

	content := fmt.Sprintf(`#!/bin/bash

# Navigate to the project directory
cd %s || exit

# Run npm install and then the Hardhat script
npm install && npx hardhat run %s --network localhost`, solV2Folder, solV2AutoDeploy)

	fname := fmt.Sprintf(pkg.ENV_CONTRACT_V1_DEPLOY_SH, pkg.CurrentDir())
	err := pkg.CreateFile(fname, []byte(content))
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(fname)
	}()

	// Run the Hardhat script
	return pkg.CMDWithStream("bash", fname)
}

func (c *cmdLocalChainV1) DeployContractLogic() error {
	cnf := c.ReadLocalChainCnf()
	rpc := cnf.Rpc
	chainID := cnf.ChainID
	privKey := cnf.PrivateKey

	_, err := c.DeployContracts(rpc, chainID, privKey)
	if err != nil {
		fmt.Println("Deployed contracts error: ", err)
		return err
	}

	// 3. Mint WEAI.
	_, err = c.MintWrappedEAI(rpc, chainID, "100000", privKey)
	if err != nil {
		fmt.Println("Mint WEAI error: ", err)
		return err
	}

	return nil
}

func (c *cmdLocalChainV1) StartMinerLogic() error {
	fmt.Print(pkg.Line)
	fmt.Println("Start miners")

	cnf := c.ReadLocalChainCnf()
	rpc := cnf.Rpc
	chainID := cnf.ChainID
	privKey := cnf.PrivateKey

	fmt.Println("Create miners: ")
	numberOfMiners := 3
	names := ""

	// clear the created miners
	cnf.Miners = make(map[string]model.Miners)
	_b, err := json.Marshal(cnf)
	if err == nil {
		pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b)
	}

	for i := 1; i <= numberOfMiners; i++ {
		fmt.Print(pkg.Line)
		// 5. Create a miner's private key (3 miner)
		minerAddress, minerPrvKey, err := c.CreateMinerAddress(rpc, chainID, privKey)
		if err != nil {
			continue
		}

		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d address", i), *minerAddress))
		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d private key", i), *minerPrvKey))

		// send WEAI
		tx, _, err := c.SendWEIToMiner(rpc, *minerAddress)
		if err != nil {
			fmt.Println("SendWEIToMiner error", err)
			continue
		}
		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d received WEAI TX", i), tx.Hash().Hex()))

		// send fee
		txFee := new(types.Transaction)
		gas := pkg.LOCAL_CHAIN_GAS_LIMIT

		_loop := 1
		for {
			if _loop >= 50 {
				break
			}

			time.Sleep(time.Second * 2)
			txFee, _, err = c.SendFeeToMiner(rpc, *minerAddress, uint64(gas))
			if err != nil {
				if strings.Contains(err.Error(), "is too low for the next block, which has a baseFeePerGas of") {
					ar := strings.Split(err.Error(), " ")
					gasStr := ar[len(ar)-1]
					gasStrInt, errP := strconv.Atoi(gasStr)
					if errP == nil {
						gas = gasStrInt
					} else {
						gas += 10_000
					}
				} else {
					gas += 10_000
				}

				if strings.Contains(err.Error(), "and exceeds block gas limit of") {
					ar := strings.Split(err.Error(), " ")
					gasStr := ar[len(ar)-1]
					gasStrInt, errP := strconv.Atoi(gasStr)
					if errP == nil {
						gas = gasStrInt
					} else {
						gas += 10_000
					}
				} else {
					gas += 10_000
				}

				fmt.Print(pkg.PrintText("SendFeeToMiner error", err))
				continue
			}
			// there is no error
			fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d gas limit", i), gas))
			_loop++
			break
		}

		err = c.CreateConfigENV(*minerAddress, i)
		if err != nil {
			fmt.Println(fmt.Sprintf("Create config for miner %d error", i), err)
			gas += gas
			continue
		}

		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d received Fee TX", i), txFee.Hash().Hex()))

		name := fmt.Sprintf("%s_%d", pkg.MINER_SERVICE_NAME, i)
		names += " " + name
	}

	errBuild := c.BuildContainers(fmt.Sprintf("%s_base", pkg.MINER_SERVICE_NAME))
	if errBuild == nil {
		// i don't want to down all services
		if names != "" {
			cnf1 := c.ReadLocalChainCnf()
			pkg.DockerCommand(names, pkg.CurrentDir(), cnf1.Platform, "down", "-local")
		}

		c.StartContainersNoBuild(names)
	}

	return nil
}

func (c *cmdLocalChainV1) CreateConfigENV(minerAddress string, index int) error {
	// sample := fmt.Sprintf(pkg.ENV_SAMPLE_FILE, pkg.CurrentDir())
	envFile := fmt.Sprintf(pkg.ENV_LOCAL_MINERS_FILE, pkg.CurrentDir(), index)
	// config.ReadConfig(sample)

	f, _ := os.Stat(envFile)
	if f != nil {
		err := os.Remove(envFile)
		if err != nil {
			return err
		}
	}

	cnf := c.ReadLocalChainCnf()

	apiURL := cnf.RunPodInternal
	if strings.Contains(apiURL, "localhost") {
		_os := runtime.GOOS
		if _os == "darwin" {
			apiURL = strings.ReplaceAll(apiURL, "localhost", "host.docker.internal")
			fmt.Print(pkg.PrintText("OS", _os))
		}
	}

	env := ""
	env += fmt.Sprintf("PUBSUB_URL=%v\n", cnf.PubSubURL)
	env += fmt.Sprintf("PLATFORM=%v\n", cnf.Platform)
	env += fmt.Sprintf("API_URL=%v\n", apiURL)
	env += fmt.Sprintf("API_KEY=%v\n", cnf.RunPodAPIKEY)
	env += fmt.Sprintf("LIGHT_HOUSE_API_KEY=%v\n", os.Getenv("LIGHT_HOUSE_API_KEY"))
	env += fmt.Sprintf("CLUSTER_ID=%v\n", cnf.ModelID)
	env += fmt.Sprintf("MODEL_ID=%v\n", cnf.ModelID)
	env += fmt.Sprintf("CHAIN_ID=%v\n", cnf.ChainID)
	env += fmt.Sprintf("CHAIN_RPC=%v\n", fmt.Sprintf(`http://%s:8545`, "localhost"))
	env += fmt.Sprintf("ACCOUNT_PRIV=%v\n", cnf.Miners[strings.ToLower(minerAddress)].PrivateKey)
	env += fmt.Sprintf("MODEL_NAME=%v\n", cnf.ModelName)
	env += fmt.Sprintf("STAKING_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER_V1])
	env += fmt.Sprintf("MODEL_LOAD_BALANCER_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_load_balancer])
	env += fmt.Sprintf("WORKER_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER_V1])
	env += fmt.Sprintf("ERC20_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI_V1])
	env += fmt.Sprintf("COLLECTION_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION_V1])

	err := pkg.CreateFile(envFile, []byte(env))
	if err != nil {
		return err
	}
	return nil
}

func (c *cmdLocalChainV1) MintWrappedEAI(rpc, chainID, mintAmount, prvkey string) (*types.Transaction, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Minting WEAI")

	cnf := c.ReadLocalChainCnf()
	ctx := context.Background()

	client, err := c.NewClient(rpc)
	if err != nil {
		return nil, err
	}
	weaiAddress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI_V1]
	if !ok {
		err := errors.New("MintWrappedEAI: weai contract was not deployed")
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, err
	}

	contract, err := w_eai.NewWEai(common.HexToAddress(weaiAddress), client)
	if err != nil {
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, err
	}

	amount := big.NewInt(1).Mul(big.NewInt(1e18), big.NewInt(5_000_000))

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit

	_, address, err := eth.GetAccountInfo(prvkey)
	if err != nil {
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, err
	}

	// fmt.Println(address.Hex())
	tx, err := contract.Mint(auth, *address, amount)
	if err != nil {
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, err
	}

	fmt.Print(pkg.PrintText("MintWrappedEAI tx", tx.Hash().Hex()))
	return tx, nil
}

func (c *cmdLocalChainV1) MintCollection(rpc, prvkey string, modelName string) (*types.Transaction, *big.Int, error) {
	return nil, nil, nil
}

func (c *cmdLocalChainV1) SendWEIToMiner(rpc, minerAddress string) (*types.Transaction, *string, error) {
	cnf := c.ReadLocalChainCnf()
	prvkey := cnf.PrivateKey
	ctx := context.Background()

	client, err := c.NewClient(rpc)
	if err != nil {
		return nil, nil, err
	}
	weaiAddress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI_V1]
	if !ok {
		err := errors.New("CreateMinerAddress: weai contract was not deployed")
		return nil, nil, err
	}

	// 6. Owner transfer to miner 25k EAI for staking
	erc20Contract, err := erc20.NewErc20(common.HexToAddress(weaiAddress), client)
	if err != nil {
		fmt.Println("CreateMinerAddress error: ", err)
		return nil, nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		fmt.Println("CreateMinerAddress error: ", err)
		return nil, nil, err
	}

	value := big.NewInt(1).Mul(big.NewInt(pkg.MIN_STAKE), big.NewInt(1e18))
	transferTX, err := erc20Contract.Transfer(auth, common.HexToAddress(minerAddress), value)
	if err != nil {
		fmt.Println("CreateMinerAddress error: ", err)
		return nil, nil, err
	}

	return transferTX, &minerAddress, nil
}

func (c *cmdLocalChainV1) CreateInfer(prompt []model.LLMInferMessage) (*types.Transaction, *uint64, *string, error) {
	ctx := context.Background()
	cnf := c.ReadLocalChainCnf()
	privKey := cnf.PrivateKey
	client, err := eth.NewEthClient(cnf.Rpc)
	if err != nil {
		return nil, nil, nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, privKey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		return nil, nil, nil, err
	}

	p, err := hybrid_model.NewHybridModel(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1]), client)
	if err != nil {
		return nil, nil, nil, err
	}

	_, pubkey, err := eth.GetAccountInfo(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	request := model.LLMInferRequest{
		Model:    cnf.ModelName,
		Messages: prompt,
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

	wkHub, err := worker_hub.NewWorkerHub(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER_V1]), client)
	if err != nil {
		return nil, nil, nil, err
	}

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

	// if chatCompletion == nil {
	// 	return tx, &inferId, nil, errors.New("error while parse response")
	// }

	if len(chatCompletion.Choices) == 0 {
		return tx, &inferId, nil, errors.New("error get data")
	}

	_ = txReceipt
	_ = receipt
	_ = pubkey

	return tx, &inferId, &chatCompletion.Choices[0].Message.Content, nil
}

func (c *cmdLocalChainV1) CreateInferWithStream(prompt []model.LLMInferMessage, out chan model.StreamDataChannel) (*types.Transaction, *uint64, *string, error) {
	ctx := context.Background()
	cnf := c.ReadLocalChainCnf()
	privKey := cnf.PrivateKey
	chatCompletion := &model.LLMInferResponse{}
	pubsub := c.rdb.Subscribe(pkg.STREAM_DATA_CHANNEL)

	defer pubsub.Close()

	client, err := eth.NewEthClient(cnf.Rpc)
	if err != nil {
		return nil, nil, nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, privKey, pkg.LOCAL_CHAIN_GAS_LIMIT)
	if err != nil {
		return nil, nil, nil, err
	}

	p, err := hybrid_model.NewHybridModel(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_HYBRID_MODEL_V1]), client)
	if err != nil {
		return nil, nil, nil, err
	}

	_, pubkey, err := eth.GetAccountInfo(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	request := model.LLMInferRequest{
		Model:    cnf.ModelName,
		Messages: prompt,
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

	wkHub, err := worker_hub.NewWorkerHub(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER_V1]), client)
	if err != nil {
		return nil, nil, nil, err
	}

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
	func(pubsub *redis.PubSub) {
		existed := make(map[string]bool)
		format := "%s_%d"

		//PUSH data to the stream
		fullMSG := ""
		for {
			//currentTime := time.Now().UTC()
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				continue
			}

			//fmt.Printf("[%s] - Received message: %s\n", "INFO", msg.Payload)

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

			if inferIdBig.String() != _dt.InferenceID {
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
	}(pubsub)

	// if chatCompletion == nil {
	// 	return tx, &inferId, nil, errors.New("error while parse response")
	// }

	if len(chatCompletion.Choices) == 0 {
		return tx, &inferId, nil, errors.New("error get data")
	}

	_ = txReceipt
	_ = receipt
	_ = pubkey

	return tx, &inferId, &chatCompletion.Choices[0].Message.Content, nil
}

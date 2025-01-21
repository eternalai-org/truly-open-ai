package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"solo/internal/contracts/erc20"
	"solo/internal/contracts/gpu_manager"
	"solo/internal/contracts/load_balancer"
	"solo/internal/contracts/model_collection"
	"solo/internal/contracts/prompt_scheduler"
	"solo/internal/contracts/proxy"
	"solo/internal/contracts/w_eai"
	"solo/internal/model"
	"solo/pkg"
	"solo/pkg/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CMD_Local_Chain struct {
	gasPrice *big.Int
	gasLimit uint64

	rpc     string
	chainID string
	prvKey  string
}

func NewCMDLocalChain() (*CMD_Local_Chain, error) {
	gasPrice := big.NewInt(pkg.LOCAL_CHAIN_GAS_PRICE)
	c := &CMD_Local_Chain{
		gasPrice: gasPrice,
		gasLimit: pkg.LOCAL_CHAIN_GAS_LIMIT,
	}
	localCnf := c.ReadLocalChainCnf()
	if localCnf != nil {
		c.rpc = localCnf.Rpc
		c.chainID = localCnf.ChainID
		c.prvKey = localCnf.PrivateKey
	}

	return c, nil
}

func (c *CMD_Local_Chain) GetPrivateKey() string {
	return c.prvKey
}

func (c *CMD_Local_Chain) SetGasPrice(gp *big.Int) {
	c.gasPrice = gp
}

func (c *CMD_Local_Chain) SetGasLimit(gl uint64) {
	c.gasLimit = gl
}

func (c *CMD_Local_Chain) newClient(rpc string) (*ethclient.Client, error) {
	client, err := eth.NewEthClient(rpc)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *CMD_Local_Chain) BuildMiner(minerIndex int) error {
	folderPath := pkg.CurrentDir()
	return pkg.DockerCommand(fmt.Sprintf("%s_%d", pkg.MINER_SERVICE_NAME, minerIndex), folderPath, os.Getenv("PLATFORM"), "build", "-local")
}

func (c *CMD_Local_Chain) BuildContainers(names string) error {
	folderPath := pkg.CurrentDir()
	pkg.DockerCommand(names, folderPath, os.Getenv("PLATFORM"), "down", "-local")
	return pkg.DockerCommand(names, folderPath, os.Getenv("PLATFORM"), "build", "-local")
}

func (c *CMD_Local_Chain) StartContainers(names string) error {
	folderPath := pkg.CurrentDir()
	return pkg.DockerCommand(names, folderPath, os.Getenv("PLATFORM"), "up -d", "-local")
}

func (c *CMD_Local_Chain) StartContainersNoBuild(names string) error {
	folderPath := pkg.CurrentDir()
	return pkg.DockerCommand(names, folderPath, os.Getenv("PLATFORM"), "up -d --no-build", "-local")
}

func (c *CMD_Local_Chain) DeployContracts(rpc, chainID, prvkey string) (*model.LocalChain, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Deploying contracts")

	cnf := c.ReadLocalChainCnf()

	env := "#HARDHAT\n"
	for i := 1; i <= 3; i++ {
		row := fmt.Sprintf("HARDHAT_PRIVATE_KEY_WORKER_%d=%s\n", i, cnf.PrivateKey)
		env += row
	}

	env += "HARDHAT_TREASURY_ADDRESS=\n"
	env += "HARDHAT_COLLECTION_ADDRESS=\n"
	env += "HARDHAT_GPU_MANAGER_ADDRESS=\n"
	env += "HARDHAT_PROMPT_SCHEDULER_ADDRESS=\n"
	env += "HARDHAT_DAGENT_721_ADDRESS=\n"
	env += "HARDHAT_MODEL_LOAD_BALANCER_ADDRESS=\n"
	env += "HARDHAT_WEAI=\n"

	err := pkg.CreateFile(fmt.Sprintf(pkg.ENV_CONTRACT_ENV, pkg.CurrentDir()), []byte(env))
	if err != nil {
		return nil, err
	}

	err = c.ContractDeployment()
	if err != nil {
		return nil, err
	}

	_b, err := os.ReadFile(fmt.Sprintf(pkg.ENV_CONTRACT_DEPLOYED_ADDRESS, pkg.CurrentDir()))
	if err != nil {
		return nil, err
	}

	resp := make(map[string]map[string]string)
	err = json.Unmarshal(_b, &resp)
	if err != nil {
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

	err = pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b1)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *CMD_Local_Chain) DeployContract(rpc, chainID, prvkey, contractName string) (*model.LocalChain, error) {
	var err error
	client, err := c.newClient(rpc)
	if err != nil {
		return nil, err
	}
	cname := &common.Address{}

	resp := c.ReadLocalChainCnf()
	resp.Rpc = rpc
	resp.ChainID = chainID
	resp.PrivateKey = prvkey

	switch contractName {
	case pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI:
		cname, _, err = c.DeployContractWrappedEAI(client, prvkey)
	case pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION:
		cname, _, err = c.DeployContractModelCollection(client, prvkey)
	case pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER:
		cname, _, err = c.DeployContractPromptScheduler(client, prvkey)
	case pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER:
		cname, _, err = c.DeployContractGpuManager(client, prvkey)
	case pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_load_balancer:
		cname, _, err = c.DeployContractLoadBalancer(client, prvkey)

	}

	if err != nil {
		return nil, err
	}
	resp.Contracts[contractName] = cname.Hex()
	_b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b)
	return resp, nil
}

func (c *CMD_Local_Chain) DeployContractPromptScheduler(client *ethclient.Client, prvkey string) (*common.Address, *types.Transaction, error) {
	// deploy contracts here
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Println(pkg.PrintText("Prompt scheduler was deployed with err: ", err))
		return nil, nil, err
	}

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit

	contractAddress, tx, _p, err := prompt_scheduler.DeployPromptScheduler(auth, client)
	if err != nil {
		fmt.Println(pkg.PrintText("Prompt scheduler was deployed with err: ", err))
		return nil, nil, err
	}

	pContract, _, err := c.DeployProxy(ctx, client, prvkey, contractAddress)
	if err != nil {
		fmt.Println(pkg.PrintText("Prompt scheduler was deployed with err: ", err))
		return nil, nil, err
	}
	_ = _p

	fmt.Print(pkg.PrintText("Prompt scheduler address: ", pContract.Hex()))
	return pContract, tx, nil
}

func (c *CMD_Local_Chain) DeployContractGpuManager(client *ethclient.Client, prvkey string) (*common.Address, *types.Transaction, error) {
	// deploy contracts here
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Println(pkg.PrintText("Gpu manager was deployed with err: ", err))
		return nil, nil, err
	}

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	contractAddress, tx, _, err := gpu_manager.DeployGpuManager(auth, client)
	if err != nil {
		fmt.Print(pkg.PrintText("Gpu manager was deployed with err: ", err))
		return nil, nil, err
	}

	// fmt.Print(pkg.PrintText("Gpu manager owner address: ", ownerStr))
	pContract, _, err := c.DeployProxy(ctx, client, prvkey, contractAddress)
	if err != nil {
		fmt.Println(pkg.PrintText("Gpu manager was deployed with err: ", err))
		return nil, nil, err
	}

	fmt.Print(pkg.PrintText("Gpu manager address: ", pContract.Hex()))
	return pContract, tx, nil
}

func (c *CMD_Local_Chain) DeployContractModelCollection(client *ethclient.Client, prvkey string) (*common.Address, *types.Transaction, error) {
	//deploy contracts here
	/*
		ctx := context.Background()
		auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
		if err != nil {
			fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
			return nil, nil, err
		}

		auth.GasPrice = c.gasPrice
		auth.GasLimit = c.gasLimit
		contractAddress, tx, _p, err := model_collection.DeployModelCollection(auth, client)
		if err != nil {
			fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
			return nil, nil, err
		}
		_ = _p

		//fmt.Print(pkg.PrintText("Gpu manager owner address: ", ownerStr))
		pContract, tx, err := c.DeployProxy(ctx, client, prvkey, contractAddress)
		if err != nil {
			fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
			return nil, nil, err
		}

		//initialize
		cnf := c.ReadLocalChainCnf()

		auth, err = eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
		if err != nil {
			fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
			return nil, nil, err
		}

		_, receiverAddress, _ := eth.GetAccountInfo(prvkey)
		contract, err := model_collection.NewModelCollection(*pContract, client)
		if err != nil {
			fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
			return nil, nil, err
		}

		fmt.Print(pkg.PrintText("Model collection: ", pContract.Hex()))
		wEAI := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI])
		intializeTX, err := contract.Initialize(auth, "model-collection", "WEAI", big.NewInt(1e16), *receiverAddress, 0, big.NewInt(100), wEAI)
		if err != nil {
			fmt.Print(pkg.PrintText("Model collection initialize with err: ", err))
			return nil, nil, err
		}

		fmt.Print(pkg.PrintText("Model collection Initialize tx: ", intializeTX))*/

	cnf := c.ReadLocalChainCnf()
	wEAI := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI])
	address, tx, err := eth.Deploy(client, prvkey, model_collection.ModelCollectionABI, model_collection.ModelCollectionBin, "model-collection", "WEAI", big.NewInt(1e16), common.Address{}, 0, big.NewInt(100), wEAI)
	if err != nil {
		return nil, nil, err
	}

	contract, err := model_collection.NewModelCollection(address, client)
	if err != nil {
		return nil, nil, err
	}

	s, err := contract.Symbol(nil)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("====>", s)
	return &address, tx, nil
}

func (c *CMD_Local_Chain) DeployContractLoadBalancer(client *ethclient.Client, prvkey string) (*common.Address, *types.Transaction, error) {
	// deploy contracts here
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Println(pkg.PrintText("Load balancer was deployed with err: ", err))
		return nil, nil, err
	}

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	contractAddress, tx, _p, err := load_balancer.DeployLoadBalancer(auth, client)
	if err != nil {
		fmt.Println(pkg.PrintText("Load balancer was deployed with err: ", err))
		return nil, nil, err
	}

	_ = _p
	pContract, _, err := c.DeployProxy(ctx, client, prvkey, contractAddress)
	if err != nil {
		fmt.Println(pkg.PrintText("Load balancer was deployed with err: ", err))
		return nil, nil, err
	}

	fmt.Print(pkg.PrintText("Load balancer address: ", pContract.Hex()))

	// initialize
	cnf := c.ReadLocalChainCnf()

	auth, err = eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
		return nil, nil, err
	}

	contract, err := load_balancer.NewLoadBalancer(*pContract, client)
	if err != nil {
		fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
		return nil, nil, err
	}

	// prompt
	pC := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER])
	wEAI := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI])
	modelCollectionC := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION])
	intializeTX, err := contract.Initialize(auth, *pContract, pC, wEAI, modelCollectionC, big.NewInt(1))
	if err != nil {
		fmt.Print(pkg.PrintText("Load balancer with err: ", err))
		// return nil, nil, err
	} else {
		fmt.Print(pkg.PrintText("Initialize tx: ", intializeTX))
	}
	_ = contract
	_ = cnf

	return pContract, tx, nil
}

func (c *CMD_Local_Chain) DeployProxy(ctx context.Context, client *ethclient.Client, prvkey string, contractAddress common.Address) (*common.Address, *types.Transaction, error) {
	// re-auth
	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		return nil, nil, err
	}

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	pContractAddress, tx, _, err := proxy.DeployProxy(auth, client, contractAddress)
	if err != nil {
		return nil, nil, err
	}

	return &pContractAddress, tx, nil
}

func (c *CMD_Local_Chain) DeployContractWrappedEAI(client *ethclient.Client, prvkey string) (*common.Address, *types.Transaction, error) {
	// deploy contracts here
	ctx := context.Background()
	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Println(pkg.PrintText("Wrapped EAI was deployed with err: ", err))
		return nil, nil, err
	}

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	contractAddress, tx, _p, err := w_eai.DeployWEai(auth, client)
	if err != nil {
		fmt.Println(pkg.PrintText("Wrapped EAI was deployed with err: ", err))
		return nil, nil, err
	}

	_ = _p
	err = eth.WaitForTx(client, tx.Hash())
	if err != nil {
		fmt.Println(pkg.PrintText("Model collection was deployed with err: ", err))
		return nil, nil, err
	}

	fmt.Print(pkg.PrintText("Wrapped EAI address: ", contractAddress.Hex()))
	return &contractAddress, tx, nil
}

func (c *CMD_Local_Chain) MintWrappedEAI(rpc, chainID, mintAmount, prvkey string) (*types.Transaction, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Minting WEAI")

	cnf := c.ReadLocalChainCnf()
	ctx := context.Background()

	client, err := c.newClient(rpc)
	if err != nil {
		return nil, err
	}
	weaiAddress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI]
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

func (c *CMD_Local_Chain) MintCollection(rpc, prvkey string, modelName string) (*types.Transaction, *big.Int, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Minting Collection")

	cnf := c.ReadLocalChainCnf()
	ctx := context.Background()

	client, err := c.newClient(rpc)
	if err != nil {
		return nil, nil, err
	}
	modelCollectionAddrress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION]
	if !ok {
		err := errors.New("MintCollection: weai contract was not deployed")
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Print(pkg.PrintText("MintCollection with err: ", err))
		return nil, nil, err
	}

	contract, err := model_collection.NewModelCollection(common.HexToAddress(modelCollectionAddrress), client)
	if err != nil {
		fmt.Print(pkg.PrintText("MintCollection with err: ", err))
		return nil, nil, err
	}

	_, address, err := eth.GetAccountInfo(prvkey)
	if err != nil {
		fmt.Print(pkg.PrintText("MintWrappedEAI with err: ", err))
		return nil, nil, err
	}

	clusterMetaData := model.ClusterMetaData{
		Version:     1,
		ModelName:   modelName,
		ModelType:   "text",
		MinHardware: 1,
	}
	metadata, err := json.MarshalIndent(clusterMetaData, "", "\t")
	if err != nil {
		fmt.Print(pkg.PrintText("MintCollection with err: ", err))
		return nil, nil, err
	}

	auth, err = eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Print(pkg.PrintText("MintCollection with err: ", err))
		return nil, nil, err
	}
	tx, err := contract.Mint(auth, *address, string(metadata))
	if err != nil {
		fmt.Print(pkg.PrintText("MintCollection with err: ", err))
		return nil, nil, err
	}

	err = eth.WaitForTx(client, tx.Hash())
	if err != nil {
		return tx, nil, err
	}

	tokenID, err := eth.GetTokenIDFromTx(client, tx.Hash())
	if err != nil {
		return tx, nil, err
	}

	fmt.Print(pkg.PrintText("MintCollection tx", tx.Hash().Hex()))
	fmt.Print(pkg.PrintText("MintCollection collectionID", tokenID.String()))

	cnf.ModelName = modelName
	cnf.ModelID = tokenID.String()

	_b, err := json.Marshal(cnf)
	if err != nil {
		return nil, nil, err
	}
	pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b)
	return tx, tokenID, nil
}

func (c *CMD_Local_Chain) SetWEAIForStakingHub(client *ethclient.Client, prvkey string) (*types.Transaction, error) {
	cnf := c.ReadLocalChainCnf()
	ctx := context.Background()

	weaiAddress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI]
	if !ok {
		err := errors.New("SetWEAIForStakingHub: weai contract was not deployed")
		return nil, err
	}

	gpuMangerContract, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER]
	if !ok {
		err := errors.New("SetWEAIForStakingHub: GPU Manager contract was not deployed")
		return nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, 0)
	if err != nil {
		fmt.Print(pkg.PrintText("SetWEAIForStakingHub with err: ", err))
		return nil, err
	}

	contract, err := gpu_manager.NewGpuManager(common.HexToAddress(gpuMangerContract), client)
	if err != nil {
		fmt.Print(pkg.PrintText("SetWEAIForStakingHub with err: ", err))
		return nil, err
	}

	modelCollection := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION])
	tresury := common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION]) // ever address is allowed for testing
	minstake := big.NewInt(1).Mul(big.NewInt(pkg.MIN_STAKE), big.NewInt(1e18))
	blockPerEpoch := big.NewInt(pkg.BLOCK_PER_EPOCH)
	rewardPerEpoch := big.NewInt(1).Mul(big.NewInt(pkg.REWARD_PER_EPOCH), big.NewInt(1e18))
	unstakeDelayTime := big.NewInt(pkg.UNSTAK_DEPLAY_TIME)
	penaltyDuration := big.NewInt(pkg.PENALTY_DURATION)
	finePercentage := uint16(pkg.FINE_PERCENTAGE)
	minFeeToUse := big.NewInt(pkg.MIN_FEE_TO_USE)

	auth.GasPrice = c.gasPrice
	auth.GasLimit = c.gasLimit
	tx, err := contract.Initialize(auth, common.HexToAddress(weaiAddress), modelCollection, tresury, minstake, blockPerEpoch, rewardPerEpoch, unstakeDelayTime, penaltyDuration, finePercentage, minFeeToUse)
	if err != nil {
		fmt.Print(pkg.PrintText("SetWEAIForStakingHub with err: ", err))
		return nil, err
	}

	fmt.Print(pkg.PrintText("SetWEAIForStakingHub tx", tx.Hash().Hex()))
	return tx, nil
}

func (c *CMD_Local_Chain) CreateMinerAddress(rpc, chainID, prvkey string) (*string, *string, error) {
	cnf := c.ReadLocalChainCnf()

	minerPrivateKey, _, minerAddress, err := eth.GenerateAddress()
	if err != nil {
		fmt.Println("CreateMinerAddress error: ", err)
		return nil, nil, err
	}

	// send ETH for fee
	// auth.GasPrice = big.NewInt(pkg.LOCAL_CHAIN_GAS_PRICE)
	// auth.GasLimit = pkg.LOCAL_CHAIN_GAS_LIMIT

	// update env
	cnf.Miners[strings.ToLower(minerAddress)] = model.Miners{
		Address:    strings.ToLower(minerAddress),
		PrivateKey: minerPrivateKey,
	}

	_b, err := json.Marshal(cnf)
	if err != nil {
		return nil, nil, err
	}

	cnf.ChainID = chainID
	cnf.PrivateKey = prvkey
	pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b)

	return &minerAddress, &minerPrivateKey, nil
}

func (c *CMD_Local_Chain) SendWEIToMiner(rpc, minerAddress string) (*types.Transaction, *string, error) {
	cnf := c.ReadLocalChainCnf()
	prvkey := cnf.PrivateKey
	ctx := context.Background()

	client, err := c.newClient(rpc)
	if err != nil {
		return nil, nil, err
	}
	weaiAddress, ok := cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI]
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

func (c *CMD_Local_Chain) SendFeeToMiner(rpc, minerAddress string, gasLimit uint64) (*types.Transaction, *string, error) {
	cnf := c.ReadLocalChainCnf()
	ctx := context.Background()
	prvkey := cnf.PrivateKey

	client, err := c.newClient(rpc)
	if err != nil {
		return nil, nil, err
	}

	sendTX, err := eth.SendToken(ctx, prvkey, minerAddress, big.NewInt(1e18), client, int64(gasLimit))
	if err != nil {
		return nil, nil, err
	}

	return sendTX, &minerAddress, nil
}

func (c *CMD_Local_Chain) ReadLocalChainCnf() *model.LocalChain {
	resp := new(model.LocalChain)
	resp.Contracts = make(map[string]string)
	resp.Miners = make(map[string]model.Miners)
	_b, err := os.ReadFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()))
	if err != nil {
		return resp
	}

	err = json.Unmarshal(_b, resp)
	if err != nil {
		return resp
	}

	return resp
}

func (c *CMD_Local_Chain) CreateInfer(prompt []model.LLMInferMessage) (*types.Transaction, *uint64, *string, error) {
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

	p, err := prompt_scheduler.NewPromptScheduler(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER]), client)
	if err != nil {
		return nil, nil, nil, err
	}

	_, pubkey, err := eth.GetAccountInfo(privKey)
	if err != nil {
		return nil, nil, nil, err
	}

	modelID := cnf.ModelID
	modelIDInt, _ := strconv.Atoi(modelID)
	request := model.LLMInferRequest{
		Model:    cnf.ModelName,
		Messages: prompt,
	}
	_b, err := json.Marshal(request)
	if err != nil {
		return nil, nil, nil, err
	}

	tx, err := p.Infer(auth, uint32(modelIDInt), _b, *pubkey, true)
	if err != nil {
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
	logs := receipt.Logs
	var inferId uint64
	for _, item := range logs {
		inferData, err := p.ParseNewInference(*item)
		if err == nil {
			inferId = inferData.InferenceId
		}
	}

	// wait for result
	chatCompletion := &model.LLMInferResponse{}
	index := 0
	for index < 150 {
		time.Sleep(2 * time.Second)
		infer, err := p.GetInferenceInfo(nil, inferId)
		if err != nil {
			return tx, &inferId, nil, err
		}

		out := string(infer.Output)

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
			break
		}
		index += 1
	}

	if chatCompletion == nil {
		return tx, &inferId, nil, errors.New("error while parse response")
	}

	if len(chatCompletion.Choices) == 0 {
		return tx, &inferId, nil, errors.New("error get data")
	}

	return tx, &inferId, &chatCompletion.Choices[0].Message.Content, nil
}

func (c *CMD_Local_Chain) SetGPUAddressRegisterModel(rpc string, modelID uint32, prvkey string) (*types.Transaction, error) {
	fmt.Print(pkg.Line)
	fmt.Println("Register modelID")

	ctx := context.Background()

	client, err := eth.NewEthClient(rpc)
	if err != nil {
		return nil, err
	}

	cnf := c.ReadLocalChainCnf()
	gpuContract, err := gpu_manager.NewGpuManager(common.HexToAddress(cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER]), client)
	if err != nil {
		return nil, err
	}

	auth, err := eth.CreateBindTransactionOpts(ctx, client, prvkey, int64(c.gasLimit))
	if err != nil {
		return nil, err
	}

	tier := uint16(1)
	minimunFee := big.NewInt(0)
	tx, err := gpuContract.RegisterModel(auth, modelID, tier, minimunFee)
	if err != nil {
		return nil, err
	}

	err = eth.WaitForTx(client, tx.Hash())
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *CMD_Local_Chain) StartHardHat() error {
	// err := pkg.DockerCommand(pkg.MINER_SERVICE_HARDHAT, pkg.CurrentDir(), "", "rm -f", "-local")
	// if err != nil {
	// return err
	// }

	err := pkg.DockerCommand(pkg.MINER_SERVICE_HARDHAT, pkg.CurrentDir(), "", "down", "-local")
	if err != nil {
		return err
	}

	err = pkg.DockerCommand(pkg.MINER_SERVICE_HARDHAT, pkg.CurrentDir(), "", "up -d --build", "-local")
	if err != nil {
		return err
	}

	c.PingRpc()
	return nil
}

func (c *CMD_Local_Chain) StartOllama() error {
	cnf := c.ReadLocalChainCnf()
	modelName := cnf.ModelName
	if modelName == "" {
		cnfFile := fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir())
		err := errors.New("model_name is empty, please goto `Setup > Manual>1` to  update your: " + cnfFile)
		return err
	}

	if cnf.RunPodExternal == "" {
		cnfFile := fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir())
		err := errors.New("run pod is empty, please `Setup > Manual>1` to update your: " + cnfFile)
		return err
	}

	if !cnf.UseExternalRunPod {

		// err := pkg.DockerCommand(pkg.MINER_SERVICE_OLLAMA, pkg.CurrentDir(), "", "rm -f", "-local")
		// if err != nil {
		// 	return err
		// }

		entrypoint := fmt.Sprintf(`#!/bin/bash
# Start Ollama in the background.
/bin/ollama serve &
# Record Process ID.
pid=$!

# Pause for Ollama to start.
sleep 5

echo "🔴 Retrieve %s model..."
ollama run %s
echo "🟢 Done!"

# Wait for Ollama process to finish.
wait $pid`, modelName, modelName)

		entryPointPath := fmt.Sprintf(pkg.ENTRY_POINT_FILE, pkg.CurrentDir())

		err := pkg.CreateFile(entryPointPath, []byte(entrypoint))
		if err != nil {
			return err
		}

		err = pkg.DockerCommand(pkg.MINER_SERVICE_OLLAMA, pkg.CurrentDir(), "", "down", "-local")
		if err != nil {
			return err
		}

		err = pkg.DockerCommand(pkg.MINER_SERVICE_OLLAMA, pkg.CurrentDir(), "", "up -d --build", "-local")
		if err != nil {
			return err
		}

	}

	// health check ollama
	c.PingOllam()

	return nil
}

func (c *CMD_Local_Chain) StartMiner(minerIndex int) error {
	return pkg.DockerCommand(fmt.Sprintf("%s_%d", pkg.MINER_SERVICE_NAME, minerIndex), pkg.CurrentDir(), "", "up -d", "-local")
}

func (c *CMD_Local_Chain) ContractDeployment() error {
	_, ok := c.RpcHealthCheck()
	if !ok {
		err := errors.New("rpc is not started")
		return err
	}

	content := fmt.Sprintf(`#!/bin/bash

# Navigate to the project directory
cd %s/sol || exit

# Run npm install and then the Hardhat script
npm install && npx hardhat run %s/sol/scripts/auto_deploy.ts --network localhost`, pkg.CurrentDir(), pkg.CurrentDir())

	fname := fmt.Sprintf("%s/sol/deploy.sh", pkg.CurrentDir())

	err := pkg.CreateFile(fname, []byte(content))
	if err != nil {
		return err
	}

	// Run the Hardhat script
	err = pkg.CMDWithStream("bash", fname)
	if err != nil {
		return err
	}

	return nil
}

func (c *CMD_Local_Chain) OllamaHealthCheck() ([]byte, bool) {
	cnf := c.ReadLocalChainCnf()

	url := cnf.RunPodExternal
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "1"

	request := model.LLMInferRequest{
		Model:    cnf.ModelName,
		MaxToken: 5,
		Messages: []model.LLMInferMessage{
			{
				Role:    "user",
				Content: "hello",
			},
		},
	}

	_b, _, st, err := pkg.HttpRequest(url, "POST", headers, request)
	if err != nil {
		return nil, false
	}
	if st != 200 {
		return nil, false
	}

	// time.Sleep(5 * time.Second)
	return _b, true
}

func (c *CMD_Local_Chain) PingOllam() model.LLMInferResponse {
	isReady := false
	_resp := []byte{}
	ping := 1
	for {

		if ping >= 1000 {
			return model.LLMInferResponse{}
		}

		fmt.Print(pkg.PrintText("OLLAM", fmt.Sprintf("ping (%d)...", ping)))
		_resp, isReady = c.OllamaHealthCheck()
		if isReady {
			fmt.Print(pkg.PrintText("OLLAM", "READY!!!!"))
			isReady = true
			break
		}

		fmt.Print(pkg.PrintText("OLLAM", "not ready"))
		time.Sleep(5 * time.Second)
		ping++
	}

	resp := model.LLMInferResponse{}
	err := json.Unmarshal(_resp, &resp)
	if err != nil {
		return model.LLMInferResponse{}
	}

	if len(resp.Choices) > 0 {
		msg := resp.Choices[0].Message.Content
		fmt.Print(pkg.PrintText("Ollam resp: ", msg))
	}

	return resp
}

func (c *CMD_Local_Chain) RpcHealthCheck() ([]byte, bool) {
	cnf := c.ReadLocalChainCnf()

	url := cnf.Rpc
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	type reqHealthCheck struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		Id      int           `json:"id"`
	}

	req := reqHealthCheck{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      1,
	}

	_b, _, st, err := pkg.HttpRequest(url, "POST", headers, req)
	if err != nil {
		return nil, false
	}
	if st != 200 {
		return nil, false
	}

	// time.Sleep(5 * time.Second)
	return _b, true
}

func (c *CMD_Local_Chain) PingRpc() interface{} {
	isReady := false
	_resp := []byte{}
	ping := 1
	for {

		if ping >= 1000 {
			return nil
		}

		fmt.Print(pkg.PrintText("RPC", fmt.Sprintf("ping (%d)...", ping)))
		_resp, isReady = c.RpcHealthCheck()
		if isReady {
			fmt.Print(pkg.PrintText("RPC", "READY!!!!"))
			isReady = true
			break
		}

		fmt.Print(pkg.PrintText("RPC", "not ready"))
		time.Sleep(5 * time.Second)
		ping++
	}

	fmt.Println("RPC resp: ", string(_resp))
	return _resp
}

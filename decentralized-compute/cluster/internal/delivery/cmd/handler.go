package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"solo/config"
	"solo/internal/model"
	"solo/pkg"
	"solo/pkg/eth"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *CMD) cliCommand() []*pkg.Command {
	localContractCMDs := []*pkg.Command{
		{
			Key:      pkg.COMMAND_LOCAL_PRIV_KEY,
			Help:     "Private Key",
			Default:  c.localChainCMD.GetPrivateKey(),
			Required: true,
		},
		{
			Key:     pkg.COMMAND_LOCAL_RUN_POD_URL,
			Help:    "Runpod URL (Default: empty)",
			Default: "",
		},
		{
			Key:     pkg.COMMAND_LOCAL_RUN_POD_API_KEY,
			Help:    "Runpod API-Key (Default: empty)",
			Default: "",
		},
		{
			Key:      pkg.COMMAND_LOCAL_MODEL_NAME,
			Help:     "Model name",
			Default:  "hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q3_K_S",
			Required: true,
		},
		/*
				{
					Key:      pkg.COMMAND_LOCAL_CHAIN_RPC,
					Help:     "Chain RPC",
					Required: true,
					Default:  "http://localhost:8545",
				},
				{
					Key:     pkg.COMMAND_LOCAL_CHAIN_ID,
					Help:    "Chain ID",
					Default: "31337",
				},
			{
				Key:     pkg.COMMAND_LOCAL_GAS_PRICE,
				Help:    "Custom gas price",
				Default: fmt.Sprintf("%d", pkg.LOCAL_CHAIN_GAS_PRICE),
			},
			{
				Key:     pkg.COMMAND_LOCAL_GAS_LIMIT,
				Help:    "Custom gas limit",
				Default: fmt.Sprintf("%d", pkg.LOCAL_CHAIN_GAS_LIMIT),
			},*/
	}
	spt := ""
	for _, v := range pkg.SupportedContracts {
		spt += pkg.PrintText(v, "")
	}
	setupCommands := []*pkg.Command{
		{
			Key:  pkg.COMMAND_SETUP,
			Help: `1. Setup local cluster`,
			Name: "Setup",
			Children: []*pkg.Command{
				{
					Key:      pkg.COMMAND_SETUP_AUTOMATIC,
					Help:     `This command sets up everything automatically in one step.`,
					Children: localContractCMDs,
					Name:     "Automatic",
					Function: c.SetUpAutomatically,
				},
				{
					Key:  pkg.COMMAND_SETUP_MANUAL,
					Name: "Manual",
					Help: `The manual command lets you set up step-by-step, giving you full control.`,
					//Function: c.handleMinerReadConfig,
					Children: []*pkg.Command{
						{
							Key:      pkg.COMMAND_LOCAL_START_CONFIG,
							Help:     "1. Create `./env/local_contracts.json`",
							Function: c.handleLocalConfig,
							Name:     "Config",
							Children: localContractCMDs,
						},
						{
							Key:      pkg.COMMAND_LOCAL_START_HARDHAT,
							Help:     "2. Start HardHat",
							Name:     "Hardhat",
							Function: c.handleStartHardHat,
						},
						{
							Key:      pkg.COMMAND_LOCAL_START_OLLAMA,
							Help:     "3. Start Ollama",
							Name:     "Ollama",
							Function: c.handleStartOllama,
							Children: localContractCMDs,
						},
						{
							Key:      pkg.COMMAND_LOCAL_DEPLOY_CONTRACT,
							Help:     "4. Deploy contracts",
							Name:     "Contracts",
							Function: c.handleStartDeployContracts,
							Children: localContractCMDs,
						},
						{
							Key:      pkg.COMMAND_LOCAL_START_MINERS,
							Help:     "5. Start miners",
							Name:     "Miners",
							Function: c.handleStartMiners,
							Children: localContractCMDs,
						},
					},
				},
			},
		},
		{
			Key:  pkg.COMMAND_INFER,
			Help: "2. Create the Testing Infer (Chat with AI)",
			Children: []*pkg.Command{
				{
					Key:  pkg.COMMAND_INFER_PROMPT,
					Help: fmt.Sprintf("Prompt: To exit at any time, simply press (%s)", pkg.COMMAND_BACK),
					Name: "Infer",
				},
			},
			Function: c.handleCreateInfer,
		},
	}

	c.buildTree(setupCommands, nil)
	return setupCommands
}

// cluster
func (c *CMD) handleClusterCreate(reader *bufio.Reader, node *pkg.Command) {
	/*
		err := c.verify()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		input := c.buildInputData(reader, node)
		c.clusterCMD.SetWatcher(c.taskWatcher)
		tx, clusterID, err := c.clusterCMD.CreateCluster(input)
		if err != nil {
			fmt.Print(pkg.PrintText("create cluster error:", err))
			return
		}
		fmt.Print(pkg.PrintText("Create cluster tx:", tx.Hash().Hex()))
		fmt.Print(pkg.PrintText("Create cluster ID:", clusterID.String()))
	*/
}

func (c *CMD) handleClusterGroupCreate(reader *bufio.Reader, node *pkg.Command) {
	/*
		err := c.verify()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		input := c.buildInputData(reader, node)
		c.clusterCMD.SetWatcher(c.taskWatcher)
		tx, err := c.clusterCMD.CreateAGroupOfCluster(input)
		if err != nil {
			fmt.Print(pkg.PrintText("create cluster error:", err))
			return
		}
		fmt.Print(pkg.PrintText("Create group cluster tx:", tx.Hash().Hex()))
	*/
}

func (c *CMD) handleJoinCluster(reader *bufio.Reader, node *pkg.Command) {
	/*
		err := c.verify()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = c.minerCMD.SetWatcher(c.taskWatcher)
		if err != nil {
			fmt.Print(pkg.PrintText("miner error:", err))
			return
		}

		stakeTx, JoinForMintingTx, err := c.clusterCMD.JoinCluster()
		if err != nil {
			fmt.Println("cluster joined with error: ", err)
			return
		}

		fmt.Print(pkg.PrintText("Stake tx:", stakeTx.Hash().Hex()))
		fmt.Print(pkg.PrintText("Join for minting tx:", JoinForMintingTx.Hash().Hex()))
	*/
}

// local chain
func (c *CMD) handleLocalConfig(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	err := c._startCreateConfigLogic(input)
	if err != nil {
		fmt.Println(pkg.PrintText("Config err", err))
		return
	}
}

func (c *CMD) handleDeployContracts(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	//env := ``

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		fmt.Println("deployed contracts error: private key is required")
		return
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		fmt.Println("deployed contracts error: rpc is required")
		return
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		fmt.Println("deployed contracts error: chainID is required")
		return
	}

	gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]
	if !ok {
		gasPriceBigInt, ok1 := big.NewInt(1).SetString(gasPrice, 10)
		if ok1 {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	}

	gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]
	if !ok {
		gasLimitUint64, err1 := strconv.Atoi(gasLimit)
		if err1 == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	}

	fmt.Println("deploying contracts")
	resp, err := c.localChainCMD.DeployContracts(rpc, chainID, privKey)
	if err != nil {
		fmt.Println("deployed contracts error: ", err)
		return
	}

	_ = resp
	fmt.Println("deployed contracts success")
}

func (c *CMD) handleDeployContract(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	//env := ``

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		fmt.Println("deployed contracts error: private key is required")
		return
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		fmt.Println("deployed contracts error: rpc is required")
		return
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		fmt.Println("deployed contracts error: chainID is required")
		return
	}

	gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]
	if !ok {
		gasPriceBigInt, ok1 := big.NewInt(1).SetString(gasPrice, 10)
		if ok1 {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	}

	gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]
	if !ok {
		gasLimitUint64, err1 := strconv.Atoi(gasLimit)
		if err1 == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	}

	contractName, ok := input[pkg.COMMAND_LOCAL_CONTRACT_NAME]
	if !ok {
		fmt.Println("contract name is required")
		return
	}

	fmt.Println("deploying contract")
	resp, err := c.localChainCMD.DeployContract(rpc, chainID, privKey, contractName)
	if err != nil {
		fmt.Println("deployed contract error: ", err)
		return
	}

	_ = resp
	fmt.Println("deployed contract success")

}

func (c *CMD) handleMintWEAI(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	//env := ``

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		fmt.Println("deployed contracts error: private key is required")
		return
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		fmt.Println("deployed contracts error: rpc is required")
		return
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		fmt.Println("deployed contracts error: chainID is required")
		return
	}

	gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]
	if !ok {
		gasPriceBigInt, ok1 := big.NewInt(1).SetString(gasPrice, 10)
		if ok1 {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	}

	gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]
	if !ok {
		gasLimitUint64, err1 := strconv.Atoi(gasLimit)
		if err1 == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	}

	fmt.Println("Minting WEAI")
	resp, err := c.localChainCMD.MintWrappedEAI(rpc, chainID, "100000", privKey)
	if err != nil {
		fmt.Println("deployed contracts error: ", err)
		return
	}

	_ = resp
	fmt.Println("Minting WEAI success")
}

func (c *CMD) handleSetWEAIForGpuManager(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	//env := ``

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		fmt.Println("deployed contracts error: private key is required")
		return
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		fmt.Println("deployed contracts error: rpc is required")
		return
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		fmt.Println("deployed contracts error: chainID is required")
		return
	}
	_ = chainID

	gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]
	if !ok {
		gasPriceBigInt, ok1 := big.NewInt(1).SetString(gasPrice, 10)
		if ok1 {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	}

	gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]
	if !ok {
		gasLimitUint64, err1 := strconv.Atoi(gasLimit)
		if err1 == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	}

	client, err := eth.NewEthClient(rpc)
	if err != nil {
		fmt.Println(pkg.PrintText("SetWEAIForStakingHub error", err.Error()))
		return
	}

	fmt.Println(pkg.PrintText("SetWEAIForStakingHub", ""))
	resp, err := c.localChainCMD.SetWEAIForStakingHub(client, privKey)
	if err != nil {
		fmt.Println(pkg.PrintText("SetWEAIForStakingHub error", err.Error()))
		return
	}
	fmt.Println(pkg.PrintText("SetWEAIForStakingHub with tx", resp.Hash().Hex()))
}

func (c *CMD) handleCreateMinerInfo(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	//env := ``

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		fmt.Println("deployed contracts error: private key is required")
		return
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		fmt.Println("deployed contracts error: rpc is required")
		return
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		fmt.Println("deployed contracts error: chainID is required")
		return
	}

	minerAddress, minerPrvKey, err := c.localChainCMD.CreateMinerAddress(rpc, chainID, privKey)
	if err != nil {
		fmt.Println("Create miner address error: ", err)
		return
	}

	fmt.Print(pkg.PrintText("Miner address", minerAddress))
	fmt.Print(pkg.PrintText("Miner private key", minerPrvKey))
}

// setup automatically
func (c *CMD) SetUpAutomatically(reader *bufio.Reader, node *pkg.Command) {
	var err error
	fmt.Println("Setup cluster")
	input := c.buildInputData(reader, node)
	//env := ``
	err = c._startCreateConfigLogic(input)
	if err != nil {
		fmt.Println(pkg.PrintText("Config err", err))
		return
	}

	err = c.localChainCMD.StartHardHat()
	if err != nil {
		fmt.Println(pkg.PrintText("Hardhat start with err", err))
		return
	}

	time.Sleep(2 * time.Second)
	//deploy all needed contracts
	//c.localChainCMD.ContractDeployment()

	//1. Deploy all contracts
	err = c._deployContractLogic()
	if err != nil {
		fmt.Println("_deployContractLogic error: ", err)
		return
	}

	err = c._startMinerLogic()
	if err != nil {
		fmt.Println("_startMinerLogic error: ", err)
		return
	}

	c.handleStartOllama(reader, node)

	fmt.Print(pkg.Line)
	fmt.Println("Done!!!")
	fmt.Print(pkg.Line)
	//ALL done!!!
}

func (c *CMD) handleCreateInfer(reader *bufio.Reader, node *pkg.Command) {
	contextMsg := []model.LLMInferMessage{}

	for {
		input := c.buildInputData(reader, node)
		_prompt, ok := input[pkg.COMMAND_INFER_PROMPT]
		if !ok {
			return
		}

		if strings.EqualFold(_prompt, pkg.COMMAND_BACK) {
			return
		}

		if _prompt != "" {

			contextMsg = append(contextMsg, model.LLMInferMessage{
				Role:    "user",
				Content: _prompt,
			})

			_, inferID, result, err := c.localChainCMD.CreateInfer(contextMsg)
			if err != nil {
				fmt.Println("create infer error", err)
				return
			}

			if result == nil {
				err = errors.New("error while get result")
				fmt.Println("create infer error", err)
				return
			}

			fmt.Print(pkg.Line)
			fmt.Print(pkg.PrintText("InferID#", *inferID))
			fmt.Print(pkg.PrintText("Prompt", _prompt))
			fmt.Print(pkg.PrintText("Result", *result))
			fmt.Print(pkg.Line)

			contextMsg = append(contextMsg, model.LLMInferMessage{
				Role:    "assistant",
				Content: *result,
			})
		}
	}
}

func (c *CMD) createConfigENV(minerAddress string, index int) error {
	sample := fmt.Sprintf(pkg.ENV_SAMPLE_FILE, pkg.CurrentDir())
	envFile := fmt.Sprintf(pkg.ENV_LOCAL_MINERS_FILE, pkg.CurrentDir(), index)

	config.ReadConfig(sample)

	f, _ := os.Stat(envFile)
	if f != nil {
		err := os.Remove(envFile)
		if err != nil {
			return err
		}
	}

	cnf := c.localChainCMD.ReadLocalChainCnf()

	env := ""
	env += fmt.Sprintf("PLATFORM=%v\n", pkg.PLATFROM_INTEL)
	env += fmt.Sprintf("API_URL=%v\n", cnf.RunPodInternal)
	env += fmt.Sprintf("API_KEY=%v\n", cnf.RunPodAPIKEY)
	env += fmt.Sprintf("LIGHT_HOUSE_API_KEY=%v\n", os.Getenv("LIGHT_HOUSE_API_KEY"))
	env += fmt.Sprintf("CLUSTER_ID=%v\n", cnf.ModelID)
	env += fmt.Sprintf("MODEL_ID=%v\n", cnf.ModelID)
	env += fmt.Sprintf("CHAIN_ID=%v\n", cnf.ChainID)
	env += fmt.Sprintf("CHAIN_RPC=%v\n", fmt.Sprintf(`http://%s:8545`, pkg.MINER_SERVICE_HARDHAT))
	env += fmt.Sprintf("ACCOUNT_PRIV=%v\n", cnf.Miners[strings.ToLower(minerAddress)].PrivateKey)
	env += fmt.Sprintf("MODEL_NAME=%v\n", cnf.ModelName)
	env += fmt.Sprintf("STAKING_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_GPU_MANAGER])
	env += fmt.Sprintf("MODEL_LOAD_BALANCER_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_load_balancer])
	env += fmt.Sprintf("WORKER_HUB_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_PROMPT_SCHEULER])
	env += fmt.Sprintf("ERC20_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_WEAI])
	env += fmt.Sprintf("COLLECTION_ADDRESS=%v\n", cnf.Contracts[pkg.COMMAND_LOCAL_CONTRACTS_DEPLOY_ONE_C_MODEL_COLLECTION])

	err := pkg.CreateFile(envFile, []byte(env))
	if err != nil {
		return err
	}
	return nil
}

// setup automatically - manual
// 1.
func (c *CMD) handleStartHardHat(reader *bufio.Reader, node *pkg.Command) {
	err := c.localChainCMD.StartHardHat()
	if err != nil {
		fmt.Println(pkg.PrintText("Hardhat start with err", err))
		return
	}
}

// 2.
func (c *CMD) handleStartOllama(reader *bufio.Reader, node *pkg.Command) {
	err := c.localChainCMD.StartOllama()
	if err != nil {
		fmt.Println(pkg.PrintText("Ollama start with err", err))
		return
	}
}

// 3.
func (c *CMD) handleStartDeployContracts(reader *bufio.Reader, node *pkg.Command) {
	err := c._deployContractLogic()
	if err != nil {
		fmt.Println("Deployed contracts error: ", err)
		return
	}
}

func (c *CMD) handleStartMiners(reader *bufio.Reader, node *pkg.Command) {
	err := c._startMinerLogic()
	if err != nil {
		fmt.Println("Start miner error: ", err)
		return
	}
}

func (c *CMD) _deployContractLogic() error {
	cnf := c.localChainCMD.ReadLocalChainCnf()
	rpc := cnf.Rpc
	chainID := cnf.ChainID
	privKey := cnf.PrivateKey
	modelName := cnf.ModelName

	_, err := c.localChainCMD.DeployContracts(rpc, chainID, privKey)
	if err != nil {
		fmt.Println("Deployed contracts error: ", err)
		return err
	}

	//3. Mint WEAI.
	_, err = c.localChainCMD.MintWrappedEAI(rpc, chainID, "100000", privKey)
	if err != nil {
		fmt.Println("Mint WEAI error: ", err)
		return err
	}

	_, tokenID, err := c.localChainCMD.MintCollection(rpc, privKey, modelName)
	if err != nil {
		fmt.Println("Mint collection error: ", err)
		return err
	}

	txRegister, err := c.localChainCMD.SetGPUAddressRegisterModel(rpc, uint32(tokenID.Int64()), privKey)
	if err != nil {
		fmt.Println("Mint collection error: ", err)
		return err
	}

	fmt.Print(pkg.PrintText("register model tx", txRegister.Hash().Hex()))
	return nil
}

func (c *CMD) _startMinerLogic() error {
	fmt.Print(pkg.Line)
	fmt.Println("Start miners")

	cnf := c.localChainCMD.ReadLocalChainCnf()
	rpc := cnf.Rpc
	chainID := cnf.ChainID
	privKey := cnf.PrivateKey

	fmt.Println("Create miners: ")
	numberOfMiners := 3
	names := ""

	//clear the created miners
	cnf.Miners = make(map[string]model.Miners)
	_b, err := json.Marshal(cnf)
	if err == nil {
		pkg.CreateFile(fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir()), _b)
	}

	for i := 1; i <= numberOfMiners; i++ {
		fmt.Print(pkg.Line)
		//5. Create a miner's private key (3 miner)
		minerAddress, minerPrvKey, err := c.localChainCMD.CreateMinerAddress(rpc, chainID, privKey)
		if err != nil {
			continue
		}

		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d address", i), *minerAddress))
		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d private key", i), *minerPrvKey))

		//send WEAI
		tx, _, err := c.localChainCMD.SendWEIToMiner(rpc, *minerAddress)
		if err != nil {
			fmt.Println("SendWEIToMiner error", err)
			continue
		}
		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d received WEAI TX", i), tx.Hash().Hex()))

		//send fee
		txFee := new(types.Transaction)
		gas := pkg.LOCAL_CHAIN_GAS_LIMIT

		_loop := 1
		for {
			if _loop >= 50 {
				break
			}

			time.Sleep(time.Second * 2)
			txFee, _, err = c.localChainCMD.SendFeeToMiner(rpc, *minerAddress, uint64(gas))
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
			//there is no error
			fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d gas limit", i), gas))

			_loop++

			break

		}

		err = c.createConfigENV(*minerAddress, i)
		if err != nil {
			fmt.Println(fmt.Sprintf("Create config for miner %d error", i), err)
			gas += gas
			continue
		}

		fmt.Print(pkg.PrintText(fmt.Sprintf("Miner %d received Fee TX", i), txFee.Hash().Hex()))

		name := fmt.Sprintf("%s_%d", pkg.MINER_SERVICE_NAME, i)
		names += " " + name
	}

	errBuild := c.localChainCMD.BuildContainers(names)
	if errBuild == nil {
		c.localChainCMD.StartContainers(names)
	}

	return nil
}

func (c *CMD) _startCreateConfigLogic(input map[string]string) error {
	cnf := c.localChainCMD.ReadLocalChainCnf()

	var err error
	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		err = errors.New("deployed contracts error: private key is required")
		return err
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		//err = errors.New("deployed contracts error: rpc is required")
		//return err
		rpc = "http://localhost:8545"
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		//err = errors.New("deployed contracts error: chainID is required")
		//return err
		chainID = "31337"
	}

	gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]
	if ok {
		gasPriceBigInt, ok1 := big.NewInt(1).SetString(gasPrice, 10)
		if ok1 {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	} else {
		c.localChainCMD.SetGasPrice(big.NewInt(int64(pkg.LOCAL_CHAIN_GAS_PRICE)))
	}

	gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]
	if ok {
		gasLimitUint64, err1 := strconv.Atoi(gasLimit)
		if err1 == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	} else {
		c.localChainCMD.SetGasLimit(uint64(pkg.LOCAL_CHAIN_GAS_LIMIT))
	}

	modelName, ok := input[pkg.COMMAND_LOCAL_MODEL_NAME]
	if !ok {
		err = errors.New("deployed contracts error: modelName is required")
		return err
	}

	runPod, ok := input[pkg.COMMAND_LOCAL_RUN_POD_URL]
	if !ok {
		return err
	}

	runPodAPIKey, ok := input[pkg.COMMAND_LOCAL_RUN_POD_API_KEY]
	if !ok {
		return err
	}

	cnf.PrivateKey = privKey
	cnf.Rpc = rpc
	cnf.ChainID = chainID
	cnf.ModelName = modelName
	cnf.RunPodAPIKEY = runPodAPIKey

	if runPod != "" {
		cnf.RunPodInternal = runPod
		cnf.RunPodExternal = runPod
		cnf.UseExternalRunPod = true
	} else {
		cnf.RunPodInternal = fmt.Sprintf("http://%s:11434/v1/chat/completions", pkg.MINER_SERVICE_OLLAMA)
		cnf.RunPodExternal = fmt.Sprintf("http://%s:11436/v1/chat/completions", "localhost")
		cnf.UseExternalRunPod = false
	}

	_b, err := json.Marshal(cnf)
	if err != nil {
		return err
	}

	envFile := fmt.Sprintf(pkg.LOCAL_CHAIN_INFO, pkg.CurrentDir())
	err = pkg.CreateFile(envFile, _b)
	if err != nil {
		return err
	}

	return nil
}

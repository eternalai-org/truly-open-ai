package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"strconv"
	"strings"
	"time"

	"solo/internal/model"
	"solo/pkg"
	"solo/pkg/eth"
)

func (c *CMD) cliCommand() []*pkg.Command {
	localContractCMDs := []*pkg.Command{
		// {
		// 	Key:      pkg.COMMAND_LOCAL_PRIV_KEY,
		// 	Help:     "Private Key",
		// 	Default:  "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		// 	Required: true,
		// },
		{
			Key:      pkg.PLATFORM,
			Required: true,
			Default:  c.getArch(),
			Help:     "Platform: " + pkg.PLATFROM_INTEL + " OR " + pkg.PLATFROM_APPLE_SILLICON,
		},
		// {
		// 	Key:     pkg.COMMAND_LOCAL_RUN_POD_URL,
		// 	Help:    "Runpod URL (Default: empty)",
		// 	Default: "",
		// },
		// {
		// 	Key:     pkg.COMMAND_LOCAL_RUN_POD_API_KEY,
		// 	Help:    "Runpod API-Key (Default: empty)",
		// 	Default: "",
		// },
		// {
		// 	Key:      pkg.COMMAND_LOCAL_MODEL_NAME,
		// 	Help:     "Model name",
		// 	Default:  "hf.co/bartowski/Meta-Llama-3.1-8B-Instruct-GGUF:Q3_K_S",
		// 	Required: true,
		// },
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
					// Function: c.handleMinerReadConfig,
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
						// {
						// 	Key:      pkg.COMMAND_LOCAL_START_OLLAMA,
						// 	Help:     "3. Start Ollama",
						// 	Name:     "Ollama",
						// 	Function: c.handleStartOllama,
						// 	Children: localContractCMDs,
						// },
						{
							Key:      pkg.COMMAND_LOCAL_DEPLOY_CONTRACT,
							Help:     "3. Deploy contracts",
							Name:     "Contracts",
							Function: c.handleStartDeployContracts,
							Children: localContractCMDs,
						},
						{
							Key:      pkg.COMMAND_LOCAL_START_MINERS,
							Help:     "4. Start miners",
							Name:     "Miners",
							Function: c.handleStartMiners,
							Children: localContractCMDs,
						},
						{
							Key:      pkg.COMMAND_LOCAL_START_APIS,
							Help:     "5. Start APIs",
							Name:     "APIs",
							Function: c.handleStartApi,
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
	}
}

func (c *CMD) handleDeployContracts(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	// env := ``

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
	_, err := c.localChainCMD.DeployContracts(rpc, chainID, privKey)
	if err != nil {
		fmt.Println("deployed contracts error: ", err)
		return
	}

	fmt.Println("deployed contracts success")
}

func (c *CMD) handleDeployContract(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	// env := ``

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
	_, err := c.localChainCMD.DeployContract(rpc, chainID, privKey, contractName)
	if err != nil {
		fmt.Println("deployed contract error: ", err)
		return
	}
	fmt.Println("deployed contract success")
}

func (c *CMD) handleMintWEAI(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	// env := ``

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
	_, err := c.localChainCMD.MintWrappedEAI(rpc, chainID, "100000", privKey)
	if err != nil {
		fmt.Println("deployed contracts error: ", err)
		return
	}
	fmt.Println("Minting WEAI success")
}

func (c *CMD) handleSetWEAIForGpuManager(reader *bufio.Reader, node *pkg.Command) {
	input := c.buildInputData(reader, node)
	// env := ``

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
	// env := ``

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

func (c *CMD) _setupAutomaticallyLogic(input map[string]string) error {
	var err error
	err = c._startCreateConfigLogic(input)
	if err != nil {
		fmt.Println(pkg.PrintText("Config err", err))
		return err
	}

	err = c.localChainCMD.StartHardHat()
	if err != nil {
		fmt.Println(pkg.PrintText("Hardhat start with err", err))
		return err
	}

	time.Sleep(2 * time.Second)
	// deploy all needed contracts
	// c.localChainCMD.ContractDeployment()

	// 1. Deploy all contracts
	err = c._deployContractLogic()
	if err != nil {
		fmt.Println("_deployContractLogic error: ", err)
		return err
	}

	err = c._startMinerLogic()
	if err != nil {
		fmt.Println("_startMinerLogic error: ", err)
		return err
	}

	err = c._startAPILogic()
	if err != nil {
		fmt.Println("_startMinerLogic error: ", err)
		return err
	}

	return nil
}

// setup automatically
func (c *CMD) SetUpAutomatically(reader *bufio.Reader, node *pkg.Command) {
	var err error
	fmt.Println("Setup cluster")
	input := c.buildInputData(reader, node)

	// env := ``
	err = c._setupAutomaticallyLogic(input)
	if err != nil {
		return
	}

	c.handleStartOllama(reader, node)
	fmt.Print(pkg.Line)
	fmt.Println("Done!!!")
	fmt.Print(pkg.Line)
	// ALL done!!!
}

// setup automatically
func (c *CMD) SetUpAutomaticallyOneStep() {
	var err error
	fmt.Println("Setup cluster one step")
	// input := c.buildInputData(reader, node)
	input := map[string]string{}

	err = c._setupAutomaticallyLogic(input)
	if err != nil {
		return
	}

	c.handleStartOllama(nil, nil)
	fmt.Print(pkg.Line)
	fmt.Println("Done!!!")
	fmt.Print(pkg.Line)
	// ALL done!!!
}

func (c *CMD) handleCreateInfer(reader *bufio.Reader, node *pkg.Command) {
	contextMsg := []model.LLMInferMessage{}
	useStream := true

	for {
		var err error
		var inferID *uint64

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

			fmt.Print(pkg.Line)
			fmt.Print(pkg.PrintText("Prompt", _prompt))
			result := new(string)
			if useStream {
				outChan := make(chan model.StreamDataChannel)
				errFChan := make(chan error)

				go func(errFChan chan error, outChan chan model.StreamDataChannel) {
					_, _, _, err = c.localChainCMD.CreateInferWithStream(contextMsg, outChan)
				}(errFChan, outChan)

				if err != nil {
					fmt.Println("create infer error", err)
					return
				}

				for v := range outChan {
					if v.Err != nil {
						err = v.Err
						fmt.Println("create infer error", err)
						return
					}

					if v.Data != nil {
						if len(v.Data.Choices) > 0 {
							msg := v.Data.Choices[0].Message.Content
							fmt.Print(msg)
							*result += msg
						}

						if inferID == nil {
							inferID = &v.InferID
						}
					}

				}

				//fmt.Print(pkg.PrintText("InferID#", *inferID))

			} else {
				_, inferID, result, err = c.localChainCMD.CreateInfer(contextMsg)
				if err != nil {
					fmt.Println("create infer error", err)
					return
				}

				if result == nil {
					err = errors.New("error while get result")
					fmt.Println("create infer error", err)
					return
				}

				fmt.Print(pkg.PrintText("InferID#", *inferID))
				fmt.Print(pkg.PrintText("Result", *result))
			}

			fmt.Println("")
			fmt.Print(pkg.Line)
			contextMsg = append(contextMsg, model.LLMInferMessage{
				Role:    "assistant",
				Content: *result,
			})
		}
	}
}

func (c *CMD) createConfigENV(minerAddress string, index int) error {
	return c.localChainCMD.CreateConfigENV(minerAddress, index)
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
func (c *CMD) handleStartOllama(_ *bufio.Reader, _ *pkg.Command) {
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

func (c *CMD) handleStartApi(reader *bufio.Reader, node *pkg.Command) {
	err := c._startAPILogic()
	if err != nil {
		fmt.Println("Start APIs error: ", err)
		return
	}
}

func (c *CMD) _deployContractLogic() error {
	return c.localChainCMD.DeployContractLogic()
}

func (c *CMD) _startMinerLogic() error {
	return c.localChainCMD.StartMinerLogic()
}

func (c *CMD) _startAPILogic() error {
	return c.localChainCMD.StartApiLogic()
}

func (c *CMD) _startCreateConfigLogic(input map[string]string) error {
	cnf := c.localChainCMD.ReadLocalChainCnf()
	input[pkg.PLATFORM] = c.getArch()

	privKey, ok := input[pkg.COMMAND_LOCAL_PRIV_KEY]
	if !ok {
		privKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	}

	rpc, ok := input[pkg.COMMAND_LOCAL_CHAIN_RPC]
	if !ok {
		rpc = "http://localhost:8545"
	}

	pubsub, ok := input[pkg.COMMAND_LOCAL_PUBSUB]
	if !ok {
		pubsub = "localhost:6380"
	}

	chainID, ok := input[pkg.COMMAND_LOCAL_CHAIN_ID]
	if !ok {
		chainID = "31337"
	}

	c.localChainCMD.SetGasPrice(big.NewInt(int64(pkg.LOCAL_CHAIN_GAS_PRICE)))
	if gasPrice, ok := input[pkg.COMMAND_LOCAL_GAS_PRICE]; ok {
		if gasPriceBigInt, ok := big.NewInt(1).SetString(gasPrice, 10); ok {
			c.localChainCMD.SetGasPrice(gasPriceBigInt)
		}
	}

	c.localChainCMD.SetGasLimit(uint64(pkg.LOCAL_CHAIN_GAS_LIMIT))
	if gasLimit, ok := input[pkg.COMMAND_LOCAL_GAS_LIMIT]; ok {
		if gasLimitUint64, err := strconv.Atoi(gasLimit); err == nil {
			c.localChainCMD.SetGasLimit(uint64(gasLimitUint64))
		}
	}

	modelName, ok := input[pkg.COMMAND_LOCAL_MODEL_NAME]
	if !ok {
		modelName = "deepseek-r1:1.5b-qwen-distill-q8_0"

		//TODO - wait util fix bug of DeepSeek-R1-Distill-Qwen-1.5B-Q8
		//modelName = "DeepSeek-R1-Distill-Qwen-1.5B-Q8"
	}

	runPod, ok := input[pkg.COMMAND_LOCAL_RUN_POD_URL]
	if !ok {
		runPod = "1"
	}

	runPodAPIKey, ok := input[pkg.COMMAND_LOCAL_RUN_POD_API_KEY]
	if !ok {
		runPodAPIKey = "1"
	}

	platform, ok := input[pkg.PLATFORM]
	if !ok {
		return errors.New("deployed contracts error: platform is required")
	}

	cnf.Platform = platform
	cnf.PrivateKey = privKey
	cnf.Rpc = rpc
	cnf.ChainID = chainID
	cnf.ModelName = modelName
	cnf.RunPodAPIKEY = runPodAPIKey
	cnf.PubSubURL = pubsub

	if runPod != "" && runPod != "1" {
		cnf.RunPodInternal = runPod
		cnf.RunPodExternal = runPod
		cnf.UseExternalRunPod = true
	} else if runPod == "1" {
		// cnf.RunPodInternal = "http://host.docker.internal:11434/v1/chat/completions"
		// we are using docker network_mode=host
		cnf.RunPodInternal = "http://localhost:11434/v1/chat/completions"
		cnf.RunPodExternal = "http://localhost:11434/v1/chat/completions"
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
	return pkg.CreateFile(envFile, _b)
}

func (c *CMD) getArch() string {
	switch runtime.GOARCH {
	case "amd64", "386":
		return pkg.PLATFROM_INTEL
	case "arm", "arm64":
		return pkg.PLATFROM_APPLE_SILLICON
	}
	return ""
}

package chat

import (
	"bufio"
	"bytes"
	"context"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/libs/http_client"
	"decentralized-inference/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func AgentTerminalChat(ctx context.Context) error {
	chatConfig, err := LoadChatConfig()
	if err != nil {
		//fmt.Println("Error loading chat config:", err)
		return err
	}

	fmt.Println("Welcome to the EAI chat terminal!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)
		if strings.ToLower(userInput) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		response, err := getLLMResponse(userInput, chatConfig)
		if err != nil {
			fmt.Println("Error getting response from server:", err)
			continue
		}

		respBytes, _ := json.MarshalIndent(response, "", "  ")
		fmt.Printf("Infer response: %s\n\n", string(respBytes))
		// show loading by  / - \
		stopChan := make(chan bool)
		go showLoading(stopChan)
		timeout := time.After(3 * time.Minute)
		for {
			select {
			case <-timeout:
				stopChan <- true
				fmt.Println("Timeout, please try again later.")
				break
			default:
				getResult(response.InferId, chatConfig, stopChan)
				time.Sleep(5 * time.Second)
			}
		}

	}

	return nil
}

func getResult(inferID uint64, chatConfig *config.ChatConfig, stopChan chan bool) {
	fmt.Println("Getting result from server...")
	uri := fmt.Sprintf("infer/%v", inferID)
	fullUrl := fmt.Sprintf("%v/%v", chatConfig.ServerBaseUrl, uri)

	request := &models.InferResultRequest{
		ChainInfo: models.ChainInfoRequest{
			Rpc: chatConfig.ChainRpc,
		},
		WorkerHubAddress: chatConfig.PromptSchedulerContractAddress,
		InferId:          inferID,
	}
	inputBytes, _ := json.Marshal(request)
	respBytes, statusCode, err := http_client.RequestHttp(fullUrl, "POST", nil, bytes.NewBuffer(inputBytes), 5)
	if err != nil {
		fmt.Println("Error getting result from server:", err)
		return
	}
	if statusCode != http.StatusOK {
		//fmt.Println("Error getting result from server: status code", statusCode)
		return
	}

	var response struct {
		Data models.InferResultResponse `json:"data"`
	}

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	// Display response
	if response.Data.Output == "" {
		//fmt.Println("No result yet. Please try again later.")
		return
	}

	stopChan <- true

	//respBytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Printf("Your agent: %s\n\n", response.Data.Output)
}

func showLoading(stopChan chan bool) {
	// Define the loading symbols
	symbols := []string{"/", "-", "\\", "|"}
	index := 0

	for {
		select {
		case <-stopChan:
			// Stop the loading animation
			return
		default:
			// Display the current loading symbol
			fmt.Printf("\rLoading %s", symbols[index])
			index = (index + 1) % len(symbols) // Cycle through the symbols
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func getLLMResponse(prompt string, chatConfig *config.ChatConfig) (*models.DecentralizeInferResponse, error) {
	request := models.DecentralizeInferRequest{
		ChainInfo: models.ChainInfoRequest{
			Rpc: chatConfig.ChainRpc,
		},
		AgentContractAddress: chatConfig.Dagent721ContractAddress,
		WorkerHubAddress:     chatConfig.PromptSchedulerContractAddress,
		InferPriKey:          chatConfig.InferWalletKey,
		Input:                prompt,
		AgentId:              chatConfig.AgentID,
	}

	uri := "infer/create"
	fullUrl := fmt.Sprintf("%v/%v", chatConfig.ServerBaseUrl, uri)

	input, _ := json.Marshal(request)
	respBytes, statusCode, err := http_client.RequestHttp(fullUrl, "POST", nil, bytes.NewBuffer(input), 5)
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %v != 200", statusCode)
	}

	var response struct {
		Data models.DecentralizeInferResponse `json:"data"`
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

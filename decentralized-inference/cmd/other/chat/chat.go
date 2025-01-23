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

	"github.com/sashabaranov/go-openai"
)

func AgentTerminalChat(ctx context.Context, agentID string) error {
	chatConfig, err := LoadChatConfig()
	if err != nil {
		return err
	}
	chatConfig.AgentID = agentID

	fmt.Println("Welcome to the EAI chat terminal!")
	err = chatConfig.VerifyBeforeChat()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Your angel ID is %v was minted at contract address: %v", agentID, chatConfig.AgentContractAddress))
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

		stopChan := make(chan bool)
		go showLoading(stopChan)
		response, err := getLLMResponseV2(userInput, chatConfig, stopChan)
		if err != nil {
			fmt.Println("Error getting response from server:", err)
			continue
		}

		fmt.Printf("\nYour agent: %v\n\n", response.Choices[0].Message.Content)
	}

	return nil
}

func getResult(inferID uint64, chatConfig *config.ChatConfig, stopChan chan bool) bool {
	fmt.Println("Getting result from server...")
	fullUrl := fmt.Sprintf("%v/infer/get_result", chatConfig.ServerBaseUrl)

	request := &models.InferResultRequest{
		ChainInfo: models.ChainInfoRequest{
			Rpc: chatConfig.Rpc,
		},
		WorkerHubAddress: chatConfig.Contracts.WorkerHubAddress,
		InferId:          inferID,
	}
	inputBytes, _ := json.Marshal(request)
	respBytes, statusCode, err := http_client.RequestHttp(fullUrl, "POST", nil, bytes.NewBuffer(inputBytes), 5)
	if err != nil {
		fmt.Println("Error getting result from server:", err)
		return false
	}
	if statusCode != http.StatusOK {
		//fmt.Println("Error getting result from server: status code", statusCode)
		return false
	}

	var response struct {
		Data models.InferResultResponse `json:"data"`
	}

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return false
	}

	// Display response
	if response.Data.Output == "" {
		//fmt.Println("No result yet. Please try again later.")
		return false
	}

	stopChan <- true

	//respBytes, _ = json.MarshalIndent(response, "", "  ")
	fmt.Printf("Your agent: %s\n\n", response.Data.Output)
	return true
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
			fmt.Printf("\r%s", symbols[index])
			index = (index + 1) % len(symbols) // Cycle through the symbols
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func getLLMResponseV2(prompt string, chatConfig *config.ChatConfig, stopChain chan bool) (*openai.ChatCompletionResponse, error) {
	defer func() {
		stopChain <- true
	}()
	request := models.DecentralizeInferRequest{
		ChainInfo: models.ChainInfoRequest{
			Rpc: chatConfig.Rpc,
		},
		AgentContractAddress: chatConfig.Contracts.SystemPromptManagerAddress,
		WorkerHubAddress:     chatConfig.Contracts.WorkerHubAddress,
		InferPriKey:          chatConfig.PrivateKey,
		Input:                prompt,
		AgentId:              chatConfig.AgentID,
		Model:                chatConfig.ModelName,
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
		Data openai.ChatCompletionResponse `json:"data"`
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func getLLMResponse(prompt string, chatConfig *config.ChatConfig) (*models.DecentralizeInferResponse, error) {
	request := models.DecentralizeInferRequest{
		ChainInfo: models.ChainInfoRequest{
			Rpc: chatConfig.Rpc,
		},
		AgentContractAddress: chatConfig.Contracts.SystemPromptManagerAddress,
		WorkerHubAddress:     chatConfig.Contracts.WorkerHubAddress,
		InferPriKey:          chatConfig.PrivateKey,
		Input:                prompt,
		AgentId:              chatConfig.AgentID,
		Model:                chatConfig.ModelName,
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

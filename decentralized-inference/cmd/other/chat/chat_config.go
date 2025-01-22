package chat

import (
	"bytes"
	"context"
	"decentralized-inference/internal/config"
	"decentralized-inference/internal/libs/http_client"
	"decentralized-inference/internal/rest"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func collectChatConfigInformation() *config.ChatConfig {
	var chatConfig config.ChatConfig

	for {
		fmt.Print("What is your network RPC: ")
		fmt.Scanln(&chatConfig.ChainRpc)
		if chatConfig.ChainRpc == "" {
			fmt.Println("Network RPC cannot be empty")
			continue
		}
		break
	}

	for {
		fmt.Print("What is your decentralize-inference API server's base URL (default: http://localhost:8484)? ")
		fmt.Scanln(&chatConfig.ServerBaseUrl)
		if chatConfig.ServerBaseUrl == "" {
			chatConfig.ServerBaseUrl = "http://localhost:8484"
		}
		break
	}

	for {
		fmt.Print("What is your DAgent721 contract address: ")
		fmt.Scanln(&chatConfig.Dagent721ContractAddress)
		if chatConfig.Dagent721ContractAddress == "" {
			fmt.Println("DAgent721 contract address cannot be empty")
			continue
		}
		break
	}

	for {
		fmt.Print("What is your agent ID: ")
		fmt.Scanln(&chatConfig.AgentID)
		if chatConfig.AgentID == "" {
			fmt.Println("Agent ID cannot be empty")
			continue
		}
		break
	}

	for {
		fmt.Print("What is your prompt scheduler contract address: ")
		fmt.Scanln(&chatConfig.PromptSchedulerContractAddress)
		if chatConfig.PromptSchedulerContractAddress == "" {
			fmt.Println("Prompt scheduler contract address cannot be empty")
			continue
		}
		break
	}

	for {
		fmt.Print("What is your infer wallet key: ")
		fmt.Scanln(&chatConfig.InferWalletKey)
		if chatConfig.InferWalletKey == "" {
			fmt.Println("Infer wallet key cannot be empty")
			continue
		}
		break
	}

	return &chatConfig
}

func AgentTerminalChatConfig(ctx context.Context) error {
	var chatConfig = collectChatConfigInformation()

	// Confirm before saving
	fmt.Println("\nReview chat config information:")
	bytes, _ := json.MarshalIndent(&chatConfig, "", "  ")
	fmt.Println(string(bytes))

	fmt.Print("\nDo you want to save this to config.json? (yes/no): ")
	var saveConfirmation string
	fmt.Scanln(&saveConfirmation)
	fmt.Println("confirm:", saveConfirmation)

	if !(strings.EqualFold(saveConfirmation, "yes") || strings.EqualFold(saveConfirmation, "y")) {
		fmt.Println("Configuration not saved.")
		return nil
	}

	// Send Config to Server
	err := SendChatConfigToServer(chatConfig)
	if err != nil {
		fmt.Println("Error while sending chat config to server: ", err, "\tTry again.")
		return nil
	}
	// Save to JSON file
	file, err := os.Create("chat_config.json")
	if err != nil {
		return fmt.Errorf("failed to create config.json: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&chatConfig); err != nil {
		return fmt.Errorf("failed to write config to file: %v", err)
	}

	fmt.Println("Chat configuration saved to chat_config.json")
	return nil
}

func LoadChatConfig() (*config.ChatConfig, error) {
	file, err := os.Open("chat_config.json")
	if err != nil {
		return nil, fmt.Errorf("chat_config.json not found, please run 'chat config-all' to create it")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var chatConfig config.ChatConfig
	if err := decoder.Decode(&chatConfig); err != nil {
		return nil, fmt.Errorf("chat_config.json is invalid, please check the values")
	}

	return &chatConfig, nil
}

func SendChatConfigToServer(chatConfig *config.ChatConfig) error {
	fullUrl := fmt.Sprintf("%v/chain_config/insert", chatConfig.ServerBaseUrl)
	inputBytes, _ := json.Marshal(chatConfig)
	respBytes, statusCode, err := http_client.RequestHttp(fullUrl, "POST", nil, bytes.NewBuffer(inputBytes), 5)
	if err != nil {
		return fmt.Errorf("sending chat_config to server failed, please check the values")
	}
	if statusCode != http.StatusOK {
		return fmt.Errorf("sending chat_config to server failed with status code %d , response :%v", statusCode, string(respBytes))
	}
	var response rest.Response
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return err
	}
	if response.Status != rest.StatusSuccess {
		return fmt.Errorf("sending chat_config to server failed with response :%v", string(respBytes))
	}
	return nil
}

package chat

import (
	"context"
	"decentralized-inference/internal/config"
	"encoding/json"
	"fmt"
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
		fmt.Print("What is your decentralize-inference API server's base URL (default: localhost:8484)? ")
		fmt.Scanln(&chatConfig.ServerBaseUrl)
		if chatConfig.ServerBaseUrl == "" {
			chatConfig.ServerBaseUrl = "localhost:8484"
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

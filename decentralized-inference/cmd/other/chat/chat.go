package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func AgentTerminalChat(ctx context.Context) error {
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

		response, err := getLLMResponse(nil, userInput)
		if err != nil {
			fmt.Println("Error getting response from LLM:", err)
			continue
		}

		// Display response
		fmt.Printf("Response: %s\n\n", response)

	}

	return nil
}

func getLLMResponse(client *openai.Client, prompt string) (string, error) {
	return "Response from LLM", nil
}

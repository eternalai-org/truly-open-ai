package openai

import (
	openai2 "github.com/sashabaranov/go-openai"
)

func GetSystemPromptFromLLMMessage(messages []openai2.ChatCompletionMessage) string {
	systemPrompt := ""
	for _, message := range messages {
		if message.Role == openai2.ChatMessageRoleSystem {
			systemPrompt = message.Content
			break
		}
	}

	return systemPrompt
}

func UpdateSystemPromptInLLMRequest(message []openai2.ChatCompletionMessage, systemPrompt string) []openai2.ChatCompletionMessage {
	for i, m := range message {
		if m.Role == openai2.ChatMessageRoleSystem {
			message[i].Content = systemPrompt
			break
		}
	}

	return message
}

func LastUserPrompt(messages []openai2.ChatCompletionMessage) string {
	lastUserPrompt := ""
	for i := len(messages) - 1; i >= 0; i-- {
		if messages[i].Role == openai2.ChatMessageRoleUser {
			lastUserPrompt = messages[i].Content
			break
		}
	}

	return lastUserPrompt
}

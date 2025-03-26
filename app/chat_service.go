package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/openai/openai-go"
)

type ChatService struct {
	chatManager  *ChatManager
	openaiClient *openai.Client
	systemPrompt string
}

func NewChatService() *ChatService {
	client := openai.NewClient()

	return &ChatService{
		chatManager:  NewChatManager(),
		openaiClient: &client,
		systemPrompt: loadSystemPrompt("app/data/system_prompt.txt"),
	}
}

func loadSystemPrompt(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error loading system prompt: %v\n", err)
		return "You are a helpful assistant."
	}
	return string(data)
}

func (cs *ChatService) CreateChat(userID string) string {
	chatID := uuid.New().String()
	cs.chatManager.CreateChat(userID, chatID, cs.systemPrompt)
	return chatID
}

func (cs *ChatService) ProcessMessage(userID, chatID, message string) (string, error) {
	_, exists := cs.chatManager.GetChat(userID, chatID)
	if !exists {
		return "", fmt.Errorf("Chat not found")
	}

	cs.chatManager.AddMessage(userID, chatID, openai.UserMessage(message))

	conversation := cs.chatManager.GetConversation(userID, chatID)

	response, err := cs.openaiClient.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages:            conversation,
			Model:               openai.ChatModelGPT4,
			Temperature:         openai.Float(0.7),
			MaxCompletionTokens: openai.Int(500),
		},
	)
	if err != nil {
		return "", fmt.Errorf("error getting AI response: %v", err)
	}

	aiMessage := response.Choices[0].Message.Content

	cs.chatManager.AddMessage(userID, chatID, openai.AssistantMessage(aiMessage))

	return aiMessage, nil
}

package main

import (
	"github.com/openai/openai-go"
)

type ChatManager struct {
	chats map[string]map[string]openai.ChatCompletionNewParams // user_id -> chat_id -> chat_data
}

func NewChatManager() *ChatManager {
	return &ChatManager{
		chats: make(map[string]map[string]openai.ChatCompletionNewParams),
	}
}

func (cm *ChatManager) CreateChat(userID, chatID, systemPrompt string) {
	if _, exists := cm.chats[userID]; !exists {
		cm.chats[userID] = make(map[string]openai.ChatCompletionNewParams)
	}

	cm.chats[userID][chatID] = openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
		},
	}
}

func (cm *ChatManager) GetChat(userID, chatID string) (openai.ChatCompletionNewParams, bool) {
	userChats, exists := cm.chats[userID]
	if !exists {
		return openai.ChatCompletionNewParams{}, false
	}
	chat, exists := userChats[chatID]
	return chat, exists
}

func (cm *ChatManager) AddMessage(userID, chatID string, message openai.ChatCompletionMessageParamUnion) {
	if chat, exists := cm.GetChat(userID, chatID); exists {
		chat.Messages = append(chat.Messages, message)
		cm.chats[userID][chatID] = chat
	}
}

func (cm *ChatManager) GetConversation(userID, chatID string) []openai.ChatCompletionMessageParamUnion {
	if chat, exists := cm.GetChat(userID, chatID); exists {
		return chat.Messages
	}
	return []openai.ChatCompletionMessageParamUnion{}
}

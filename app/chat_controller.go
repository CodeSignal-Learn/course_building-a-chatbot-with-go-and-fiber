package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
)

type ChatController struct {
	chatService *ChatService
	store       *session.Store
}

func NewChatController(store *session.Store) *ChatController {
	return &ChatController{
		chatService: NewChatService(),
		store:       store,
	}
}

func (cc *ChatController) EnsureUserSession(c *fiber.Ctx) string {
	sess, err := cc.store.Get(c)
	if err != nil {
		// Handle error
		return ""
	}

	userID := sess.Get("user_id")
	if userID == nil {
		// Create a new user ID
		userID = uuid.New().String()
		sess.Set("user_id", userID)
		if err := sess.Save(); err != nil {
			// Handle error
			return ""
		}
	}

	return userID.(string)
}

func (cc *ChatController) CreateChat(c *fiber.Ctx) (map[string]string, error) {
	sess, err := cc.store.Get(c)
	if err != nil {
		return nil, errors.New("session error")
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return nil, errors.New("session expired")
	}

	chatID := cc.chatService.CreateChat(userID.(string))
	return map[string]string{
		"chat_id": chatID,
		"message": "Chat created successfully",
	}, nil
}

func (cc *ChatController) SendMessage(c *fiber.Ctx) (map[string]string, int) {
	sess, err := cc.store.Get(c)
	if err != nil {
		return map[string]string{"error": "Session error"}, 500
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return map[string]string{"error": "Session expired"}, 401
	}

	// Parse request body
	var request struct {
		ChatID      string `json:"chat_id"`
		UserMessage string `json:"user_message"`
	}

	if err := c.BodyParser(&request); err != nil {
		return map[string]string{"error": "Invalid request body"}, 400
	}

	chatID := request.ChatID
	userMessage := request.UserMessage

	if chatID == "" || userMessage == "" {
		return map[string]string{"error": "Missing chat_id or message"}, 400
	}

	aiResponse, err := cc.chatService.ProcessMessage(userID.(string), chatID, userMessage)
	if err != nil {
		if err.Error() == "Chat not found" {
			return map[string]string{"error": err.Error()}, 404
		}
		return map[string]string{"error": err.Error()}, 500
	}

	return map[string]string{"message": aiResponse}, 200
}

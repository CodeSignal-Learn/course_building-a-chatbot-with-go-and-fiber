package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize the Fiber application with HTML template engine
	engine := html.New("./app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static/style.css", "./app/static/style.css")

	// Setup session store
	store := session.New()

	// Create a new chat controller
	chatController := NewChatController(store)

	// Define a route for the index page
	app.Get("/", func(c *fiber.Ctx) error {
		// Ensure user session exists
		chatController.EnsureUserSession(c)
		return c.Render("chat", fiber.Map{})
	})

	// Define a route for creating a chat session
	app.Post("/api/create_chat", func(c *fiber.Ctx) error {
		response, err := chatController.CreateChat(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(response)
	})

	// Define a route for sending messages
	app.Post("/api/send_message", func(c *fiber.Ctx) error {
		response, statusCode := chatController.SendMessage(c)
		return c.Status(statusCode).JSON(response)
	})

	// Run the Fiber application
	app.Listen(":3000")
}

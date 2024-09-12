package app

import (
	"fmt"
	"main/config"
	"main/config/db"
	message_app "main/src/message/application"
	message_adapter "main/src/message/infrastructure/adapter"

	chat_app "main/src/chat/application"
	chat_adapter "main/src/chat/infrastructure/adapter"

	"os"
	"path/filepath"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	currentDir, err := os.Getwd()
	if err != nil {
		fiberlog.Fatalf("Error getting current working directory: %v", err)
	}

	fiberlog.Infof("Current Dir: %s", currentDir)

	sql_tables_inyection_path := filepath.Join(currentDir, "/config/db/init.sql")
	sql_data_inyection_path := filepath.Join(currentDir, "/config/db/data.sql")
	env_path := filepath.Join(currentDir, ".env")

	// Load environment variables
	ENV_CONFIG, err := config.LoadConfig(env_path)
	if err != nil {
		fiberlog.Fatalf("Error while loading .env file: %s", err)
	}

	fiberlog.Infof("Environment: %s", ENV_CONFIG.ENV)

	psqlClient := config.GetPostgreSQLClient(env_path)
	defer psqlClient.Close()

	if ENV_CONFIG.ENV == "dev" {
		db.SQLInjection(sql_tables_inyection_path, psqlClient)
		db.SQLInjection(sql_data_inyection_path, psqlClient)
	}

	// Create a new Fiber app with custom JSON encoder and decoder
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, customize for production
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Enable Logger for fiber
	app.Use(logger.New(logger.Config{
		Format:     "${pid} [${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/Lima",
	}))

	// Set up the root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BACKEND TO USE CHAT-GPT API: https://github.com/luizarnoldch/next_go_chatgpt_clone")
	})

	// Set up the file upload route
	app.Post("/upload", UploadFile)

	app.Post("/chat/completion", ChatCompletion)

	// Initialize repositories, services, and controllers
	messageRepo := message_adapter.NewMessagePSQLRepository(psqlClient)
	messageService := message_app.NewMessagePSQLServie(messageRepo)
	messageController := MessageController{s: messageService}

	message_api_group := app.Group("/messages") // /api

	// Add the routes for message functions
	message_api_group.Get("/", messageController.GetAllMessages)      // Get all messages
	message_api_group.Get("/:id", messageController.GetMessageByID)   // Get a specific message by ID
	message_api_group.Post("/", messageController.CreateMessage)      // Create a new message
	message_api_group.Put("/:id", messageController.UpdateMessage)    // Update a message by ID
	message_api_group.Delete("/:id", messageController.DeleteMessage) // Delete a message by ID

	// Initialize Chat-related components
	chatRepo := chat_adapter.NewChatPSQLRepository(psqlClient)
	chatService := chat_app.NewChatPSQLServie(chatRepo)
	chatController := ChatController{s: chatService}

	chat_api_group := app.Group("/chats") // /api

	// Add the routes for chat functions
	chat_api_group.Get("/", chatController.GetAllChats)      // Get all chats
	chat_api_group.Get("/:id", chatController.GetChatByID)   // Get a specific chat by ID
	chat_api_group.Post("/", chatController.CreateChat)      // Create a new chat
	chat_api_group.Put("/:id", chatController.UpdateChat)    // Update a chat by ID
	chat_api_group.Delete("/:id", chatController.DeleteChat) // Delete a chat by ID

	// Start the server
	URL_API := fmt.Sprint(ENV_CONFIG.MICRO.API.API_HOST, ":", ENV_CONFIG.MICRO.API.API_PORT)
	err = app.Listen(URL_API)
	if err != nil {
		fiberlog.Fatalf("Error starting server: %v", err)
	}
}

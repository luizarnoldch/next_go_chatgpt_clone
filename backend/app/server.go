package app

import (
	"fmt"
	"main/config"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	// Load environment variables
	ENV, err := config.LoadConfig(".env")
	if err != nil {
		fiberlog.Fatalf("Error while loading .env file: %s", err)
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

	// Start the server
	URL_API := fmt.Sprint(ENV.MICRO.API.API_HOST, ":", ENV.MICRO.API.API_PORT)
	app.Listen(URL_API)
}

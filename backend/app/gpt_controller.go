package app

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func ChatCompletion(c *fiber.Ctx) error {

	var chatRequest ChatRequest

	// Parsear el cuerpo JSON enviado por el cliente
	if err := c.BodyParser(&chatRequest); err != nil {
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	MY_CHATGPT_TOKEN := ""

	client := openai.NewClient(MY_CHATGPT_TOKEN)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chatRequest.Content,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(resp)
}

package app

import (
	"main/src/chat/application"
	"main/src/chat/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

type ChatController struct {
	s application.ChatService
}

// GetAllChats handles retrieving all chats
func (c *ChatController) GetAllChats(ctx *fiber.Ctx) error {
	chats, err := c.s.GetAllChats()
	if err != nil {
		fiberlog.Errorf("Error retrieving all chats: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve chats",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Chats retrieved successfully",
		"data":    chats,
	})
}

// CreateChat handles creating a new chat
func (c *ChatController) CreateChat(ctx *fiber.Ctx) error {
	var chat model.Chat

	// Parse the request body into the chat model
	if err := ctx.BodyParser(&chat); err != nil {
		fiberlog.Errorf("Error parsing chat request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
		})
	}

	// Call the service to create the chat
	if err := c.s.CreateChat(&chat); err != nil {
		fiberlog.Errorf("Error creating chat: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create chat",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Chat created successfully",
		"data":    chat,
	})
}

// GetChatByID handles retrieving a chat by ID
func (c *ChatController) GetChatByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid chat ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid chat ID",
		})
	}

	chat, err := c.s.GetChatByID(id)
	if err != nil {
		fiberlog.Errorf("Error retrieving chat by ID: %v", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Chat not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Chat retrieved successfully",
		"data":    chat,
	})
}

// UpdateChat handles updating an existing chat
func (c *ChatController) UpdateChat(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid chat ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid chat ID",
		})
	}

	var chat model.Chat
	if err := ctx.BodyParser(&chat); err != nil {
		fiberlog.Errorf("Error parsing chat request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
		})
	}

	chat.ID = id

	if err := c.s.UpdateChat(&chat); err != nil {
		fiberlog.Errorf("Error updating chat: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update chat",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Chat updated successfully",
		"data":    chat,
	})
}

// DeleteChat handles deleting a chat by ID
func (c *ChatController) DeleteChat(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid chat ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid chat ID",
		})
	}

	if err := c.s.DeleteChat(id); err != nil {
		fiberlog.Errorf("Error deleting chat: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete chat",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Chat deleted successfully",
	})
}

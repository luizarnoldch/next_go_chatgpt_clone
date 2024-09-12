package app

import (
	"main/src/message/application"
	"main/src/message/domain/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

type MessageController struct {
	s application.MessageService
}

// GetAllMessages retrieves all messages
func (c *MessageController) GetAllMessages(ctx *fiber.Ctx) error {
	messages, err := c.s.GetAllMessages()
	if err != nil {
		fiberlog.Errorf("Error retrieving all messages: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to retrieve messages",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Messages retrieved successfully",
		"data":    messages,
	})
}

// CreateMessage creates a new message
func (c *MessageController) CreateMessage(ctx *fiber.Ctx) error {
	var message model.Message

	// Parse request body into message object
	if err := ctx.BodyParser(&message); err != nil {
		fiberlog.Errorf("Error parsing message request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
		})
	}

	// Attempt to create the message
	if err := c.s.CreateMessage(&message); err != nil {
		fiberlog.Errorf("Error creating message: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create message",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Message created successfully",
		"data":    message,
	})
}

// GetMessageByID retrieves a single message by its ID
func (c *MessageController) GetMessageByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid message ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid message ID",
		})
	}

	// Fetch message by ID
	message, err := c.s.GetMessageByID(id)
	if err != nil {
		fiberlog.Errorf("Error retrieving message by ID: %v", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Message not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Message retrieved successfully",
		"data":    message,
	})
}

// UpdateMessage updates an existing message
func (c *MessageController) UpdateMessage(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid message ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid message ID",
		})
	}

	var message model.Message
	// Parse request body into message object
	if err := ctx.BodyParser(&message); err != nil {
		fiberlog.Errorf("Error parsing message request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request payload",
		})
	}

	message.ID = id

	// Attempt to update the message
	if err := c.s.UpdateMessage(&message); err != nil {
		fiberlog.Errorf("Error updating message: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to update message",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Message updated successfully",
		"data":    message,
	})
}

// DeleteMessage deletes a message by its ID
func (c *MessageController) DeleteMessage(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fiberlog.Errorf("Invalid message ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid message ID",
		})
	}

	// Attempt to delete the message
	if err := c.s.DeleteMessage(id); err != nil {
		fiberlog.Errorf("Error deleting message: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete message",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Message deleted successfully",
	})
}

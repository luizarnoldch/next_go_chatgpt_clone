package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func UploadFile(c *fiber.Ctx) error {
	// Get the file from the form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to retrieve the file")
	}

	// Get the file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))

	allowedExtensions := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
		".pdf":  true,
	}

	fiberlog.Info(file.Filename)
	fiberlog.Info(ext)

	if !allowedExtensions[ext] {
		return c.Status(fiber.StatusBadRequest).SendString("File type not allowed")
	}

	// Validate the content type
	contentType := file.Header.Get("Content-Type")
	allowedContentTypes := map[string]bool{
		"image/png":       true,
		"image/jpeg":      true,
		"application/pdf": true,
	}

	if !allowedContentTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).SendString("File type not allowed")
	}

	// Generate a new file name using UUID and keep the original extension
	newFileName := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Log the received file
	fiberlog.Infof("Received file: %s, Type: %s", file.Filename, contentType)

	// Specify the path where the file will be saved with the new name
	filePath := fmt.Sprintf("./uploads/%s", newFileName)

	// Create the 'uploads' folder if it doesn't exist
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err = os.Mkdir("./uploads", 0755)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create the uploads directory")
		}
	}

	// Save the file to the specified folder with the new name
	err = c.SaveFile(file, filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save the file")
	}

	return c.SendString(fmt.Sprintf("File uploaded and saved as %s successfully", newFileName))
}
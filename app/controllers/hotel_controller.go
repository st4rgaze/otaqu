package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/st4rgaze/otaqu/app/models"
)

// get all hotel
func GetAllHotels(c *fiber.Ctx) error {
	hotels, err := models.GetAll()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    hotels,
		"success": true,
	})
}

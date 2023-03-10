package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nkolentcev/fly_service_srv/config"
	"github.com/nkolentcev/fly_service_srv/internal/models"
	"gorm.io/gorm"
)

func FindUserByPin(c *fiber.Ctx) error {
	userPin := c.Params("userPin")
	fmt.Println(userPin)
	var user models.User
	result := config.DB.First(&user, "personal_number = ?", userPin)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

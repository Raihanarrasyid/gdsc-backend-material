package controller

import (
	"restful-api/db"
	"restful-api/helper"
	"restful-api/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	db := db.Connection

	if db == nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Database connection failed",
		})
	}

	var users []model.User

	db.Model(&model.User{}).Preload("Notes").Find(&users)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User found",
		"data":    users,
	})
}

func UserRegister(c *fiber.Ctx) error {
	db := db.Connection

	request := new(helper.RegisterUserRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Request body is not valid",
		})
	}

	user := model.User{
		Username: request.Username,
	}

	db.Create(&user)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User created",
		"data":    user,
	})
}
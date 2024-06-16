package controller

import (
	"restful-api/db"
	"restful-api/model"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	db := db.Connection

	note := model.Note{
		Title:  "Title",
		Body:   "Body",
		UserID: 1,
	}

	db.Create(&note)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Note created",
		"data":    note,
	})
}
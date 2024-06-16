package routes

import (
	"restful-api/controller"

	"github.com/gofiber/fiber/v2"
)
	

func UserRoutes(app *fiber.App) {
	app.Get("/api/user", controller.GetAllUser)
	app.Get("/api/note", controller.CreateNote)
	app.Post("/api/user", controller.UserRegister)
}
package main

import (
	"log"
	"restful-api/db"
	"restful-api/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var app = fiber.New();

func init() {
	var err error
	dialect := mysql.Open("root:@tcp(localhost:3306)/fiber?charset=utf8mb4&parseTime=True&loc=Local")
	db.Connection, err = gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Database connection failed")
	}
}

func main() {
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	routes.UserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
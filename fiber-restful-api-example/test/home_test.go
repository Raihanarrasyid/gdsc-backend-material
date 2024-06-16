package test

import (
	"restful-api/db"
	"testing"

	"github.com/gofiber/fiber/v2/log"
)

func TestHelloWorld(t *testing.T) {
}

func TestDatabase(t *testing.T) {
	db.Connect()

	if db.Connection == nil {
		t.Error("Expected *gorm.DB, got nil")
	} else {
		log.Info("Database connection established")
	}
}
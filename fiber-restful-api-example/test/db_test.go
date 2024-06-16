package test

import (
	"restful-api/db"
	"restful-api/model"
	"testing"
)

func TestMigrate(t *testing.T) {
	db := db.Connect()
	db.AutoMigrate(&model.User{}, &model.Note{}, &model.Post{})
}
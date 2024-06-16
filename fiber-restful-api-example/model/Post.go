package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      int64  `gorm:"primary_key;column:id;autoIncrement"`
	Title   string `gorm:"column:title"`
	Body    string `gorm:"column:body"`
	UserID  int64  `gorm:"column:user_id"`
	User    User   `gorm:"foreignKey:user_id;references:id"`
	LikedBy []User `gorm:"many2many:post_likes;foreignKey:id;joinForeignKey:post_id;references:id;joinReferences:user_id;"`
}
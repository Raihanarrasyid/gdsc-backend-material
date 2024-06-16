package model

import "gorm.io/gorm"

type User struct {
	gorm.Model   
	ID       int64  `gorm:"primary_key;column:id;autoIncrement"`
	Username string `gorm:"column:username"`
	Notes	 []Note `gorm:"foreignKey:user_id;references:id"`
	Posts	 []Post `gorm:"foreignKey:user_id;references:id"`
	PostLikes []Post `gorm:"many2many:post_likes;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:post_id;"`
}

func (u *User) TableName() string {
	return "users"
}
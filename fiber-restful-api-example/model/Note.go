package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	Title  string `gorm:"column:title"`
	Body   string `gorm:"column:body"`
	UserID int64  `gorm:"column:user_id"`
	User   User   `gorm:"foreignKey:user_id;references:id"`
}

func (n *Note) TableName() string {
	return "notes"
}
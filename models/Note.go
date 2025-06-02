package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

type NoteInput struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

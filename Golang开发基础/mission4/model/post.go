package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
	UserID   uint   `json:"userID" gorm:"not null"`
	Comments []Comment
}

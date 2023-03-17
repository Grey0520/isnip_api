package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"password"`
	Snippets []Snippet `json:"snippets"`
	Folders  []Folder  `json:"folders"`
}

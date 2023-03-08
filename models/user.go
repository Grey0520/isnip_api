package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"-"`
	Snippets []Snippet `json:"snippets"`
	Folders  []Folder  `json:"folders"`
}

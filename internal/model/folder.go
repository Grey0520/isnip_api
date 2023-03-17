package model

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name     string    `json:"name"`
	Snippets []Snippet `json:"snippets"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user"`
}

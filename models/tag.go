package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string    `gorm:"uniqueIndex" json:"name"`
	Snippets []Snippet `gorm:"many2many:snippet_tags;" json:"snippets"`
}

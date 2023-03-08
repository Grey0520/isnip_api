package models

import "gorm.io/gorm"

type Snippet struct {
	gorm.Model
	Title    string `json:"title"`
	Code     string `json:"code"`
	FolderID uint   `json:"folder_id"`
	Folder   Folder `json:"folder"`
	Tags     []Tag  `gorm:"many2many:snippet_tags;" json:"tags"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
}

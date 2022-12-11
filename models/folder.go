package models

import (
	"time"
)

type Folder struct {
	FolderID        uint64 `json:"id" db:"folder_id"`
	FolderName      string `json:"name" db:"folder_name"`
	UserID          uint64 `json:"user_id" db:"created_by"`
    CreateTime  time.Time `json:"created_time" db:"created_at"`
    UpdateTime time.Time `json:"updated_time" db:"updated_at"`
	DefaultLanguage string `json:"default_language" db:"defaultLanguage"`
}

type FolderDetail struct {
	FolderID     uint64    `json:"community_id" db:"folder_id"`
	FolderIDName string    `json:"community_name" db:"folder_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	CreateTime   time.Time `json:"create_time" db:"created_time"`
}

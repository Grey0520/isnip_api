package models
type Snippet struct {
	SnipID      uint64    `db:"id" json:"snip_id"`
	Name        string    `db:"name"  json:"name"`
	Language    string    `db:"language" json:"language"`
	FolderID    uint64    `db:"folder_id" json:"folder_id"`
	TagID       uint64    `db:"tag_id" json:"tag_id"`
	Value       string    `db:"content" json:"value"`
	Desc        string    `db:"desc"  json:"description"`
	IsDelete    bool      `db:"isDeleted" json:"isDeleted"`
	IsFavorites bool      `db:"isFavorites" json:"isFavourites"`
	CreateTime  uint64 `db:"created_at" json:"create_time"`
	CreateBy    uint64    `db:"created_by" json:"create_by`
    UpdateTime  uint64 `db:"updated_at" json:"update_time"`
    ModifyBy    string  `db:"modified_by"   json:"update_by`
}

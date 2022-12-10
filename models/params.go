package models

const OrderTime = "time"

type ParamSnippetList struct {
	FolderID uint64 `json:"folder_id" form:"folder_id"`              // 可以为空
	Page     int64  `json:"page" form:"page"`                        // 页码
	Size     int64  `json:"size" form:"size"`                        // 每页数量
	Order    string `json:"order" form:"order" example:"created_id"` // 排序依据
}

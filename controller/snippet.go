package controller

import (
	"fmt"

	"github.com/Grey0520/isnip_api/dao/mysql"
	"github.com/Grey0520/isnip_api/models"
	"github.com/Grey0520/isnip_api/utils/snowflake"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 评论

// SnippetHandler 创建Snippet
func SnippetHandler(c *gin.Context) {
	var Snippet models.Snippet
	if err := c.BindJSON(&Snippet); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 生成SnippetID
	SnippedID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	Snippet.SnipID = SnippedID
	Snippet.CreateBy = userID

	// 创建Snippet
	if err := mysql.CreateSnippet(&Snippet); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func SnippetListHandler(c *gin.Context) {
    userID, err := getCurrentUserID(c)
	if err!=nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	posts, err := mysql.GetSnippetListByUserID(userID)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)
}

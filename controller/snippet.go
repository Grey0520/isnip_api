package controller

import (
	"fmt"

	"github.com/Grey0520/isnip_api/dao/mysql"
	"github.com/Grey0520/isnip_api/models"
	"github.com/Grey0520/isnip_api/utils/snowflake"
	//    "github.com/Grey0520/isnip_api/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 评论

// CreateSnippetHandler 创建Snippet
func CreateSnippetHandler(c *gin.Context) {
	var Snippet models.Snippet
	if err := c.ShouldBindJSON(&Snippet); err != nil {
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
	fmt.Println(Snippet)
	// 创建Snippet
	if err := mysql.CreateSnippet(&Snippet); err != nil {
        zap.L().Error("mysql.CreateSnippet(&Snippet) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SnippetListHandler 代码片段列表
func SnippetListHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	snippets, err := mysql.GetSnippetListByUserID(userID)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, snippets)
}

func UpdateSnippetHandler(c *gin.Context) {
	var Snippet models.Snippet
	if err := c.ShouldBindJSON(&Snippet); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	
	fmt.Println(Snippet)
	// 创建Snippet
	if err := mysql.UpdateSnippet(&Snippet); err != nil {
        zap.L().Error("mysql.CreateSnippet(&Snippet) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

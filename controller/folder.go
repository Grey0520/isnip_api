package controller

import (
	"fmt"

	"github.com/Grey0520/isnip_api/models"
	"github.com/Grey0520/isnip_api/utils/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
    "github.com/Grey0520/isnip_api/dao/mysql"
)

func CreateFolderHandler(c *gin.Context) {
	var Folder models.Folder
	if err := c.ShouldBindJSON(&Folder); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	FolderID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetId() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	Folder.FolderID = FolderID
	Folder.UserID = userID

    if err:= mysql.CreateFolder(&Folder); err != nil {
        zap.L().Error("mysql.CreateFolder(&Folder) failed", zap.Error(err))
        ResponseError(c, CodeServerBusy)
        return
    }
    ResponseSuccess(c, nil)
}

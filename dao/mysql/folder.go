package mysql

import (
	"database/sql"

	"github.com/Grey0520/isnip_api/models"

	"go.uber.org/zap"
)

func GetFolderByID(id uint64) (community *models.FolderDetail, err error) {
	community = new(models.FolderDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
from community
where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err == sql.ErrNoRows { // 查询为空
		err = ErrorInvalidID // 无效的ID
		return
	}
	if err != nil {
		zap.L().Error("query folder failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
	}
	return community, err
}

func CreateFolder(folder *models.Folder) (err error) {
    sqlStr := `insert into folders(folder_id,folder_name,created_by,defaultLanguage) values(?,?,?,?)`
    _, err = db.Exec(sqlStr, folder.FolderID,folder.FolderName,folder.UserID,folder.DefaultLanguage)
    if err!=nil {
        zap.L().Error("insert folder failed", zap.Error(err))
        err = ErrorInsertFailed
        return
    }
    return
}

func GetFolderListByUserID(userID uint64) (FolderList []*models.Folder, err error) {
    sqlStr := `SELECT * FROM folders Where created_by = ?`
    err = db.Select(&FolderList, sqlStr, userID)
    if err == sql.ErrNoRows {
        err = ErrorInvalidID
        return
    }
    if err != nil {
        zap.L().Error("query folder failed", zap.String("sql", sqlStr),zap.Error(err))
        err = ErrorQueryFailed
        return
    }
    return
}
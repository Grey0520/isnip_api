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

package mysql

import (
	"github.com/Grey0520/isnip_api/models"
    "database/sql"
//	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreateSnippet(snippet *models.Snippet) (err error) {
	sqlStr := `insert into snippets(
id, content, language, created_by , folder_id)
values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, snippet.SnipID, snippet.Value, snippet.Language,
		snippet.CreateBy, snippet.FolderID)
	if err != nil {
		zap.L().Error("insert snippet failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetSnippetListByUserID(userID uint64) (SnippetList *models.Snippet, err error) {
    Snippet := new(models.Snippet)
    sqlStr := `SELECT * FROM snippets Where created_by = ?`
    err = db.Get(Snippet,sqlStr,userID)
    if err == sql.ErrNoRows {
        err = ErrorInvalidID
        return
    }
    if err != nil {
        zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
        err = ErrorQueryFailed
        return
    }
    return Snippet,err

}
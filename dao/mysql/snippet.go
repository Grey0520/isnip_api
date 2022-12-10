package mysql

import (
	"database/sql"
	"strings"
	"time"

	"github.com/Grey0520/isnip_api/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func CreateSnippet(snippet *models.Snippet) (err error) {
	sqlStr := `insert into
snippets(id,name,language,folder_id,tag_id,content,snippets.desc,isDeleted,isFavorites,created_at,created_by)
values(?,?,?,?,?,?,?,?,?,?,?)`
	snippet.CreateTime = time.Now()
	snippet.UpdateTime = time.Now()
	_, err = db.Exec(sqlStr, snippet.SnipID, snippet.Name, snippet.Language, snippet.FolderID, snippet.TagID, snippet.Value, snippet.Desc, snippet.IsDelete, snippet.IsFavorites, snippet.CreateTime, snippet.CreateBy)
	if err != nil {
		zap.L().Error("insert snippet failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetSnippetListByUserID(userID uint64) (SnippetList []*models.Snippet, err error) {
	//    Snippet := new(models.Snippet)
	sqlStr := `SELECT * FROM snippets Where created_by = ?`
	err = db.Select(&SnippetList, sqlStr, userID)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
		return
	}
	if err != nil {
		zap.L().Error("query post failed", zap.String("sql", sqlStr), zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	return
}

func GetSnippetListByIDs(ids []string) (snippetList []*models.Snippet, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
from post
where post_id in (?)
order by FIND_IN_SET(post_id, ?)`
	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&snippetList, query, args...)
	return
}

func GetSnippetList(page, size int64) (posts []*models.Snippet, err error) {
	sqlStr := `select id,name,language,folder_id,tag_id,content,snippets.desc,isDeleted,isFavorites,created_at,created_by
from post
ORDER BY created_at
DESC
limit ?,?
`
	posts = make([]*models.Snippet, 0, 2) // 0：长度  2：容量
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

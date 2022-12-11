package logic

import (
	"fmt"

	"github.com/Grey0520/isnip_api/dao/mysql"
	"github.com/Grey0520/isnip_api/models"
	//    "github.com/Grey0520/isnip_api/dao/redis"
	"go.uber.org/zap"
)

//func GetPostListNew(p *models.ParamSnippetList) (data []*models.ApiSnippetDetail, err error) {
//	// 根据请求参数的不同,执行不同的业务逻辑
//	if p.FolderID == 0 {
//		// 查所有
//		data, err = GetPostList2(p)
//	} else {
//		// 根据社区id查询
//		data, err = GetCommunityPostList(p)
//	}
//	if err != nil {
//		zap.L().Error("GetPostListNew failed", zap.Error(err))
//		return nil, err
//	}
//	return
//}

func GetSnippetList(page, size int64) (data []*models.ApiSnippetDetail, err error) {
	snippetList, err := mysql.GetSnippetList(page, size)
	if err != nil {
		fmt.Println(err)
		return
	}
	data = make([]*models.ApiSnippetDetail, 0, len(snippetList)) // data 初始化
	for _, snippet := range snippetList {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserByID(snippet.CreateBy)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.Uint64("postID", snippet.CreateBy),
				zap.Error(err))
			continue
		}
		// 根据文件夹id查询文件夹内容
		community, err := mysql.GetFolderByID(snippet.FolderID)
		if err != nil {
			zap.L().Error("mysql.GetFolderByID() failed",
				zap.Uint64("folder_id", snippet.FolderID),
				zap.Error(err))
			continue
		}
		// 接口数据拼接
		postdetail := &models.ApiSnippetDetail{
			Snippet:      snippet,
			FolderDetail: community,
			AuthorName:   user.UserName,
		}
		data = append(data, postdetail)
	}
	return
}

//func GetSnippetListNew(p *models.ParamSnippetList) (data []*models.ApiSnippetDetail, err error) {
//    // 根据请求参数的不同,执行不同的业务逻辑
//    if p.FolderID == 0 {
//        // 查所有
//        data, err = GetPostList2(p)
//    } else {
//        // 根据社区id查询
//        data, err = GetFolderSnippetList(p)
//    }
//    if err != nil {
//        zap.L().Error("GetPostListNew failed",zap.Error(err))
//        return nil, err
//    }
//    return
//}
//
//func GetPostList2(p *models.ParamSnippetList) (data []*models.ApiSnippetDetail, err error) {
//    // 2、去redis查询id列表
//    ids, err := redis.GetPostIDsInOrder(p)
//    if err != nil {
//        return
//    }
//    if len(ids) == 0 {
//        zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
//        return
//    }
//    zap.L().Debug("GetPostList2", zap.Any("ids", ids))
//    // 提前查询好每篇帖子的投票数
//    voteData, err := redis.GetPostVoteData(ids)
//    if err != nil {
//        return
//    }
//
//    // 3、根据id去数据库查询帖子详细信息
//    // 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
//    posts, err := mysql.GetSnippetListByIDs(ids)
//    if err != nil {
//        return
//    }
//    // 将帖子的作者及分区信息查询出来填充到帖子中
//    for idx, post := range posts {
//        // 根据作者id查询作者信息
//        user, err := mysql.GetUserByID(post.AuthorId)
//        if err != nil {
//            zap.L().Error("mysql.GetUserByID() failed",
//                zap.Uint64("postID",post.AuthorId),
//                zap.Error(err))
//            continue
//        }
//        // 根据社区id查询社区详细信息
//        community, err := mysql.GetCommunityByID(post.CommunityID)
//        if err != nil {
//            zap.L().Error("mysql.GetCommunityByID() failed",
//                zap.Uint64("community_id",post.CommunityID),
//                zap.Error(err))
//            continue
//        }
//        // 接口数据拼接
//        postdetail := &models.ApiPostDetail{
//            VoteNum: voteData[idx],
//            Post:            post,
//            CommunityDetail: community,
//            AuthorName:      user.UserName,
//            }
//            data = append(data,postdetail)
//    }
//    return
//}
//
//func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
//    // 2、去redis查询id列表
//    ids, err := redis.GetCommunityPostIDsInOrder(p)
//    if err != nil {
//        return
//    }
//    if len(ids) == 0 {
//        zap.L().Warn("redis.GetCommunityPostList(p) return 0 data")
//        return
//    }
//    zap.L().Debug("GetPostList2", zap.Any("ids", ids))
//    // 提前查询好每篇帖子的投票数
//    voteData, err := redis.GetPostVoteData(ids)
//    if err != nil {
//        return
//    }
//
//    // 3、根据id去数据库查询帖子详细信息
//    // 返回的数据还要按照我给定的id的顺序返回  order by FIND_IN_SET(post_id, ?)
//    posts, err := mysql.GetPostListByIDs(ids)
//    if err != nil {
//        return
//    }
//    // 将帖子的作者及分区信息查询出来填充到帖子中
//    for idx, post := range posts {
//        // 根据作者id查询作者信息
//        user, err := mysql.GetUserByID(post.AuthorId)
//        if err != nil {
//            zap.L().Error("mysql.GetUserByID() failed",
//                zap.Uint64("postID",post.AuthorId),
//                zap.Error(err))
//            continue
//        }
//        // 根据社区id查询社区详细信息
//        community, err := mysql.GetCommunityByID(post.CommunityID)
//        if err != nil {
//            zap.L().Error("mysql.GetCommunityByID() failed",
//                zap.Uint64("community_id",post.CommunityID),
//                zap.Error(err))
//            continue
//        }
//        // 接口数据拼接
//        postdetail := &models.ApiPostDetail{
//            VoteNum: voteData[idx],
//            Post:            post,
//            CommunityDetail: community,
//            AuthorName:      user.UserName,
//            }
//            data = append(data,postdetail)
//    }
//    return
//}

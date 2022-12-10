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

// 根据前端传来的参数动态的获取片段列表
// 按创建时间排序
// 1、获取请求的query string 参数
// 2、去redis查询id列表
// 3、根据id去数据库查询切片详细信息
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamSnippetList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /snippetList [get]

//func SnippetList2Handler(c *gin.Context) {
//	// GET请求参数(query string)： /api/v1/snippet2?page=1&size=10&order=time
//	// 获取分页参数
//	p := &models.ParamSnippetList{
//		Page:  1,
//		Size:  10,
//		Order: models.OrderTime, // magic string
//	}
//	// c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
//	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
//	if err := c.ShouldBindQuery(p); err != nil {
//        zap.L().Error("SnippetList2Handler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParams)
//		return
//	}
//
//	// 获取数据
//	data, err := logic.GetSnippetListNew(p) // 更新：合二为一
//	if err != nil {
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}

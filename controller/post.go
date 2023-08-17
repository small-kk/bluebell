package controller

import (
	"app/logic"
	"app/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子处理函数
func CreatePostHandler(c *gin.Context) {
	//获取参数，参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取用户的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Debug("getCurrentUserID(c) failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = uint64(userID)

	//创建贴子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 根据帖子id查询帖子详情处理函数
func GetPostDetailHandler(c *gin.Context) {
	//获取帖子id
	postIDStr := c.Param("id")
	postId, err := strconv.Atoi(postIDStr)
	if err != nil {
		zap.L().Debug("strconv.Atoi(postID) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//根据帖子id从数据库中查询帖子内容
	apiPostDetail, err := logic.GetPostDetailByPostID(postId)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByPostID(postId) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, apiPostDetail)
}

// GetPostListHandler 获取帖子列表处理函数
func GetPostListHandler(c *gin.Context) {
	//获取分页信息
	page, pageSize := getPageInfo(c)

	//获取数据
	apiPostList, err := logic.GetPostList(page, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, apiPostList)
}

// GetPostList2Handler 根据分数来分页查询帖子列表
func GetPostList2Handler(c *gin.Context) {
	//获取分页信息
	page, pageSize := getPageInfo(c)

	//获取数据
	apiPostList2, err := logic.GetPostList2(page, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, apiPostList2)
}

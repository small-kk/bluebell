package controller

import (
	"app/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 社区处理函数
func CommunityHandler(c *gin.Context) {
	// 查询到所有社区 (community_id,community_name)
	dataList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		// 不轻易将服务端报错暴露给外面
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, dataList)
}

// CommunityDetailHandler 根据id查询社区分类详情处理函数
func CommunityDetailHandler(c *gin.Context) {
	//获取社区分类id
	communityID := c.Param("id")
	communityId, err := strconv.Atoi(communityID)
	if err != nil {
		zap.L().Debug("strconv.Atoi(communityID) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	communityDetailRes, err := logic.GetCommunityDetailByID(communityId)
	if err != nil {
		zap.L().Error("GetCommunityDetailByID(communityId) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityDetailRes)
}

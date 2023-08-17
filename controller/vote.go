package controller

import (
	"app/logic"
	"app/models"
	pkgValidator "app/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 帖子投票处理函数
func PostVoteHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := pkgValidator.RemoveTopStruct(errs.Translate(pkgValidator.Trans))
		zap.L().Error("invalid param", zap.Any("error", errData))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//为帖子投票
	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Debug("VoteForPost",
			zap.Int64("userID", userID),
			zap.String("postID", p.PostID),
			zap.Int8("direction", p.Direction),
		)
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

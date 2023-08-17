package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

var CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前用户UserID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (page int, pageSize int) {
	//获取分页参数
	var err error
	pageStr := c.Query("page")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		zap.L().Error("strconv.Atoi(offsetStr) failed", zap.Error(err))
		page = 0
	}
	//每页显示记录数
	pageSizeStr := c.Query("size")
	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	return page, pageSize
}

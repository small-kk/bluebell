package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code":1001  //程序中的错误码
	"msg" : xxx  //提示信息
	"data": {}  //数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, data interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

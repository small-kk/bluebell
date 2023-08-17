package controller

import (
	"app/logic"
	"app/models"
	pkgValidator "app/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理用户注册请求的函数
func SignUpHandler(c *gin.Context) {

	//1，获取参数和参数校验
	var p = new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		//判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, pkgValidator.RemoveTopStruct(errs.Translate(pkgValidator.Trans)))

		return
	}

	//2，业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("user register failed", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	//3，返回响应
	ResponseSuccess(c, "注册成功")
}

// LoginHandler 处理用户登录处理函数
func LoginHandler(c *gin.Context) {
	//获取请求参数以及参数校验
	var p = new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, pkgValidator.RemoveTopStruct(errs.Translate(pkgValidator.Trans)))
		return
	}

	//创建用户实例
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	//业务逻辑处理
	exist := logic.Login(user)
	//返回响应
	if !exist {
		zap.L().Info("user login failed", zap.String("username", p.Username))
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, user)
}

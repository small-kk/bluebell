package logic

import (
	"app/dao/mysql"
	"app/models"
	appJWT "app/pkg/jwt"
	"app/pkg/snowflake"
	"go.uber.org/zap"
)

// SignUp 用户注册
func SignUp(p *models.ParamSignUp) error {
	//判断用户存不存在
	if err := mysql.CheckUserExistByUserName(p.Username); err != nil {
		return err
	}

	//生成UID
	userID := snowflake.GenID()

	//构造用户实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存到数据库
	return mysql.InsertUser(user)
}

func Login(user *models.User) (exist bool) {
	// 判断用户是否存在于用户表（用户是否已经注册）
	exist = mysql.CheckUserExistByUsernameAndPassword(user)
	//用户存在，则生成jwt
	if exist {
		var err error
		user.AccessToken, user.RefreshToken, err = appJWT.GenToken(user.UserID, user.Username)
		if err != nil {
			zap.L().Error("generate token failed", zap.String("username", user.Username), zap.Error(err))
			return false
		}
	}
	return
}

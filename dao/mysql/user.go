package mysql

import (
	"app/models"
	"app/pkg/encryptMD5"
	"errors"
	"go.uber.org/zap"
)

// InsertUser 插入用户记录
func InsertUser(user *models.User) error {
	//对用户密码进行加密
	user.Password = encryptMD5.EncryptPassword(user.Password)
	//插入数据库
	err := db.Create(user).Error
	return err
}

// CheckUserExistByUserName 判断指定用户名的用户是否存在
func CheckUserExistByUserName(username string) (err error) {
	var count int64
	err = db.Table("users").Where("username=?", username).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// CheckUserExistByUsernameAndPassword 根据用户名和密码判断用户是否存在
func CheckUserExistByUsernameAndPassword(user *models.User) bool {
	//对用户密码加密处理
	encryptPassword := encryptMD5.EncryptPassword(user.Password)

	//从用户表中查询数据是否存在
	// affected 返回找到的记录数
	affected := db.Where("username=?", user.Username).Find(user).RowsAffected
	//if err != nil {
	//	zap.L().Error("user not in mysql_table_users", zap.Error(err))
	//	return false
	//}
	if affected == 0 || user.Password != encryptPassword {
		return false
	}
	return true
}

// GetUserByUserID 根据用户id获取用户信息
func GetUserByUserID(userID uint64) (user *models.User, err error) {
	user = new(models.User)
	err = db.Where("user_id = ?", userID).Find(user).Error
	if err != nil {
		zap.L().Error("db.Where('user_id = ?', userID).Find(user) failed", zap.Error(err))
		return
	}
	return user, err
}

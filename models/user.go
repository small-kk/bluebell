package models

type User struct {
	UserID       int64  `json:"user_id" gorm:"column:user_id"`
	Username     string `json:"user_name" gorm:"column:username"`
	Password     string `json:"-" gorm:"column:password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

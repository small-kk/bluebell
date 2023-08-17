package models

// ParamSignUp 用户注册参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` //注意：在gorm中用 分号，这里是validator用逗号
}

// ParamLogin 用户登录参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 帖子投票参数结构体
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1),反对票(-1),取消投票(0)
}

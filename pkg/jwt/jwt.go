package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var myScrect = []byte("hello,world")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return myScrect, nil
}

// AccessTokenExpireDuration 定义JWT过期时间
const accessTokenExpireDuration = time.Hour * 24      //access token过期时间
const refreshTokenExpireDuration = time.Hour * 24 * 7 //refresh token过期时间

// GenToken 生成JWT
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	//创建一个自己声明的数据
	c := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenExpireDuration).Unix(), //过期时间
			Issuer:    "bluebell",
		},
	}
	// 加密并获取完整编码后的 access token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(myScrect)

	//加密获取完整编码 refresh token
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenExpireDuration).Unix(), //过期时间
		Issuer:    "bluebell",
	}).SignedString(myScrect)
	return aToken, rToken, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, mc, keyFunc)
	if err != nil {
		return nil, err
	}
	// 校验token
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新access token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	//refresh token 无效直接返回
	var token *jwt.Token
	if token, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	if !token.Valid {
		return
	}

	//从旧access token 中解析出claims数据，解析payload负载信息
	var claims = new(MyClaims)
	_, err = jwt.ParseWithClaims(aToken, claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	//从当前access token 是过期错误，并且refresh token 没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return

}

package encryptMD5

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = "hello,world"

// EncryptPassword 加密用户密码
func EncryptPassword(password string) (password_ciphertext string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

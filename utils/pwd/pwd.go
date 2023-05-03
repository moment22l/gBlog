package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd 加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CheckPwd 验证密码 hashedPwd: hash后的密码	plainPwd: 输入的密码
func CheckPwd(hashedPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

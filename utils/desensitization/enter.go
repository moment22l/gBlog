package desensitization

import "strings"

// Tel 手机号脱敏
func Tel(tel string) string {
	// 137 1234 5678
	// 137 **** 5678
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}

// Email 邮箱脱敏
func Email(email string) string {
	// 298913@qq.com
	// 2*****@qq.com
	strList := strings.Split(email, "@")
	if len(strList) != 2 {
		return ""
	}
	return strList[0][:1] + "****@" + strList[1]
}

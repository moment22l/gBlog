package flag

import (
	sysFlag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool   `structs:""`
	User string // -u admin
}

// Parse 解析命令行参数
func Parse() Option {
	db := sysFlag.Bool("db", false, "初始化数据库")
	user := sysFlag.String("u", "", "创建新用户")
	sysFlag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	m := structs.Map(&option)
	for _, val := range m {
		switch val.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	sysFlag.Usage()
}

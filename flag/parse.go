package flag

import "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := flag.Bool("db", false, "初始化数据库")
	flag.Parse()
	return Option{DB: *db}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}

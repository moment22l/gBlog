package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DB           string `yaml:"db"`
	LogLevel     string `yaml:"log_level"`
	Config       string `yaml:"config"`         // 高级配置, 例如charset, parseTime
	MaxIdleConns int    `yaml:"max_idle_conns"` // 最大连接数
	MaxOpenConns int    `yaml:"max_open_conns"` // 最高可容纳
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}

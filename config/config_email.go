package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` // 发件人邮箱
	AuthCode         string `json:"auth_code" yaml:"auth_code"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` // 默认发件人名字
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       // 是否使用ssl
	UseTLS           bool   `json:"use_tls" yaml:"use_tls"`
}

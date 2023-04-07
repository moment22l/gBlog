package config

// Config 整体配置
type Config struct {
	System   System   `yaml:"system"`
	Logger   Logger   `yaml:"logger"`
	Mysql    Mysql    `yaml:"mysql"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
}

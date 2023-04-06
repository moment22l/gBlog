package config

type Config struct {
	System   System   `yaml:"system"`
	Logger   Logger   `yaml:"logger"`
	Mysql    Mysql    `yaml:"mysql"`
	SiteInfo SiteInfo `yaml:"site_info"`
}

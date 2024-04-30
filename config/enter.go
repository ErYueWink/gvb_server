package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	System   System   `yaml:"system"`
	Logger   Logger   `yaml:"logger"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	QQ       QQ       `yaml:"qq"`
	QiNiu    QiNiu    `yaml:"qiniu"`
	Upload   Upload   `yaml:"upload"`
	Redis    Redis    `yaml:"redis"`
}

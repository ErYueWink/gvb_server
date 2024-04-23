package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"`           // ip
	Port         int    `yaml:"port"`           // 端口号
	DB           string `yaml:"db"`             // 连接的数据库
	User         string `yaml:"user"`           // 用户名
	Password     string `yaml:"password"`       // 密码
	Config       string `yaml:"config"`         // mysql相关配置
	LogLevel     string `yaml:"log_level"`      // 日期级别 debug打印所有日志
	MaxIdleConns int    `yaml:"max_idle_conns"` // 最大空闲连接数
	MaxOpenConns int    `yaml:"max_open_conns"` // 最大可容纳
}

// Dsn 数据库连接地址
func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}

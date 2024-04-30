package config

import "fmt"

type Redis struct {
	Host     string `yaml:"host"` // 主机
	Port     int    `yaml:"port"` // 端口号
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"` // 连接池大小
}

func (r Redis) GetAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

package config

type System struct {
	Host string `yaml:"host"` // 主机
	Port int    `yaml:"port"` // 端口号
	Env  string `yaml:"env"`  // 环境
}

package config

type Logger struct {
	Level        string `yaml:"level"`          // 级别
	Prefix       string `yaml:"prefix"`         // 前缀
	Director     string `yaml:"director"`       // 日志存放的目录
	ShowLine     bool   `yaml:"show-line"`      // 是否显示行号
	LogInConsole bool   `yaml:"log-in-console"` // 是否显示日志文件
}

package config

type Mysql struct {
	Host     string `yaml:"host"`      // ip
	Port     string `yaml:"port"`      // 端口号
	DB       string `yaml:"db"`        // 连接的数据库
	User     string `yaml:"user"`      // 用户名
	Password string `yaml:"password"`  // 密码
	Config   string `yaml:"config"`    // mysql相关配置
	LogLevel string `yaml:"log_level"` // 日期级别 debug打印所有日志
}

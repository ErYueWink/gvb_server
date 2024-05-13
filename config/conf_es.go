package config

import "fmt"

type Es struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	UserName string `json:"user_name" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (e Es) URL() string {
	return fmt.Sprintf("%s:%d", e.Host, e.Port)
}

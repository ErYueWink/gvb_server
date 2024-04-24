package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"
)

const yamlFile = "settings.yaml"

// InitConf 初始化配置文件
func InitConf() {
	c := &config.Config{}
	// 读取配置文件
	yamlConf, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(fmt.Sprintf("get YamlConf err%s", err.Error()))
	}
	// 将yaml文件转换为结构体
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("YamlConf Unmarshal err%s", err.Error())
	}
	log.Println("YamlConf Unmarshal success")
	global.Config = c
}

// SetYaml 修改Yaml文件
func SetYaml() error {
	// 将结构体转换为字节数组
	bytes, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(yamlFile, bytes, fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var confPath = "settings.yaml"

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type Config struct {
	System System `yaml:"system"`
}

func ReadConf() {
	byteDate, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(byteDate, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件格式错误:%s", err))
	}
	fmt.Println(config)
}

package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	Config ServerConfig
	Dev    bool //是否开发环境
	Debug  bool // 是否Debug模式
)

func init() {
	var file []byte
	var err error
	// 通过环境变量来判断是否使用默认配置文件，方便开发
	if filename, ok := os.LookupEnv("QiuBlogConfigFileName"); ok {
		file, err = os.ReadFile(filename)
	} else {
		file, err = os.ReadFile("./config/config.yaml")
	}
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误，请检查文件路径--%v", err))
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		panic(fmt.Sprintf("配置流解析错误，请检查：%v", err))
	}
	if Config.Server.AppMode == "development" {
		Dev = true
	}
	if Config.Log.Level == "debug" {
		Debug = true
	}
}

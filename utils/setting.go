package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Config ServerConfig

type ServerConfig struct {
	Server           server
	Redis            redis
	Database         database
	Oss              oss
	ConstructionTime int64
	Push             push
	Frontend         FrontendConfig
}
type FrontendConfig struct {
	UserInfo struct {
		Title  string `json:"title"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Motto  string `json:"motto"`
		MottoE string `json:"mottoE"`
	} `json:"userInfo"`
	FriendChain []struct {
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"friendChain"`
	Global struct {
		RandomImgApi string `json:"randomImgApi"`
	} `json:"global"`
}

func init() {
	var file *ini.File
	var err error
	// 通过环境变量来判断是否使用默认配置文件，方便开发
	if filename, ok := os.LookupEnv("QiuBlogConfigFileName"); ok {
		file, err = ini.Load(fmt.Sprintf("./config/%s.ini", filename))
	} else {
		file, err = ini.Load("./config/config.ini")
	}
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误，请检查文件路径--%s", err))
	}
	Config.ConstructionTime = file.Section("statistics").Key("ConstructionTime").MustInt64(1662525548)
	LoadServer(file)
	LoadRedis(file)
	LoadData(file)
	LoadOss(file)
	LoadPush(file)

	jsonData, err := os.ReadFile("config/config.json")
	if err != nil {
		panic("Error reading front-end configuration file")
	}
	err = json.Unmarshal(jsonData, &Config.Frontend)
	if err != nil {
		panic(fmt.Sprintf("Front-end configuration file parsing error,msg: %v", err))
	}
}

type server struct {
	AppMode  string
	HttpPort string
	Oss      string
	JwtKey   string
	Url      string
}

func LoadServer(file *ini.File) {
	Config.Server.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	Config.Server.HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	Config.Server.Oss = file.Section("server").Key("Oss").MustString("qiniu")
	Config.Server.JwtKey = file.Section("server").Key("JwtKey").MustString("111")
	Config.Server.Url = file.Section("server").Key("Url").MustString("")
}

type redis struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDb       int
}

func LoadRedis(file *ini.File) {
	Config.Redis.RedisHost = file.Section("redis").Key("RedisHost").MustString("localhost")
	Config.Redis.RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
	Config.Redis.RedisPassword = file.Section("redis").Key("RedisPassword").MustString("123456")
	Config.Redis.RedisDb = file.Section("redis").Key("RedisDb").MustInt(0)
}

type database struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

func LoadData(file *ini.File) {
	Config.Database.Db = file.Section("database").Key("Db").MustString("mysql")
	Config.Database.DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	Config.Database.DbPort = file.Section("database").Key("DbPort").MustString("3306")
	Config.Database.DbUser = file.Section("database").Key("DbUser").MustString("root")
	Config.Database.DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	Config.Database.DbName = file.Section("database").Key("DbName").MustString("qiublog")
}

type oss struct {
	QiniuAccessKey string
	QiniuSecretKey string
	QiniuBucket    string
	QiniuSever     string

	AliyunAccessKeyId     string
	AliyunAccessKeySecret string
	AliyunEndpoint        string
	AliyunBucketName      string
}

func LoadOss(file *ini.File) {
	switch Config.Server.Oss {
	case "aliyun":
		Config.Oss.AliyunAccessKeyId = file.Section("aliyun").Key("AccessKeyId").String()
		Config.Oss.AliyunAccessKeySecret = file.Section("aliyun").Key("AccessKeySecret").String()
		Config.Oss.AliyunEndpoint = file.Section("aliyun").Key("Endpoint").String()
		Config.Oss.AliyunBucketName = file.Section("aliyun").Key("BucketName").String()
	case "qiniu":
		Config.Oss.QiniuAccessKey = file.Section("qiniu").Key("AccessKey").String()
		Config.Oss.QiniuSecretKey = file.Section("qiniu").Key("SecretKey").String()
		Config.Oss.QiniuBucket = file.Section("qiniu").Key("Bucket").String()
		Config.Oss.QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
	}
}

type push struct {
	WxPushCorpId  string
	WxPushAgentid string
	WxPushSecret  string
}

func LoadPush(file *ini.File) {
	Config.Push.WxPushCorpId = file.Section("Push").Key("WxPushCorpId").String()
	Config.Push.WxPushAgentid = file.Section("Push").Key("WxPushAgentid").String()
	Config.Push.WxPushSecret = file.Section("Push").Key("WxPushSecret").String()
}

package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Config ServerConfig

type ServerConfig struct {
	Server   server
	SiteInfo siteInfo
	Redis    redis
	Database database
	Oss      oss
	Push     push
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
	LoadServer(file)
	LoadSiteInfo(file)
	LoadRedis(file)
	LoadData(file)
	LoadOss(file)
	LoadPush(file)
}

type server struct {
	AppMode  string
	HttpPort string
	Oss      string
	JwtKey   string
}

func LoadServer(file *ini.File) {
	Config.Server.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	Config.Server.HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	Config.Server.Oss = file.Section("server").Key("Oss").MustString("qiniu")
	Config.Server.JwtKey = file.Section("server").Key("JwtKey").MustString("111")
}

type siteInfo struct {
	Url              string
	Name             string
	User             string
	Email            string
	Desc             string
	ConstructionTime int64
}

func LoadSiteInfo(file *ini.File) {
	Config.SiteInfo.Url = file.Section("siteInfo").Key("Url").MustString("https://邱.cf")
	Config.SiteInfo.Name = file.Section("siteInfo").Key("Name").MustString("💘  Ocyss 的博客")
	Config.SiteInfo.User = file.Section("siteInfo").Key("User").MustString("Ocyss")
	Config.SiteInfo.Email = file.Section("siteInfo").Key("Email").MustString("qiudie@88.com")
	Config.SiteInfo.Desc = file.Section("siteInfo").Key("Desc").MustString("故事很短，满是遗憾。")
	Config.SiteInfo.ConstructionTime = file.Section("siteInfo").Key("ConstructionTime").MustInt64(1662525548)
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
	Config.Push.WxPushCorpId = file.Section("push").Key("WxPushCorpId").String()
	Config.Push.WxPushAgentid = file.Section("push").Key("WxPushAgentid").String()
	Config.Push.WxPushSecret = file.Section("push").Key("WxPushSecret").String()
}

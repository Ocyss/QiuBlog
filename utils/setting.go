package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var Config ConfigStruct

type ConfigStruct struct {
	Server           server
	Redis            redis
	Database         database
	Oss              oss
	ConstructionTime int64
}

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误，请检查文件路径--%s", err))
	}
	Config.ConstructionTime = file.Section("statistics").Key("ConstructionTime").MustInt64(1662525548)
	LoadServer(file)
	LoadRedis(file)
	LoadData(file)
	LoadOss(file)
}

type server struct {
	AppMode                string
	HttpPort               string
	Oss                    string
	JwtKey                 string
	BaiduVerifyCodevaName  string
	GoogleVerifyCodevaName string
}

func LoadServer(file *ini.File) {
	Config.Server.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	Config.Server.HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	Config.Server.Oss = file.Section("server").Key("Oss").MustString("qiniu")
	Config.Server.JwtKey = file.Section("server").Key("JwtKey").MustString("111")
	Config.Server.BaiduVerifyCodevaName = file.Section("server").Key("BaiduVerifyCodevaName").MustString("")
	Config.Server.GoogleVerifyCodevaName = file.Section("server").Key("GoogleVerifyCodevaName").MustString("")
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

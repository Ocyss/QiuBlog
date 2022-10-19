package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	Oss      string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	QiniuAccessKey string
	QiniuSecretKey string
	QiniuBucket    string
	QiniuSever     string

	AliyunAccessKeyId     string
	AliyunAccessKeySecret string
	AliyunEndpoint        string
	AliyunBucketName      string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		panic(fmt.Sprintf("配置文件读取错误，请检查文件路径--%s", err))
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
	LoadAliyun(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	Oss = file.Section("server").Key("Oss").MustString("qiniu")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("qiublog")
}

func LoadQiniu(file *ini.File) {
	QiniuAccessKey = file.Section("qiniu").Key("AccessKey").String()
	QiniuSecretKey = file.Section("qiniu").Key("SecretKey").String()
	QiniuBucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}

func LoadAliyun(file *ini.File) {
	AliyunAccessKeyId = file.Section("aliyun").Key("AccessKeyId").String()
	AliyunAccessKeySecret = file.Section("aliyun").Key("AccessKeySecret").String()
	AliyunEndpoint = file.Section("aliyun").Key("Endpoint").String()
	AliyunBucketName = file.Section("aliyun").Key("BucketName").String()
}

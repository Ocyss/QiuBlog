package utils

type ServerConfig struct {
	Server   Server   `yaml:"Server"`
	Log      Log      `yaml:"Log"`
	SiteInfo SiteInfo `yaml:"SiteInfo"`
}
type Server struct {
	AppMode  string   `yaml:"AppMode"`
	Host     string   `yaml:"Host"`
	Port     string   `yaml:"Port"`
	JwtKey   string   `yaml:"JwtKey"`
	Database Database `yaml:"Database"`
	Oss      Oss      `yaml:"Oss"`
	Push     Push     `yaml:"Push,omitempty"`
}
type Database struct {
	Type     string `yaml:"Type"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	PassWord string `yaml:"PassWord"`
	Name     string `yaml:"Name"`
	Redis    Redis  `yaml:"Redis"`
}
type Redis struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}
type Oss struct {
	Name  string `yaml:"Name"`
	Qiniu struct {
		AccessKey string `yaml:"AccessKey,omitempty"`
		SecretKey string `yaml:"SecretKey,omitempty"`
		Bucket    string `yaml:"Bucket,omitempty"`
		Sever     string `yaml:"Sever,omitempty"`
	} `yaml:"qiniu"`
	Aliyun struct {
		AccessKeyID     string `yaml:"AccessKeyId,omitempty"`
		AccessKeySecret string `yaml:"AccessKeySecret,omitempty"`
		Endpoint        string `yaml:"Endpoint,omitempty"`
		BucketName      string `yaml:"BucketName,omitempty"`
	} `yaml:"aliyun"`
}
type Push struct {
	WxPushCorpID  string `yaml:"WxPushCorpId,omitempty"`
	WxPushAgentid string `yaml:"WxPushAgentid,omitempty"`
	WxPushSecret  string `yaml:"WxPushSecret,omitempty"`
}
type Log struct {
	Enable     bool   `yaml:"Enable,omitempty"`
	Level      string `yaml:"Level,omitempty"`
	Filename   string `yaml:"Filename,omitempty"`
	MaxSize    int    `yaml:"Maxsize,omitempty"`
	MaxAge     int    `yaml:"Maxage,omitempty"`
	MaxBackups int    `yaml:"Maxbackups,omitempty"`
	Compress   bool   `yaml:"Compress,omitempty"`
}

type SiteInfo struct {
	URL              string `yaml:"Url"`
	Name             string `yaml:"Name"`
	User             string `yaml:"User"`
	Email            string `yaml:"Email"`
	Desc             string `yaml:"Desc"`
	ConstructionTime int64  `yaml:"ConstructionTime"`
}

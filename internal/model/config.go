package model

type Config struct {
	Server    *Server
	OssAliyun *OssAliyun `yaml:"oss-aliyun"`
	Database  *Database
	Redis     *Redis
}

type Server struct {
	Pprof   bool
	Address string
}

type OssAliyun struct {
	Endpoint        string
	AccessKeyID     string `yaml:"access-key-id"`
	AccessKeySecret string `yaml:"access-key-secret"`
	BucketName      string `yaml:"bucket-name"`
	BaseURL         string `yaml:"base-url"`
	PublicURL       string `yaml:"public-url"`
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string `yaml:"db-name"`
	Settings string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DBIndex  int `yaml:"db-index"`
}

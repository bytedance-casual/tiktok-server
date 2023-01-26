package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"tiktok-server/configs"
	"tiktok-server/internal/model"
)

const (
	ApiServiceName             = "api_service"
	FeedServiceName            = "feed_service"
	PublishServiceName         = "publish_service"
	UserServiceName            = "user_service"
	EtcdAddress                = "127.0.0.1:2379"
	CPURateLimit       float64 = 80.0
)

var Config *model.Config

// Init 加载程序运行所需的配置文件，默认为 configs/config.yml
// TODO init outside config file
func Init() {
	data, err := configs.Files.ReadFile("config.yml")
	if err != nil {
		panic(fmt.Errorf(`read default "config.yml":%w`, err))
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(fmt.Errorf(`parse "config.yml":%w`, err))
	}
}

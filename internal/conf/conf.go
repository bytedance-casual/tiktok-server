package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"tiktok-server/configs"
	"tiktok-server/internal/model"
)

var Config *model.Config

// Init 加载程序运行所需的配置文件，默认为 configs/config.yml
// TODO init outside config file
func Init() error {
	data, err := configs.Files.ReadFile("config.yml")
	if err != nil {
		return fmt.Errorf(`read default "config.yml":%w`, err)
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return fmt.Errorf(`parse "config.yml":%w`, err)
	}
	return nil
}

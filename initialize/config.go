package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
	"os"
)

// 配置文件名
const (
	ConfigFileDevEnv  = "config_dev.yaml" // 开发环境配置文件名
	ConfigFileProdEnv = "configs.yaml"    // 生产环境配置文件名
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	switch os.Getenv("GOENV") {
	case global.GOENVDev:
		v.SetConfigName(ConfigFileDevEnv) // name of configs file (without extension)
	case global.GOENVProd:
		v.SetConfigName(ConfigFileProdEnv) // name of configs file (without extension)
	}
	v.AddConfigPath(".")    // optionally look for configs in the working directory
	v.SetConfigType("yaml") // set the configs file type
	err := v.ReadInConfig() // find and read the configs file
	if err != nil {         // handle errors reading the configs file
		log.Fatalln(fmt.Errorf("fatal error configs file: %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		// 重载配置
		if err = v.Unmarshal(&global.App.Config); err != nil {
			log.Fatalln(fmt.Errorf("unmarshal configs: %w", err))
		}
	})
	if err = v.Unmarshal(&global.App.Config); err != nil {
		log.Fatalln(fmt.Errorf("unmarshal configs: %w", err))
	}
}

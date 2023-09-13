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
	ConfigFileProdEnv = "config.yaml"     // 生产环境配置文件名
)

// InitConfig 初始化配置
func InitConfig() {
	v := viper.New()
	switch os.Getenv("GOENV") {
	case global.GOENVDev:
		v.SetConfigName(ConfigFileDevEnv) // name of config file (without extension)
	case global.GOENVProd:
		v.SetConfigName(ConfigFileProdEnv) // name of config file (without extension)
	}
	v.AddConfigPath(".")    // optionally look for config in the working directory
	v.SetConfigType("yaml") // set the config file type
	err := v.ReadInConfig() // find and read the config file
	if err != nil {         // handle errors reading the config file
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		// 重载配置
		if err = v.Unmarshal(&global.App.Config); err != nil {
			log.Fatalln(fmt.Errorf("unmarshal config: %w", err))
		}
	})
	if err = v.Unmarshal(&global.App.Config); err != nil {
		log.Fatalln(fmt.Errorf("unmarshal config: %w", err))
	}
}

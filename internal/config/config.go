package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type BootStrap struct {
	Service   Service   `json:"service"`
	Data      Data      `json:"data"`
	Scheduler Scheduler `json:"scheduler"`
}

func NewBootStrap() *BootStrap {
	// 设置配置文件路径
	config := "config.yaml"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s ", err))
	}
	var bs BootStrap

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&bs); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&bs); err != nil {
		fmt.Println(err)
	}
	return &bs
}

type Service struct {
	Env     string `json:"env"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Scheduler struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Data struct {
	Database Database `json:"database"`
	RabbitMQ string   `json:"rabbitmq"`
}

type Database struct {
	Driver string `json:"driver"`
	Source string `json:"source"`
	Debug  bool   `json:"debug"`
}

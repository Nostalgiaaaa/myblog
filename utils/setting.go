package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局变量，用来保存程序所有配置信息
var Conf = new(MultipleConfig)

type MultipleConfig struct {
	*AppConfig
	*MysqlConfig
}

type AppConfig struct {
	AppMode  string
	HttpPort string
}

type MysqlConfig struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     int
}

func Init() {
	viper.SetConfigFile("./config/config.yaml")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed,err = %v\n ", err)
		return
	}

	//把读取到的配置信息反序列化到Conf中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err = %v\n ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改......")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("OnConfigChange viper.Unmarshal failed,err = %v\n ", err)
		}
	})
	//fmt.Printf("Conf = %+v\n", Conf.StartTime)
	return
}

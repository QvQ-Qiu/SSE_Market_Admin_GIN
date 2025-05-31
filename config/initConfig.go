package config

import (
	"github.com/spf13/viper"
)

// 使用viper从配置文件中读取配置

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
